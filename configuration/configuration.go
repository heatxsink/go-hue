package configuration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/heatxsink/go-hue/hue"
	"io/ioutil"
	"net/http"
)

var (
	configurationURL  = "http://%s/api/%s/config"
	fullStateURL      = "http://%s/api/%s"
	createUsernameURL = "http://%s/api"
	deleteUsernameURL = "http://%s/api/%s/config/whitelist/%s"
)

type User struct {
	Username   string `json:"username,omitempty"`
	DeviceType string `json:"devicetype"`
}

type SWUpdate struct {
	UpdateState int    `json:updatestate`
	URL         string `json:url`
	Text        string `json:text`
	Notify      bool   `json:notify`
}

type Whitelist struct {
	LastUseDate string `json:"last use date"`
	CreateDate  string `json:"create date"`
	Name        string `json:name`
}

type Configuration struct {
	Hostname     string               `json:"go_hue_hostname,omitempty"`
	ProxyPort    int                  `json:proxyport`
	UTC          string               `json:utc`
	Name         string               `json:name`
	SWUpdate     SWUpdate             `json:swupdate`
	Whitelist    map[string]Whitelist `json:whitelist`
	SWVersion    string               `json:swversion`
	ProxyAddress string               `json:proxyaddress`
	Mac          string               `json:mac`
	LinkButton   bool                 `json:linkbutton`
	IPAddress    string               `json:ipaddress`
	NetMask      string               `json:netmask`
	Gateway      string               `json:gateway`
	DHCP         bool                 `json:dhcp`
}

func New(hostname string) *Configuration {
	return &Configuration{
		Hostname: hostname,
	}
}

func (c *Configuration) CreateUser(applicationName string, deviceName string) ([]hue.ApiResponse, error) {
	data := fmt.Sprintf("%s#%s", applicationName, deviceName)
	user := User{
		DeviceType: data,
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(createUsernameURL, c.Hostname)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
	if apiResponse[0].Error != nil {
		return nil, errors.New(apiResponse[0].Error.Description)
	}
	return apiResponse, err
}

func (c *Configuration) DeleteUser(username string, usernameToBeDeleted string) ([]hue.ApiResponse, error) {
	url := fmt.Sprintf(deleteUsernameURL, c.Hostname, username, usernameToBeDeleted)
	request, err := http.NewRequest("DELETE", url, nil)
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
	var apiResponse []hue.ApiResponse
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, err
}

func (c *Configuration) GetFullState(username string) (string, error) {
	url := fmt.Sprintf(fullStateURL, c.Hostname, username)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(contents), err
}

func (c *Configuration) GetConfiguration(username string) (Configuration, error) {
	var configuration Configuration
	url := fmt.Sprintf(configurationURL, c.Hostname, username)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return configuration, err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return configuration, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return configuration, err
	}
	if response.StatusCode != 200 {
		fmt.Println(string(contents))
	}
	err = json.Unmarshal(contents, &configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, err
}

func (c *Configuration) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Name:         %s\n", c.Name))
	buffer.WriteString(fmt.Sprintf("UTC:          %s\n", c.UTC))
	buffer.WriteString(fmt.Sprintf("SWVersion:    %s\n", c.SWVersion))
	buffer.WriteString(fmt.Sprintf("ProxyAddress: %s\n", c.ProxyAddress))
	buffer.WriteString(fmt.Sprintf("ProxyPort:    %d\n", c.ProxyPort))
	buffer.WriteString(fmt.Sprintf("Mac:          %s\n", c.Mac))
	buffer.WriteString(fmt.Sprintf("LinkButton:   %t\n", c.LinkButton))
	buffer.WriteString(fmt.Sprintf("IPAddress:    %s\n", c.IPAddress))
	buffer.WriteString(fmt.Sprintf("NetMask:      %s\n", c.NetMask))
	buffer.WriteString(fmt.Sprintf("Gateway:      %s\n", c.Gateway))
	buffer.WriteString(fmt.Sprintf("DHCP:         %t\n", c.DHCP))
	buffer.WriteString(fmt.Sprintf("SWUpdate:\n"))
	buffer.WriteString(fmt.Sprintf("\tUpdateState: %d\n", c.SWUpdate.UpdateState))
	buffer.WriteString(fmt.Sprintf("\tURL:         %s\n", c.SWUpdate.URL))
	buffer.WriteString(fmt.Sprintf("\tText:        %s\n", c.SWUpdate.Text))
	buffer.WriteString(fmt.Sprintf("\tNotify:      %t\n", c.SWUpdate.Notify))
	buffer.WriteString(fmt.Sprintf("Whitelist:\n"))
	for key := range c.Whitelist {
		buffer.WriteString(fmt.Sprintf("\tKey: %s\n", key))
		buffer.WriteString(fmt.Sprintf("\t\tLastUseDate: %s\n", c.Whitelist[key].LastUseDate))
		buffer.WriteString(fmt.Sprintf("\t\tCreateDate:  %s\n", c.Whitelist[key].CreateDate))
		buffer.WriteString(fmt.Sprintf("\t\tName:        %s\n", c.Whitelist[key].Name))
	}
	return buffer.String()
}
