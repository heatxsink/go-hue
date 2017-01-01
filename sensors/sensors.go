package sensors

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
	getAllSensorsURL = "http://%s/api/%s/sensors"
	getSensorURL = "http://%s/api/%s/sensors/%d"
)

type Sensors struct {
	Hostname string
	Username string
}

type Sensor struct {
	ID               int    `json:"id,omitempty"`
	Name             string `json:"name"`
	State            State  `json:"state,omitempty"`
	Config           Config `json:"config,omitempty"`
	Type             string `json:"type,omitempty"`
	ModelID          string `json:"modelid,omitempty"`
	SWVersion        string `json:"swversion,omitempty"`
	ManufacturerName string `json:"manufacturername,omitempty"`
	UniqueID         string `json:"uniqueid,omitempty"`
}

type Config struct {
	On            bool     `json:"on"`
	Long          string   `json:"long,omitempty"`
	Lat           string   `json:"lat,omitempty"`
	SunriseOffset int16    `json:"sunriseoffset,omitempty"`
	SunsetOffset  int16    `json:"sunsetoffset,omitempty"`
}

type State struct {
	Daylight    bool      `json:"daylight,omitempty"`
	LastUpdated string    `json:"lastupdated,omitempty"`
	ButtonEvent int16     `json:"buttonevent,omitempty"`
}

func New(hostname string, username string) *Sensors {
	return &Sensors{
		Hostname: hostname,
		Username: username,
	}
}

func (l *Sensors) GetSensor(sensorID int) (Sensor, error) {
	var ll Sensor
	url := fmt.Sprintf(getSensorURL, l.Hostname, l.Username, sensorID)
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
	ll.ID = sensorID
	return ll, err
}

func (l *Sensors) UpdateSensor(sensorID int, sensorName string) ([]hue.ApiResponse, error) {
	url := fmt.Sprintf(getSensorURL, l.Hostname, l.Username, sensorID)
	data := fmt.Sprintf("{\"name\": \"%s\"}", sensorName)
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

func (l *Sensors) GetAllSensors() ([]Sensor, error) {
	url := fmt.Sprintf(getAllSensorsURL, l.Hostname, l.Username)
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
	sensorsMap := map[string]Sensor{}
	err = json.Unmarshal(contents, &sensorsMap)
	if err != nil {
		return nil, err
	}
	sensors := make([]Sensor, 0, len(sensorsMap))
	for sensorID, sensor := range sensorsMap {
		sensor.ID, _ = strconv.Atoi(sensorID)
		sensors = append(sensors, sensor)
	}
	return sensors, err
}

func (l *Sensor) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Id:              %d\n", l.ID))
	buffer.WriteString(fmt.Sprintf("Name:            %s\n", l.Name))
	buffer.WriteString(fmt.Sprintf("Type:            %s\n", l.Type))
	buffer.WriteString(fmt.Sprintf("ModelId:         %s\n", l.ModelID))
	buffer.WriteString(fmt.Sprintf("SwVersion:       %s\n", l.SWVersion))
	buffer.WriteString(fmt.Sprint("State:\n"))
	buffer.WriteString(l.State.String())
	return buffer.String()
}

func (s *State) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("ButtonEvent:     %d\n", s.ButtonEvent))
	buffer.WriteString(fmt.Sprintf("Daylight:        %t\n", s.Daylight))
	buffer.WriteString(fmt.Sprintf("LastUpdated:        %s\n", s.LastUpdated))
	return buffer.String()
}
