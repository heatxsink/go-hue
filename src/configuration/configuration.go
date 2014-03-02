package configuration

import (
	"encoding/json"
	"fmt"
	"../util"
)

var (
	configuration_url   = "http://%s/api/%s/config"
	full_state_url = "http://%s/api/%s"
	create_username_url = "http://%s/api"
	delete_username_url = "http://%s/api/%s/config/whitelist/%s"
)

type User struct {
	Username   string `json:"username"`
	DeviceType string `json:"devicetype"`
}

type SwUpdate struct {
	UpdateState int `json:updatestate`
	Url string `json:url`
	Text string `json:text`
	Notify bool `json:notify`
}

type Whitelist struct {
	LastUseDate string `json:"last use date"`
	CreateDate string `json:"create date"`
	Name string `json:name`
}

type Configuration struct {
	Hostname string `json:"go_hue_hostname,omitempty"`
	ProxyPort int `json:proxyport`
	Utc string `json:utc`
	Name string `json:name`
	SwUpdate SwUpdate `json:swupdate`
	Whitelist map[string]Whitelist `json:whitelist`
	SwVersion string `json:swversion`
	ProxyAddress string `json:proxyaddress`
	Mac string `json:mac`
	LinkButton bool `json:linkbutton`
	IpAddress string `json:ipaddress`
	NetMask string `json:netmask`
	Gateway string `json:gateway`
	Dhcp bool `json:dhcp`
}

func NewConfiguration(hostname string) *Configuration {
	configuration := Configuration{}
	configuration.Hostname = hostname
	return &configuration
}

func (cc *Configuration) CreateUser(username string, device_type string) []util.ApiResponse {
	username_md5 := util.CreateMd5Hash(username)
	api_data := User{username_md5, device_type}
	json_data, _ := json.Marshal(api_data)
	url := fmt.Sprintf(create_username_url, cc.Hostname)
	response := util.HttpPost(url, string(json_data), "application/json")
	var api_response []util.ApiResponse
	json.Unmarshal(response, &api_response)
	return api_response
}

func (cc *Configuration) DeleteUser(username string, delete_this_username string) []util.ApiResponse {
	url := fmt.Sprintf(delete_username_url, cc.Hostname, username, delete_this_username)
	response := util.HttpDelete(url)
	var api_response []util.ApiResponse
	json.Unmarshal(response, &api_response)
	return api_response
}

func (cc *Configuration) GetFullState(username string) string {
	url := fmt.Sprintf(full_state_url, cc.Hostname, username)
	response := util.HttpGet(url)
	//var api_response Configuration
	//json.Unmarshal(response, &api_response)
	//return api_response
	return string(response)
}

func (cc *Configuration) GetConfiguration(username string) Configuration {
	url := fmt.Sprintf(configuration_url, cc.Hostname, username)
	response := util.HttpGet(url)
	var api_response Configuration
	json.Unmarshal(response, &api_response)
	return api_response
}
