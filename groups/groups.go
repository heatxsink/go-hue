package groups

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/heatxsink/go-hue/hue"
	"github.com/heatxsink/go-hue/lights"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	setGroupURL      = "http://%s/api/%s/groups/%d"
	getGroupURL      = "http://%s/api/%s/groups/%d"
	deleteGroupURL   = "http://%s/api/%s/groups/%d"
	createGroupURL   = "http://%s/api/%s/groups"
	getAllGroupsURL  = "http://%s/api/%s/groups"
	setGroupStateURL = "http://%s/api/%s/groups/%d/action"
)

type Groups struct {
	Hostname string
	Username string
}

type Group struct {
	ID     int          `json:"id,omitempty"`
	Name   string       `json:"name"`
	Action lights.State `json:"action,omitempty"`
	Lights []string     `json:"lights,omitempty"`
}

func New(hostname string, username string) *Groups {
	return &Groups{
		Hostname: hostname,
		Username: username,
	}
}

func (g *Groups) GetAllGroups() ([]Group, error) {
	url := fmt.Sprintf(getAllGroupsURL, g.Hostname, g.Username)
	request, err := http.NewRequest("GET", url, nil)
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
	groupsMap := map[string]Group{}
	json.Unmarshal(contents, &groupsMap)
	groups := make([]Group, 0, len(groupsMap))
	for groupID, group := range groupsMap {
		group.ID, err = strconv.Atoi(groupID)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, err
}

func (g *Groups) CreateGroup(group Group) ([]hue.ApiResponse, error) {
	var apiResponse []hue.ApiResponse
	url := fmt.Sprintf(createGroupURL, g.Hostname, g.Username)
	jsonData, err := json.Marshal(&group)
	if err != nil {
		return apiResponse, err
	}
	postBody := strings.NewReader(string(jsonData))
	request, err := http.NewRequest("POST", url, postBody)
	if err != nil {
		return apiResponse, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResponse, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, err
}

func (g *Groups) GetGroup(groupID int) (Group, error) {
	var gg Group
	url := fmt.Sprintf(getGroupURL, g.Hostname, g.Username, groupID)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return gg, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return gg, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return gg, err
	}
	gg.ID = groupID
	err = json.Unmarshal(contents, &gg)
	if err != nil {
		return gg, err
	}
	return gg, err
}

func (g *Groups) SetGroup(groupID int, group Group) ([]hue.ApiResponse, error) {
	var apiResponse []hue.ApiResponse
	url := fmt.Sprintf(setGroupURL, g.Hostname, g.Username, groupID)
	jsonData, err := json.Marshal(&group)
	if err != nil {
		return apiResponse, err
	}
	body := strings.NewReader(string(jsonData))
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return apiResponse, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResponse, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}
	json.Unmarshal(contents, &apiResponse)
	return apiResponse, err
}

func (g *Groups) SetGroupState(groupID int, state lights.State) ([]hue.ApiResponse, error) {
	var apiResponse []hue.ApiResponse
	url := fmt.Sprintf(setGroupStateURL, g.Hostname, g.Username, groupID)
	jsonData, err := json.Marshal(&state)
	if err != nil {
		return apiResponse, err
	}
	body := strings.NewReader(string(jsonData))
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return apiResponse, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResponse, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, err
}

func (g *Groups) DeleteGroup(groupID int) ([]hue.ApiResponse, error) {
	// TODO: Implement
	return nil, nil
}

func (g *Group) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("ID:              %d\n", g.ID))
	buffer.WriteString(fmt.Sprintf("Name:            %s\n", g.Name))
	buffer.WriteString("Action:\n")
	buffer.WriteString(g.Action.String())
	buffer.WriteString("Lights:\n")
	for _, lightID := range g.Lights {
		buffer.WriteString(fmt.Sprintf("\t%s\n", lightID))
	}
	return buffer.String()
}
