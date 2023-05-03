package main

import (
	"fmt"

	"github.com/multiverse-os/uinput"
)

func main() {

	fmt.Printf("uinput-cli")
	fmt.Printf("==========")

	fmt.Printf("Creating a new test keyboard.\n")
	keyboard := uinput.NewKeyboard("test-keyboard")

	fmt.Printf("keyboard: %v\n", keyboard)

}
