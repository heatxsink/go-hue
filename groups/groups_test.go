package groups

import (
	"fmt"
	"github.com/heatxsink/go-hue/lights"
	"os"
	"testing"
	"time"
)

var (
	testUsername       string
	testHostname       string
	testGroups         *Groups
	transitionTime     uint16
	sleepSeconds       int
	sleepMilliSeconds  int
	testLightNumbers   []int
	redState           lights.State
	blueState          lights.State
	whiteState         lights.State
	offState           lights.State
	onState            lights.State
	virginAmericaState lights.State
)

func init() {
	testUsername = os.Getenv("HUE_TEST_USERNAME")
	testHostname = os.Getenv("HUE_TEST_HOSTNAME")
	testGroups = New(testHostname, testUsername)
	transitionTime = uint16(4)
	sleepSeconds = 4
	sleepMilliSeconds = 100
	testLightNumbers = []int{1, 2, 3, 4}

	virginAmericaState = lights.State{On: true, Hue: 54179, Effect: "none", Bri: 230, Sat: 253, CT: 223, XY: []float32{0.3621, 0.1491}, Alert: "none", TransitionTime: transitionTime}
	redState = lights.State{On: true, Hue: 65527, Effect: "none", Bri: 13, Sat: 253, CT: 500, XY: []float32{0.6736, 0.3221}, Alert: "none", TransitionTime: transitionTime}
	blueState = lights.State{On: true, Hue: 46573, Effect: "none", Bri: 254, Sat: 251, CT: 500, XY: []float32{0.1754, 0.0556}, Alert: "none", TransitionTime: transitionTime}
	whiteState = lights.State{On: true, Hue: 34495, Effect: "none", Bri: 203, Sat: 232, CT: 155, XY: []float32{0.3151, 0.3252}, Alert: "none", TransitionTime: transitionTime}
	offState = lights.State{On: false}
	onState = lights.State{On: true}
}

func TestCreateGroup(t *testing.T) {
	group := Group{Name: "Office", Lights: []string{"1", "2"}}
	_, err := testGroups.CreateGroup(group)
	if err != nil {
		t.Fail()
	}
}

func TestSetGroup(t *testing.T) {
	group := Group{Name: "Office", Lights: []string{"1", "2"}}
	_, err := testGroups.SetGroup(1, group)
	if err != nil {
		t.Fail()
	}
	group = Group{Name: "Bedroom", Lights: []string{"3", "4"}}
	_, err = testGroups.SetGroup(2, group)
	if err != nil {
		t.Fail()
	}
	group = Group{Name: "Living Room", Lights: []string{"5", "6"}}
	_, err = testGroups.SetGroup(3, group)
	if err != nil {
		t.Fail()
	}
	group = Group{Name: "Upstairs", Lights: []string{"1", "2", "3", "4", "5", "6", "8"}}
	_, err = testGroups.SetGroup(4, group)
	if err != nil {
		t.Fail()
	}
}

func TestGetAllGroups(t *testing.T) {
	groups, err := testGroups.GetAllGroups()
	if err != nil {
		t.Fail()
	}
	for _, g := range groups {
		t.Log(g.String())
	}
}

func TestGetGroup(t *testing.T) {
	for _, groupID := range testLightNumbers {
		g, err := testGroups.GetGroup(groupID)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		fmt.Println(g.String())
	}
}

func test_set_group_state(t *testing.T, state lights.State) {
	groups, err := testGroups.GetAllGroups()
	if err != nil {
		t.Fail()
	}
	for _, group := range groups {
		// TODO: need to test response.
		_, err := testGroups.SetGroupState(group.ID, state)
		if err != nil {
			t.Fail()
		}
		time.Sleep(time.Millisecond * time.Duration(sleepMilliSeconds))
	}
}

func TestSetGroupState(t *testing.T) {
	groupsBackup, err := testGroups.GetAllGroups()
	if err != nil {
		t.Fail()
	}
	test_set_group_state(t, offState)
	test_set_group_state(t, virginAmericaState)
	test_set_group_state(t, offState)
	test_set_group_state(t, onState)
	test_set_group_state(t, offState)
	test_set_group_state(t, blueState)
	test_set_group_state(t, offState)
	for _, group := range groupsBackup {
		if group.ID == 0 {
			continue
		}
		group.Action.TransitionTime = transitionTime
		group.Action.Alert = "none"
		// TODO: need to test response.
		_, err := testGroups.SetGroupState(group.ID, group.Action)
		if err != nil {
			t.Fail()
		}
		time.Sleep(time.Millisecond * time.Duration(sleepMilliSeconds))
	}
}
