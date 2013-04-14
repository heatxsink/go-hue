package hue

import (
	"testing"
	"fmt"
	"github.com/heatxsink/go-hue"
)

func TestGetPortal(t *testing.T) {
	portal := hue.GetPortal()
	fmt.Println("hue.GetPortal()")
	for i := range portal {
		fmt.Printf("\tId:                  %s\n", portal[i].Id)
		fmt.Printf("\tInternal Ip Address: %s\n", portal[i].InternalIpAddress)
		fmt.Printf("\tMac Address:         %s\n", portal[i].MacAddress)
		t.Log(portal[i].InternalIpAddress)
	}
}

/*
func TestCreateUsername(t *testing.T) {
	hostname := "10.0.16.16"
	username := ""
	device_type := ""
	h := hue.NewHue(hostname)
	response := h.CreateUsername(username, device_type)
	message := fmt.Sprintf("Verified Api Key (Username MD5 hashed): %s\n", response[0]["success"]["username"])
	t.Log(message)
}

func TestDeleteUsername(t *testing.T) {
	hostname := "10.0.16.16"
	username := ""
	device_type := ""
	h := hue.NewHue(hostname)
	response := h.CreateUsername(username, device_type)
	message := fmt.Sprintf("Verified Api Key (Username MD5 hashed): %s\n", response[0]["success"]["username"])
	t.Log(message)
}
*/