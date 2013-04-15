package hue

import (
	"testing"
	"fmt"
	"github.com/heatxsink/go-hue"
)

var (
	username = "testing"
	device_type = "testing desktop"
	username_api_key = "ae2b1fca515949e5d54fb22b8ed95575"
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

func TestGetConfiguration(t *testing.T) {
	portal := hue.GetPortal()
	fmt.Println("h.GetConfiguration(username_api_key)")
	for i := range portal {
		p := portal[i]
		h := hue.NewHue(p.InternalIpAddress)
		c := h.GetConfiguration(username_api_key)
		fmt.Printf("\tProxyPort: %d\n", c.ProxyPort)
		fmt.Printf("\tUtc: %s\n", c.Utc)
		fmt.Printf("\tName: %s\n", c.Name)
		fmt.Printf("\tSwUpdate: \n")
		fmt.Printf("\t\tUpdateState: %d\n", c.SwUpdate.UpdateState)
		fmt.Printf("\t\tUrl: %s\n", c.SwUpdate.Url)
		fmt.Printf("\t\tText: %s\n", c.SwUpdate.Text)
		fmt.Printf("\t\tNotify: %t\n", c.SwUpdate.Notify)
		fmt.Printf("\tWhitelist: \n")
		for j := range c.Whitelist {
			fmt.Printf("\t\t%s\n", j)
			fmt.Printf("\t\t\tLastUseDate: %s\n", c.Whitelist[j].LastUseDate)
			fmt.Printf("\t\t\tCreateDate: %s\n", c.Whitelist[j].CreateDate)
			fmt.Printf("\t\t\tName: %s\n", c.Whitelist[j].Name)
		}
		fmt.Printf("\tSwVersion: %s\n", c.SwVersion)
		fmt.Printf("\tProxyAddress: %s\n", c.ProxyAddress)
		fmt.Printf("\tMac: %s\n", c.Mac)
		fmt.Printf("\tLinkButton: %t\n", c.LinkButton)
		fmt.Printf("\tIpAddress: %s\n", c.IpAddress)
		fmt.Printf("\tNetMask: %s\n", c.NetMask)
		fmt.Printf("\tGateway: %s\n", c.Gateway)
		fmt.Printf("\tDhcp: %t\n", c.Dhcp)
		t.Log(p.InternalIpAddress)
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