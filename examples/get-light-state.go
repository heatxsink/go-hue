package main

import (
	"os"
	"flag"
	"fmt"
	"../src/portal"
	"../src/lights"
	"../src/key"
)

var (
	config_filename string = ""
)

func print_light(light lights.Light) {
	fmt.Println("\tLight: ")
	fmt.Println("\t\tId:           ", light.Id)
	fmt.Println("\t\tName:         ", light.Name)
	fmt.Println("\t\tType:         ", light.Type)
	fmt.Println("\t\tModelId:      ", light.ModelId)
	fmt.Println("\t\tSwVersion:    ", light.SwVersion)
	fmt.Println("\t\tState:")
	fmt.Println("\t\t\tOn:         ", light.State.On)
	fmt.Println("\t\t\tHue:        ", light.State.Hue)
	fmt.Println("\t\t\tEffect:     ", light.State.Effect)
	fmt.Println("\t\t\tBri:        ", light.State.Bri)
	fmt.Println("\t\t\tSat:        ", light.State.Sat)
	fmt.Println("\t\t\tCt:         ", light.State.Ct)
	fmt.Println("\t\t\tXy:         ", light.State.Xy)
	fmt.Println("\t\t\tReachable:  ", light.State.Reachable)
	fmt.Println("\t\t\tColorMode:  ", light.State.ColorMode)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: get-light-state -config=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.StringVar(&config_filename, "config", "", "Config filename.")
	flag.Parse()
	flag.Usage = usage
}

func main() {
	if config_filename != "" {
		k := key.New(config_filename)
		portal := portal.GetPortal()
		ll := lights.NewLights(portal[0].InternalIpAddress, k.Username)
		lights := ll.GetAllLights()
		fmt.Println("All Lights: ")
		for _, l := range lights {
			light := ll.GetLight(l.Id)
				print_light(light)
		}
	} else {
		usage()
	}
}
