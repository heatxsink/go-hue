package schedules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	getAllSensorsURL = "http://%s/api/%s/schedules"
	getSensorURL     = "http://%s/api/%s/schedule/%d"
)

type Schedules struct {
	Hostname string
	Username string
}

type Schedule struct {
	ID             int         `json:"id,omitempty"`
	Name           string      `json:"name"`
	Description    string      `json:"description,omitempty"`
	Status         string      `json:"status,omitempty"`
	Conditions     []Condition `json:"conditions"`
	Command        Command     `json:"command"`
	Owner          string      `json:"owner,omitempty"`
	TimesTriggered int16       `json:"timestriggered,omitempty"`
	LastTriggered  string      `json:"lasttriggered,omitempty"`
	StartTime      string      `json:"starttime,omitempty"`
	Time           string      `json:"time,omitempty"`
	AutoDelete     bool        `json:"autodelete,omitempty"`
}

type Condition struct {
	Address  string `json:"address,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Command struct {
	Address string      `json:"address,omitempty"`
	Method  string      `json:"method,omitempty"`
	Body    interface{} `json:"value,omitempty"`
}

func New(hostname string, username string) *Schedules {
	return &Schedules{
		Hostname: hostname,
		Username: username,
	}
}

func (l *Schedules) GetSchedule(ruleID int) (Schedule, error) {
	var ll Schedule
	url := fmt.Sprintf(getSensorURL, l.Hostname, l.Username, ruleID)
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
	ll.ID = ruleID
	return ll, err
}

func (l *Schedules) GetAllSchedules() ([]Schedule, error) {
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
	sensorsMap := map[string]Schedule{}
	err = json.Unmarshal(contents, &sensorsMap)
	if err != nil {
		return nil, err
	}
	sensors := make([]Schedule, 0, len(sensorsMap))
	for sensorID, sensor := range sensorsMap {
		sensor.ID, _ = strconv.Atoi(sensorID)
		sensors = append(sensors, sensor)
	}
	return sensors, err
}

func (l *Schedule) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Id:              %d\n", l.ID))
	buffer.WriteString(fmt.Sprintf("Name:            %s\n", l.Name))
	buffer.WriteString(fmt.Sprintf("Owner:           %s\n", l.Owner))
	buffer.WriteString(fmt.Sprintf("Time:            %s\n", l.Time))
	buffer.WriteString(fmt.Sprintf("StartTime:       %s\n", l.StartTime))
	buffer.WriteString(fmt.Sprintf("LastTriggered:   %s\n", l.LastTriggered))
	buffer.WriteString(fmt.Sprintf("Status:          %s\n", l.Status))
	buffer.WriteString(fmt.Sprintf("Command:\n%s\n", l.Command.String()))
	return buffer.String()
}

func (c *Command) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Address:         %s\n", c.Address))
	buffer.WriteString(fmt.Sprintf("Method:          %s\n", c.Method))
	buffer.WriteString(fmt.Sprintf("Body:            TODO\n"))
	return buffer.String()
}
