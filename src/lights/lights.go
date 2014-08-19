package lights

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/heatxsink/go-hue/src/util"
)

var (
	get_all_lights_url = "http://%s/api/%s/lights"
	get_light_url = "http://%s/api/%s/lights/%d"
	set_light_state_url = "http://%s/api/%s/lights/%d/state"
)


type Lights struct {
	Hostname string
	Username string
}

type Light struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name"`
	State State `json:"State,omitempty"`
	Type string `json:"Type,omitempty"`
	ModelId string `json:"modelid,omitempty"`
	SwVersion string `json:"swversion,omitempty"`
}

type State struct {
	On bool `json:"on"`
	Hue uint16 `json:"hue,omitempty"`
	Effect string `json:"effect,omitempty"`
	Bri uint8 `json:"bri,omitempty"`
	Sat uint8 `json:"sat,omitempty"`
	Ct uint16 `json:"ct,omitempty"`
	Xy []float32 `json:"xy,omitempty"`
	Alert string `json:"alert,omitempty"`
	TransitionTime uint16 `json:"transitiontime,omitempty"`
	Reachable bool `json:"reachable,omitempty"`
	ColorMode string `json:"colormode,omitempty"`
}

func NewLights(hostname string, username string) *Lights {
	light := Lights{Hostname: hostname, Username: username}
	return &light
}

func (l *Lights) GetLight(light_id int) Light {
	url := fmt.Sprintf(get_light_url, l.Hostname, l.Username, light_id)
	response := util.HttpGet(url)
	var light Light
	json.Unmarshal(response, &light)
	light.Id = light_id
	return light
}

func (l *Lights) RenameLight(light_id int, light_name string) []util.ApiResponse {
	url := fmt.Sprintf(get_light_url, l.Hostname, l.Username, light_id)
	data := fmt.Sprintf("{\"name\": \"%s\"}", light_name)
	response := util.HttpPut(url, data, "application/json")
	var api_response []util.ApiResponse
	json.Unmarshal(response, &api_response)
	return api_response
}

func (l *Lights) SetLightState(light_id int, state State) []util.ApiResponse {
	url := fmt.Sprintf(set_light_state_url, l.Hostname, l.Username, light_id)
	state_json, _ := json.Marshal(&state)
	data := string(state_json)
	response := util.HttpPut(url, data, "application/json")
	var api_response []util.ApiResponse
	json.Unmarshal(response, &api_response)
	return api_response
}

func (l *Lights) GetAllLights() []Light {
	url := fmt.Sprintf(get_all_lights_url, l.Hostname, l.Username)
	response := util.HttpGet(url)
	lights_map := map[string]Light{}
	json.Unmarshal(response, &lights_map)
	lights := make([]Light, 0, len(lights_map))
	for light_id, light := range lights_map {
		light.Id, _ = strconv.Atoi(light_id)
		lights = append(lights, light)
	}
	return lights
}
