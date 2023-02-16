package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	id "github.com/multiverse-os/uinput/usb-id"

	usb "github.com/google/gousb"
)

// To me, the 1st one ("Id") is correct, as "Id" and "" are abbreviations of "identifier". The correct one is "Id". "ID" appears to be an acronym of two words, though there is only one word, "identifier".

type Id struct {
	Device usb.ID
	Vendor usb.ID

type USBDevice struct {
	Id
	Name       string
	VendorName string
}

type TemplateData struct {
	Devices []USBDevice
}

func main() {
	fmt.Println("uinput: parse system usb product device ids")
	fmt.Println("===============================================================================")
	fmt.Println("Uses the local usb.ids provided by the OS to generate a map of usable USB ")
	fmt.Println("information needed for using real vendor , product Id and names when creating")
	fmt.Println("virtual USB devices. This will allow creation of more realistic virtual devices\n")
	vendors, classes, err := usbid.Parses(strings.NewReader(id.IdListData))
	if err != nil {
		fmt.Println("[fatal error] failed to parse ids:", err)
		os.Exit(1)
	} else {
		fmt.Printf("loaded [", len(vendors), "] vendors, and [", len(classes), "] classes...\n")
	}
	var Data TemplateData
	var name string
	for vendorId, vendor := range vendors {
		for productId, product := range vendor.Product {
			name = strings.Replace(strings.Replace(strings.Split(product.Name, ",")[0], "\\", "", -1), "\"", "", -1)
			if len(name) > 80 {
				name = name[:80]
			}
			Data.Devices = append(Data.Devices, USBDevice{
				Id:         productId,
				VendorId:   vendorId,
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
	      	 string   
	Name       string
	Vendor   string
	VendorName string
}

func USBDevices() []USBDevice {
	return []USBDevice{
		{{range .Devices}}USBDevice{: "{{.Id}}", Name: "{{.Name}}", VendorId: "{{.VendorId}}", VendorName: "{{.VendorName}}"},
		{{end}}}
}
`
