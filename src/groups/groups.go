package groups

import (
	"fmt"
	"github.com/heatxsink/go-hue/src/lights"
	"github.com/heatxsink/go-hue/src/util"
	"strconv"
	"encoding/json"
)

var (
	get_groups_url = "http://%s/api/%s/groups"
	get_group_url = "http://%s/api/%s/groups/%d"
	set_group_url = "http://%s/api/%s/groups/%d"
	set_group_state_url = "http://%s/api/%s/groups/%d/action"
)

type Groups struct {
	Hostname string
	Username string
}

type Group struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name"`
	Lights []string `json:"lights,omitempty"`
}

type GroupState struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name"`
	Action lights.State `json:"action,omitempty"`
	Lights []string `json:"lights"`
}

func NewGroup(hostname string, username string) *Groups {
	group := Groups{Hostname: hostname, Username: username}
	return &group
}

func (g *Groups) GetGroup(group_number int) GroupState {
	url := fmt.Sprintf(get_group_url, g.Hostname, g.Username, group_number)
	response := util.HttpGet(url)
	var gg GroupState
	gg.Id = group_number
	json.Unmarshal(response, &gg)
	return gg
}

func (g *Groups) CreateGroup(group Group) []util.ApiResponse {
	url := fmt.Sprintf(get_groups_url, g.Hostname, g.Username)
	group_json, _ := json.Marshal(&group)
	data := string(group_json)
	response := util.HttpPost(url, data, "application/json")
	var api_response []util.ApiResponse
	json.Unmarshal(response, &api_response)
	return api_response
}

func (g *Groups) SetGroup(group_number int, group Group) []util.ApiResponse {
	url := fmt.Sprintf(set_group_url, g.Hostname, g.Username, group_number)
	group_json, _ := json.Marshal(&group)
	data := string(group_json)
	response := util.HttpPut(url, data, "application/json")
	var api_response []util.ApiResponse
	json.Unmarshal(response, &api_response)
	return api_response
}

func (g *Groups) SetGroupState(group_number int, group_state lights.State) []util.ApiResponse {
	url := fmt.Sprintf(set_group_state_url, g.Hostname, g.Username, group_number)
	group_state_json, _ := json.Marshal(&group_state)
	data := string(group_state_json)
	response := util.HttpPut(url, data, "application/json")
	var api_response []util.ApiResponse
	json.Unmarshal(response, &api_response)
	return api_response
}

func (g *Groups) GetGroups() []Group {
	url := fmt.Sprintf(get_groups_url, g.Hostname, g.Username)
	response := util.HttpGet(url)
	groups_map := map[string]Group{}
	json.Unmarshal(response, &groups_map)
	groups := make([]Group, 0, len(groups_map))
	for group_id, group := range groups_map {
		group.Id, _ = strconv.Atoi(group_id)
		groups = append(groups, group)
	}
	return groups
}