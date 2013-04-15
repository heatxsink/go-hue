package hue

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

var (
	user_agent          = "hue/1.0.2 CFNetwork/609.1.4 Darwin/13.0.0"
	create_username_url = "http://%s/api"
	delete_username_url = "http://%s/api/%s/config/whitelist/%s"
	configuration_url   = "http://%s/api/%s/config"
	portal_url          = "http://www.meethue.com/api/nupnp"
)

type User struct {
	Username   string `json:"username"`
	DeviceType string `json:"devicetype"`
}

type Portal struct {
	Id string `json:"id"`
	InternalIpAddress string `json:"internalipaddress"`
	MacAddress string `json:"macaddress"`
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

type Hue struct {
	Hostname string
}

func NewHue(hostname string) *Hue {
	hue := Hue{hostname}
	return &hue
}

func (hue *Hue) CreateUser(username string, device_type string) []map[string]map[string]interface{} {
	username_md5 := create_md5_hash(username)
	api_data := User{username_md5, device_type}
	json_data, _ := json.Marshal(api_data)
	url := fmt.Sprintf(create_username_url, hue.Hostname)
	response := http_post(url, string(json_data), "application/json")
	var api_response []map[string]map[string]interface{}
	json.Unmarshal(response, &api_response)
	return api_response
}

func (hue *Hue) DeleteUser(username string) []map[string]interface{} {
	url := fmt.Sprintf(delete_username_url, hue.Hostname, username, username)
	response := http_delete(url)
	var api_response []map[string]interface{}
	json.Unmarshal(response, &api_response)
	return api_response
}

func (hue *Hue) GetConfiguration(username string) Configuration {
	url := fmt.Sprintf(configuration_url, hue.Hostname, username)
	response := http_get(url)
	var api_response Configuration
	json.Unmarshal(response, &api_response)
	return api_response
}

func GetPortal() []Portal {
	response := http_get(portal_url)
	var api_response []Portal
	json.Unmarshal(response, &api_response)
	return api_response
}

func Discover() string {
	service := "239.255.255.250:1900"
	mac_address, err := net.ResolveUDPAddr("udp4", service)
	check_error(err)
	send, err := net.DialUDP("udp4", nil, mac_address)
	check_error(err)
	defer send.Close()
	// Send SSDP Message
	ssdp_discovery_message := []byte("M-SEARCH * HTTP/1.1\r\nHOST: 239.255.255.250:1900\r\nMAN: ssdp:discover\r\nMX: 10\r\nST: \"ssdp:all\"\r\n\r\n")
	_, error := send.Write(ssdp_discovery_message)
	check_error(error)
	fmt.Println("Searching for Philip Hue Hub (Could take up to 30 secs)...")
	// Listen for SSDP/HTTP NOTIFY over UDP
	listen, err := net.ListenMulticastUDP("udp4", nil, mac_address)
	check_error(err)
	defer listen.Close()
	description_url := ""
	for {
		b := make([]byte, 256)
		_, _, err := listen.ReadFromUDP(b)
		check_error(err)
		payload_message := string(b)
		headers := strings.Split(payload_message, "\r\n")
		for _, header := range headers {
			datum := strings.Split(header, ": ")
			if len(datum) > 1 {
				if datum[0] == "LOCATION" {
					if strings.Contains(datum[1], "description.xml") {
						description_url = datum[1]
						break
					}
				}
			}
		}
		if strings.Contains(description_url, "description.xml") {
			break
		}
	}
	u, err := url.Parse(description_url)
	check_error(err)
	hostname := ""
	if strings.Contains(u.Host, ":") {
		h := strings.Split(u.Host, ":")
		hostname = h[0]
	} else {
		hostname = u.Host
	}
	fmt.Printf("Found Hub at %s\n", hostname)
	return hostname
}

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func create_md5_hash(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	md5_hash := h.Sum(nil)
	return fmt.Sprintf("%x", md5_hash)
}

func http_post(url string, data string, content_type string) []byte {
	post_body := strings.NewReader(data)
	request, err := http.NewRequest("POST", url, post_body)
	check_error(err)
	request.Header.Set("Content-Type", content_type)
	client := http.Client{}
	response, err := client.Do(request)
	check_error(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	check_error(err)
	return contents
}

func http_get(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	check_error(err)
	client := http.Client{}
	response, err := client.Do(request)
	check_error(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	check_error(err)
	return contents
}

func http_delete(url string) []byte {
	request, err := http.NewRequest("DELETE", url, nil)
	check_error(err)
	client := http.Client{}
	response, err := client.Do(request)
	check_error(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	check_error(err)
	return contents
}
