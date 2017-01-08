package rules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/heatxsink/go-hue/hue"
)

var (
	getAllSensorsURL = "http://%s/api/%s/rules"
	getSensorURL     = "http://%s/api/%s/rule/%d"
	deleteRuleURL    = getSensorURL
	createRuleURL    = getAllSensorsURL
)

type Rules struct {
	Hostname string
	Username string
}

type Rule struct {
	ID             int         `json:"id,omitempty"`
	Name           string      `json:"name"`
	Status         string      `json:"status,omitempty"`
	Conditions     []Condition `json:"conditions"`
	Actions        []Action    `json:"actions"`
	Owner          string      `json:"owner,omitempty"`
	TimesTriggered int16       `json:"timestriggered,omitempty"`
	LastTriggered  string      `json:"lasttriggered,omitempty"`
	Created        string      `json:"created,omitempty"`
}

type Condition struct {
	Address  string `json:"address,omitempty"`
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Action struct {
	Address string      `json:"address,omitempty"`
	Method  string      `json:"method,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}

func New(hostname string, username string) *Rules {
	return &Rules{
		Hostname: hostname,
		Username: username,
	}
}

func (l *Rules) GetRule(ruleID int) (Rule, error) {
	var ll Rule
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

func (l *Rules) CreateRule(rule Rule) ([]hue.ApiResponse, error) {
	url := fmt.Sprintf(createRuleURL, l.Hostname, l.Username)
	data, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
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

func (l *Rules) GetAllRules() ([]Rule, error) {
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
	sensorsMap := map[string]Rule{}
	err = json.Unmarshal(contents, &sensorsMap)
	if err != nil {
		return nil, err
	}
	sensors := make([]Rule, 0, len(sensorsMap))
	for sensorID, sensor := range sensorsMap {
		sensor.ID, _ = strconv.Atoi(sensorID)
		sensors = append(sensors, sensor)
	}
	return sensors, err
}

func (r *Rules) DeleteRule(ruleID int) error {
	var ll Rule
	url := fmt.Sprintf(deleteRuleURL, r.Hostname, r.Username, ruleID)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, &ll)
	if err != nil {
		return err
	}
	ll.ID = ruleID
	return err
}

func (l *Rule) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Id:              %d\n", l.ID))
	buffer.WriteString(fmt.Sprintf("Name:            %s\n", l.Name))
	buffer.WriteString(fmt.Sprintf("Owner:           %s\n", l.Owner))
	buffer.WriteString(fmt.Sprintf("Created:         %s\n", l.Created))
	buffer.WriteString(fmt.Sprintf("LastTriggered:   %s\n", l.LastTriggered))
	buffer.WriteString(fmt.Sprintf("Status:          %s\n", l.Status))
	buffer.WriteString(fmt.Sprintf("Conditions:\n"))
	for i := range l.Conditions {
		buffer.WriteString(fmt.Sprintf("  Address:          %s\n", l.Conditions[i].Address))
		buffer.WriteString(fmt.Sprintf("  Operator:         %s\n", l.Conditions[i].Operator))
		buffer.WriteString(fmt.Sprintf("  Value:            %s\n", l.Conditions[i].Value))
	}
	buffer.WriteString(fmt.Sprintf("Actions:\n"))
	for i := range l.Actions {
		buffer.WriteString(fmt.Sprintf("  Address:          %s\n", l.Actions[i].Address))
		buffer.WriteString(fmt.Sprintf("  Method:           %s\n", l.Actions[i].Method))
		if l.Actions[i].Body != nil {
			repr, err := json.Marshal(l.Actions[i].Body)
			if err == nil {
				buffer.WriteString(fmt.Sprintf("  Body:             %s\n", repr))
			}
		}
	}
	return buffer.String()
}
