package lights

import (
	"fmt"
	"time"
	"testing"
	"../portal"
)

var (
	username_api_key = "ae2b1fca515949e5d54fb22b8ed95575"
	transition_time = uint16(4)
	sleep_seconds = 4
	sleep_ms = 100
	//test_lights = []int { 1, 2, 3, 4, 5, 6, 8 }
	test_lights = []int { 1, 2, 3, 4 }
)

func print_light_state(light Light) {
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

func redo_state(old_state State) State {
	new_state := State {}
	new_state.TransitionTime = transition_time
	new_state.Alert = "none"
	new_state.On = old_state.On
	new_state.Hue = old_state.Hue
	new_state.Effect = old_state.Effect
	new_state.Bri = old_state.Bri
	new_state.Sat = old_state.Sat
	new_state.Ct = old_state.Ct
	new_state.Xy = old_state.Xy
	return new_state
}

func TestGetAllLights(t *testing.T) {
	fmt.Println("lights.GetAllLights()")
	portal := portal.GetPortal()
	lll := NewLights(portal[0].InternalIpAddress, username_api_key)
	lights := lll.GetAllLights()
	fmt.Println("\tLights: ")
	for _, lll := range lights {
		fmt.Println("\t\tId:   ", lll.Id)
		fmt.Println("\t\tName: ", lll.Name)
	    fmt.Println("\t\t------")
	}
}

func TestGetLight(t *testing.T) {
	fmt.Println("lights.GetLight()")
	portal := portal.GetPortal()
	lll := NewLights(portal[0].InternalIpAddress, username_api_key)
	for _, light_id := range test_lights {
		light := lll.GetLight(light_id)
		print_light_state(light)
	}
}

func TestRenameLight(t *testing.T) {
	fmt.Println("lights.RenameLight()")
	portal := portal.GetPortal()
	lll := NewLights(portal[0].InternalIpAddress, username_api_key)
	light_id := 1
	light_before := lll.GetLight(light_id)
	fmt.Println("\t [BEFORE] Name:   ", light_before.Name)
	resp := lll.RenameLight(light_id, "Testing 123")
	fmt.Println(resp[0].Success)
	light_after := lll.GetLight(light_id)
	fmt.Println("\t [AFTER] Name:    ", light_after.Name)
	resp = lll.RenameLight(light_id, light_before.Name)
	fmt.Println(resp[0].Success)
}

func TestSetLightState(t *testing.T) {
	fmt.Println("lights.SetLightState()")
	portal := portal.GetPortal()
	lll := NewLights(portal[0].InternalIpAddress, username_api_key)
	lights_before := make([]Light, len(test_lights))
	// save current state.
	fmt.Println("\n\nBACKING UP original state ...")
	for _, light_id := range test_lights {
		ls := lll.GetLight(light_id)
		lights_before = append(lights_before, ls)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))
	// lets turn them red!
	fmt.Println("\n\nRED ...")
	red := State { On: true, Hue: 65527, Effect: "none", Bri: 13, Sat: 253, Ct: 500, Xy: []float32{0.6736, 0.3221}, Alert: "none", TransitionTime: transition_time }
	for _, light_id := range test_lights {
		r := lll.SetLightState(light_id, red)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))
	// lets turn them blue!
	fmt.Println("\n\nBLUE ...")
	blue := State { On: true, Hue: 46573, Effect: "none", Bri: 254, Sat: 251, Ct: 500, Xy: []float32{0.1754, 0.0556}, Alert: "none", TransitionTime: transition_time }
	for _, light_id := range test_lights {
		r := lll.SetLightState(light_id, blue)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))

	// lets turn them white!
	fmt.Println("\n\nWHITE ...")
	white := State { On: true, Hue: 34495, Effect: "none", Bri: 203, Sat: 232, Ct: 155, Xy: []float32{0.3151, 0.3252}, Alert: "none", TransitionTime: transition_time }
	for _, light_id := range test_lights {
		r := lll.SetLightState(light_id, white)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))
	
	// lets turn them off!
	fmt.Println("\n\nOFF ...")
	off := State { On: false, Hue: 34495, Effect: "none", Bri: 203, Sat: 232, Ct: 155, Xy: []float32{0.3151, 0.3252}, Alert: "none", TransitionTime: 4 }
	for _, light_id := range test_lights {
		r := lll.SetLightState(light_id, off)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))
	
	// lets RESTORE
	fmt.Println("\n\nRESTORING original state ...")
	for _, llll := range lights_before {
		if llll.Id == 0 {
			continue
		}
		original_state := redo_state(llll.State)
		r := lll.SetLightState(llll.Id, original_state)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	fmt.Println("Fin.\n\n")
}