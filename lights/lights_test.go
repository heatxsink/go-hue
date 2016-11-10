package lights

import (
	"os"
	"testing"
	"time"
)

var (
	testUsername      string
	testHostname      string
	testLights        *Lights
	transitionTime    uint16
	sleepSeconds      int
	sleepMilliSeconds int
	testLightNumbers  []int
	redState          State
	blueState         State
	whiteState        State
	offState          State
)

func init() {
	testUsername = os.Getenv("HUE_TEST_USERNAME")
	testHostname = os.Getenv("HUE_TEST_HOSTNAME")
	testLights = New(testHostname, testUsername)

	transitionTime = uint16(4)
	sleepSeconds = 4
	sleepMilliSeconds = 100
	testLightNumbers = []int{1, 2, 3, 4}

	redState = State{On: true, Hue: 65527, Effect: "none", Bri: 13, Sat: 253, CT: 500, XY: []float32{0.6736, 0.3221}, Alert: "none", TransitionTime: transitionTime}
	blueState = State{On: true, Hue: 46573, Effect: "none", Bri: 254, Sat: 251, CT: 500, XY: []float32{0.1754, 0.0556}, Alert: "none", TransitionTime: transitionTime}
	whiteState = State{On: true, Hue: 34495, Effect: "none", Bri: 203, Sat: 232, CT: 155, XY: []float32{0.3151, 0.3252}, Alert: "none", TransitionTime: transitionTime}
	offState = State{On: false, Hue: 34495, Effect: "none", Bri: 203, Sat: 232, CT: 155, XY: []float32{0.3151, 0.3252}, Alert: "none", TransitionTime: 4}
}

func TestGetAllLights(t *testing.T) {
	lights, err := testLights.GetAllLights()
	if err != nil {
		t.Fail()
	} else {
		for _, light := range lights {
			t.Log(light)
		}
	}
}

func TestGetLight(t *testing.T) {
	for _, lightID := range testLightNumbers {
		light, err := testLights.GetLight(lightID)
		if err != nil {
			t.Fail()
		} else {
			t.Log(light)
		}
	}
}

func TestRenameLight(t *testing.T) {
	lightID := 1
	lightBefore, err := testLights.GetLight(lightID)
	if err != nil {
		t.Fail()
	}
	t.Log("[BEFORE]: ", lightBefore.Name)
	// TODO: need to test response.
	_, err = testLights.RenameLight(lightID, "Testing 123")
	if err != nil {
		t.Fail()
	}
	lightAfter, err := testLights.GetLight(lightID)
	t.Log("[AFTER]: ", lightAfter.Name)
	// TODO: need to test response.
	_, err = testLights.RenameLight(lightID, lightBefore.Name)
}

func test_set_light_state(t *testing.T, state State) {
	for _, lightID := range testLightNumbers {
		// TODO: need to test response.
		_, err := testLights.SetLightState(lightID, state)
		if err != nil {
			t.Fail()
		}
		time.Sleep(time.Millisecond * time.Duration(sleepMilliSeconds))
	}
	time.Sleep(time.Second * time.Duration(sleepSeconds))
}

func TestSetLightState(t *testing.T) {
	lightStateBackup := make([]Light, len(testLightNumbers))
	for _, lightID := range testLightNumbers {
		light, err := testLights.GetLight(lightID)
		if err != nil {
			t.Fail()
		}
		lightStateBackup = append(lightStateBackup, light)
	}
	test_set_light_state(t, redState)
	test_set_light_state(t, blueState)
	test_set_light_state(t, whiteState)
	test_set_light_state(t, offState)
	for _, light := range lightStateBackup {
		if light.ID == 0 {
			continue
		}
		light.State.TransitionTime = transitionTime
		light.State.Alert = "none"
		// TODO: need to test response.
		_, err := testLights.SetLightState(light.ID, light.State)
		if err != nil {
			t.Fail()
		}
		time.Sleep(time.Millisecond * time.Duration(sleepMilliSeconds))
	}
}
