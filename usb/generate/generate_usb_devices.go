package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	gousb "github.com/google/gousb"
	usbid "github.com/multiverse-os/vinput/libs/uinput/usbid"
)

type USBDevice struct {
	ID         gousb.ID
	Name       string
	VendorID   gousb.ID
	VendorName string
}

type TemplateData struct {
	Devices []USBDevice
}

func main() {
	fmt.Println("Uinput Generate Script: USB Product Device IDs")
	fmt.Println("===============================================================================")
	fmt.Println("Uses the local usb.ids provided by the OS to generate a map of usable USB ID")
	fmt.Println("information needed for using real vendor ID, product ID and names when creating")
	fmt.Println("virtual USB devices. This will allow creation of more realistic virtual devices\n")
	vendors, classes, err := usbid.ParseIDs(strings.NewReader(usbid.IDListData))
	if err != nil {
		fmt.Println("[fatal error] failed to parse ids:", err)
		os.Exit(1)
	} else {
		fmt.Println("loaded [", len(vendors), "] vendors, and [", len(classes), "] classes...")
	}
	var Data TemplateData
	var name string
	for vendorID, vendor := range vendors {
		for productID, product := range vendor.Product {
			name = strings.Replace(strings.Replace(strings.Split(product.Name, ",")[0], "\\", "", -1), "\"", "", -1)
			if len(name) > 80 {
				name = name[:80]
			}
			Data.Devices = append(Data.Devices, USBDevice{
				ID:         productID,
				VendorID:   vendorID,
				VendorName: vendor.Name,
				Name:       name,
			})
		}
	}

	t, err := template.New("devices").Parse(GoCodeTemplate)
	if err != nil {
		fmt.Println("[fatal error] failed to initialize the template with data usb device list:", err)
	}
	fmt.Println("succsessfully created template, now executing it with [", len(Data.Devices), "] devices to render")

	sourceCodeFile, err := os.Create("usb_devices.go")
	if err != nil {
		fmt.Println("[fatal error] failed to open file for writing:", err)
	}

	err = t.Execute(sourceCodeFile, Data)
	if err != nil {
		fmt.Println("[fatal error] failed to execute the template with the templateData:", err)
	}
	sourceCodeFile.Close()
}

var GoCodeTemplate = `package usb

type USBDevice struct {
	ID      	 string   
	Name       string
	VendorID   string
	VendorName string
}

func USBDevices() []USBDevice {
	return []USBDevice{
		{{range .Devices}}USBDevice{ID: "{{.ID}}", Name: "{{.Name}}", VendorID: "{{.VendorID}}", VendorName: "{{.VendorName}}"},
		{{end}}}
}
`
