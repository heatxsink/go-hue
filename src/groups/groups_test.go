package groups

import (
	"testing"
	"fmt"
	"time"
	"../portal"
	"../lights"
	"../key"
)

var (
	username_api_key_filename = "../../conf/username_api_key.conf"
	username_api_key = ""
	transition_time = uint16(4)
	sleep_seconds = 4
	sleep_ms = 100
	test_groups = []int { 0, 1, 2, 3, 4 }
)

func init() {
	k := key.New(username_api_key_filename)
	username_api_key = k.Username
}

/*
func TestCreateGroup(t *testing.T) {
	fmt.Println("groups.CreateGroup()")
	portal := portal.GetPortal()
	hostname := portal[0].InternalIpAddress
	gg := NewGroup(hostname, username_api_key)
	group := Group { Name: "Office", Lights: []string { "1", "2" } }
	data := gg.CreateGroup(group)
	fmt.Println(data)
}
*/

func redo_state(old_state lights.State) lights.State {
	new_state := lights.State {}
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

func print_group_state(group GroupState) {
	fmt.Println("\tGroup: ")
	fmt.Println("\t\tId:              ", group.Id)
	fmt.Println("\t\tName:              ", group.Name)
	fmt.Println("\t\tAction:")
	fmt.Println("\t\t\tOn:              ", group.Action.On)
	fmt.Println("\t\t\tHue:             ", group.Action.Hue)
	fmt.Println("\t\t\tEffect:          ", group.Action.Effect)
	fmt.Println("\t\t\tBri:             ", group.Action.Bri)
	fmt.Println("\t\t\tSat:             ", group.Action.Sat)
	fmt.Println("\t\t\tCt:              ", group.Action.Ct)
	fmt.Println("\t\t\tXy:              ", group.Action.Xy)
	fmt.Println("\t\t\tAlert:           ", group.Action.Alert)
	fmt.Println("\t\t\tReachable:  ", group.Action.Reachable)
	fmt.Println("\t\t\tColorMode:  ", group.Action.ColorMode)
	fmt.Println("\t\t\tTransitionTime:  ", group.Action.TransitionTime)
	fmt.Println("\n\tLights:")
	for _, light_id := range group.Lights {
		fmt.Println("\t\t\t", light_id)
	}
}

func TestSetGroup(t *testing.T) {
	fmt.Println("groups.SetGroup()")
	portal := portal.GetPortal()
	hostname := portal[0].InternalIpAddress
	gg := NewGroup(hostname, username_api_key)
	// office!
	group_id := 1
	group := Group { Name: "Office", Lights: []string { "1", "2" } }
	data := gg.SetGroup(group_id, group)
	fmt.Println(data)
	// bedroom!
	group_id = 2
	group = Group { Name: "Bedroom", Lights: []string { "3", "4" } }
	data = gg.SetGroup(group_id, group)
	fmt.Println(data)
	// living room!
	group_id = 3
	group = Group { Name: "Living Room", Lights: []string { "5", "6" } }
	data = gg.SetGroup(group_id, group)
	fmt.Println(data)
	// upstairs!
	group_id = 4
	group = Group { Name: "Upstairs", Lights: []string{ "1", "2", "3", "4", "5", "6", "8" } }
	data = gg.SetGroup(group_id, group)
	fmt.Println(data)
}

func TestGetGroups(t *testing.T) {
	fmt.Println("groups.GetGroups()")
	portal := portal.GetPortal()
	hostname := portal[0].InternalIpAddress
	ggg := NewGroup(hostname, username_api_key)
	groups := ggg.GetGroups()
	fmt.Println("\tGroups: ")
	for _, gg := range groups {
		fmt.Println("\t\tId:   ", gg.Id)
		fmt.Println("\t\tName: ", gg.Name)
	    fmt.Println("\t\t------")
	}
}

func TestGetGroup(t *testing.T) {
	fmt.Println("groups.GetGroup()")
	portal := portal.GetPortal()
	hostname := portal[0].InternalIpAddress
	gg := NewGroup(hostname, username_api_key)
	for _, group_id := range test_groups {
		group := gg.GetGroup(group_id)
		print_group_state(group)
	}
}

func TestSetGroupState(t *testing.T) {
	fmt.Println("groups.SetGroupState()")
	portal := portal.GetPortal()
	hostname := portal[0].InternalIpAddress
	gg := NewGroup(hostname, username_api_key)
	_test_groups := []int { 1 }
	groups_before := make([]GroupState, len(_test_groups))
	// save current state.
	fmt.Println("\n\nBACKING UP original state ...")
	for _, group_id := range _test_groups {
		gs := gg.GetGroup(group_id)
		groups_before = append(groups_before, gs)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))
	// lets turn them red!
	fmt.Println("\n\nRED ...")
	red := lights.State { On: true, Hue: 65527, Effect: "none", Bri: 13, Sat: 253, Ct: 500, Xy: []float32{0.6736, 0.3221}, Alert: "none", TransitionTime: transition_time }
	for _, group_id := range _test_groups {
		r := gg.SetGroupState(group_id, red)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))

	// lets turn them off!
	fmt.Println("\n\nOFF ...")
	off := lights.State { On: false }
	for _, group_id := range _test_groups {
		r := gg.SetGroupState(group_id, off)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))

	// lets turn them on!
	fmt.Println("\n\nON ...")
	on := lights.State { On: true }
	for _, group_id := range _test_groups {
		r := gg.SetGroupState(group_id, on)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))


	// lets turn them blue!
	fmt.Println("\n\nBLUE ...")
	blue := lights.State { On: true, Hue: 46573, Effect: "none", Bri: 254, Sat: 251, Ct: 500, Xy: []float32{0.1754, 0.0556}, Alert: "none", TransitionTime: transition_time }
	for _, group_id := range _test_groups {
		r := gg.SetGroupState(group_id, blue)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))

	// lets turn them white!
	fmt.Println("\n\nWHITE ...")
	white := lights.State { On: true, Hue: 34495, Effect: "none", Bri: 203, Sat: 232, Ct: 155, Xy: []float32{0.3151, 0.3252}, Alert: "none", TransitionTime: transition_time }
	for _, group_id := range _test_groups {
		r := gg.SetGroupState(group_id, white)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	time.Sleep(time.Second * time.Duration(sleep_seconds))

	// lets RESTORE
	fmt.Println("\n\nRESTORING original state ...")
	for _, gggg := range groups_before {
		if gggg.Id == 0 {
			continue
		}
		original_state := redo_state(gggg.Action)
		r := gg.SetGroupState(gggg.Id, original_state)
		fmt.Println(r)
		time.Sleep(time.Millisecond * time.Duration(sleep_ms))
	}
	fmt.Println("Fin.\n\n")
}
