package main

import (
	"fmt"
	"bufio"
	"os"
	"../src/portal"
	"../src/configuration"
)

func main() {
	hub_hostname := portal.Discover()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please press the link button on your hub, then press [enter] to continue.")
	reader.ReadLine()
	// username
	fmt.Println("Please enter a username:")
	data, _, _ := reader.ReadLine()
	username := string(data)
	// device type
	fmt.Println("Please enter device type:")
	data1, _, _ := reader.ReadLine()
	device_type := string(data1)
	// lets finally create a username for the api
	c := configuration.NewConfiguration(hub_hostname)
	response := c.CreateUsername(username, device_type)
	fmt.Printf("Verified Api Key (Username MD5 hashed): %s\n", response[0]["success"]["username"])
}