package makejson

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func Execute() {
	// Variable declaration and scan that space :P
	var name string
	var address string

	fmt.Printf("Enter your name: ")
	x := bufio.NewScanner(os.Stdin)

	if x.Scan() {
		name = x.Text()
	}

	fmt.Printf("Enter your address: ")
	y := bufio.NewScanner(os.Stdin)

	if y.Scan() {
		address = y.Text()
	}

	// Create map and link values with map key
	informations := make(map[string]string)

	informations["name"] = name
	informations["address"] = address

	// Create json and create a error treatment
	response, err := json.Marshal(informations)

	if err == nil {
		fmt.Printf("\n\njson:\n%s\n", response)
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}
