package lights

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/heatxsink/go-hue/hue"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	getAllLightsURL  = "http://%s/api/%s/lights"
	getLightURL      = "http://%s/api/%s/lights/%d"
	setLightStateURL = "http://%s/api/%s/lights/%d/state"
)

type Lights struct {
	Hostname string
	Username string
}

type Light struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	State     State  `json:"state,omitempty"`
	Type      string `json:"type,omitempty"`
	ModelID   string `json:"modelid,omitempty"`
	SWVersion string `json:"swversion,omitempty"`
}

type State struct {
	On             bool      `json:"on"`
	Hue            uint16    `json:"hue,omitempty"`
	Effect         string    `json:"effect,omitempty"`
	Bri            uint8     `json:"bri,omitempty"`
	Sat            uint8     `json:"sat,omitempty"`
	CT             uint16    `json:"ct,omitempty"`
	XY             []float32 `json:"xy,omitempty"`
	Alert          string    `json:"alert,omitempty"`
	TransitionTime uint16    `json:"transitiontime,omitempty"`
	Reachable      bool      `json:"reachable,omitempty"`
	ColorMode      string    `json:"colormode,omitempty"`
}

func New(hostname string, username string) *Lights {
	return &Lights{
		Hostname: hostname,
		Username: username,
	}
}

func (l *Lights) GetLight(lightID int) (Light, error) {
	var ll Light
	url := fmt.Sprintf(getLightURL, l.Hostname, l.Username, lightID)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ll, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ll, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ll, err
	}
	err = json.Unmarshal(contents, &ll)
	if err != nil {
		return ll, err
	}
	ll.ID = lightID
	return ll, err
}

func (l *Lights) RenameLight(lightID int, lightName string) ([]hue.ApiResponse, error) {
	url := fmt.Sprintf(getLightURL, l.Hostname, l.Username, lightID)
	data := fmt.Sprintf("{\"name\": \"%s\"}", lightName)
	post_body := strings.NewReader(data)
	request, err := http.NewRequest("PUT", url, post_body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var apiResponse []hue.ApiResponse
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, err
}

func (l *Lights) SetLightState(lightID int, state State) ([]hue.ApiResponse, error) {
	url := fmt.Sprintf(setLightStateURL, l.Hostname, l.Username, lightID)
	stateJSON, err := json.Marshal(&state)
	if err != nil {
		return nil, err
	}
	post_body := strings.NewReader(string(stateJSON))
	request, err := http.NewRequest("PUT", url, post_body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var apiResponse []hue.ApiResponse
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, err
}

func (l *Lights) GetAllLights() ([]Light, error) {
	url := fmt.Sprintf(getAllLightsURL, l.Hostname, l.Username)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	lightsMap := map[string]Light{}
	err = json.Unmarshal(contents, &lightsMap)
	if err != nil {
		return nil, err
	}
	lights := make([]Light, 0, len(lightsMap))
	for lightID, light := range lightsMap {
		light.ID, _ = strconv.Atoi(lightID)
		lights = append(lights, light)
	}
	return lights, err
}

func (l *Light) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Id:              %d\n", l.ID))
	buffer.WriteString(fmt.Sprintf("Name:            %s\n", l.Name))
	buffer.WriteString(fmt.Sprintf("Type:            %s\n", l.Type))
	buffer.WriteString(fmt.Sprintf("ModelId:         %s\n", l.ModelID))
	buffer.WriteString(fmt.Sprintf("SwVersion:       %s\n", l.SWVersion))
	buffer.WriteString(fmt.Sprintf("State:\n"))
	buffer.WriteString(l.State.String())
	return buffer.String()
}

func (s *State) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("On:              %t\n", s.On))
	buffer.WriteString(fmt.Sprintf("Hue:             %d\n", s.Hue))
	buffer.WriteString(fmt.Sprintf("Effect:          %s\n", s.Effect))
	buffer.WriteString(fmt.Sprintf("Bri:             %d\n", s.Bri))
	buffer.WriteString(fmt.Sprintf("Sat:             %d\n", s.Sat))
	buffer.WriteString(fmt.Sprintf("CT:              %d\n", s.CT))
	if len(s.XY) > 0 {
		buffer.WriteString(fmt.Sprintf("XY:              %g, %g\n", s.XY[0], s.XY[1]))
	}
	buffer.WriteString(fmt.Sprintf("Alert:           %s\n", s.Alert))
	buffer.WriteString(fmt.Sprintf("TransitionTime:  %d\n", s.TransitionTime))
	buffer.WriteString(fmt.Sprintf("Reachable:       %t\n", s.Reachable))
	buffer.WriteString(fmt.Sprintf("ColorMode:       %s\n", s.ColorMode))
	return buffer.String()
}
