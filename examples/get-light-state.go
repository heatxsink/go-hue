package main

import (
	"fmt"
	"../src/portal"
	"../src/lights"
)

var (
	username_api_key = "ae2b1fca515949e5d54fb22b8ed95575"
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

func print_light_state(light lights.State) {
	if light.TransitionTime == 0 {
		light.TransitionTime = 4
	}
	light_struct := fmt.Sprintf("lights.State { On: %t, Hue: %d, Effect: \"%s\", Bri: %d, Sat: %d, Ct: %d, Xy: []float32{%g, %g}, Alert: \"%s\", TransitionTime: %d }", light.On, light.Hue, light.Effect, light.Bri, light.Sat, light.Ct, light.Xy[0], light.Xy[1], light.Alert, light.TransitionTime)
	fmt.Println(light_struct)
}

func main() {
	struct_mode := true
	portal := portal.GetPortal()
	ll := lights.NewLights(portal[0].InternalIpAddress, username_api_key)
	lights := ll.GetAllLights()
	fmt.Println("All Lights: ")
	for _, l := range lights {
		light := ll.GetLight(l.Id)
		if struct_mode {
			fmt.Println(l.Id)
			print_light_state(light.State)
		} else {
			print_light(light)
		}
	}
}