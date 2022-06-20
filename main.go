package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {

	cSuccess := color.New(color.FgWhite).Add(color.BgGreen)
	cError := color.New(color.FgWhite).Add(color.BgRed)

	// Make ssh file
	f, err := os.Create("ssh")

	if err != nil {
		fmt.Println(cError.Sprintf("Error creating ssh file: %s", err))
	}

	defer f.Close()

	// Make wpa_supplicant file
	f, err = os.Create("wpa_supplicant.conf")

	if err != nil {
		fmt.Println(cError.Sprintf("Error creating wpa_supplicant file: %s", err))
	}

	defer f.Close()

	conf_string := `country=JP
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
network={
	ssid="{SSID HERE}"
	psk="{PASSWORD HERE}"
	scan_ssid=1
}`

	// Write wpa_supplicant.conf
	_, err = f.WriteString(conf_string)

	if err != nil {
		fmt.Println(cError.Sprintf("Error writing wpa_supplicant file: %s", err))
	}

	fmt.Println(cSuccess.Sprintf("Successfully created files!"))

}
