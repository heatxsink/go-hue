package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	portalURL = "https://discovery.meethue.com/"
)

type Portal struct {
	ID                string `json:"id"`
	InternalIPAddress string `json:"internalipaddress"`
	MacAddress        string `json:"macaddress"`
}

func GetPortal() ([]Portal, error) {
	request, err := http.NewRequest("GET", portalURL, nil)
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
	var apiResponse []Portal
	err = json.Unmarshal(contents, &apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, err
}

func (p *Portal) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("ID:                 %s\n", p.ID))
	buffer.WriteString(fmt.Sprintf("InternalIPAddress:  %s\n", p.InternalIPAddress))
	buffer.WriteString(fmt.Sprintf("MacAddress:         %s\n", p.MacAddress))
	return buffer.String()
}
