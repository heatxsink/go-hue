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
)

type Api struct {
	Username   string `json:"username"`
	DeviceType string `json:"devicetype"`
}

type Hue struct {
	Hostname string
}

func NewHue(hostname string) *Hue {
	hue := Hue{hostname}
	return &hue
}

func (hue *Hue) CreateUsername(username string, device_type string) []map[string]map[string]interface{} {
	username_md5 := create_md5_hash(username)
	api_data := Api{username_md5, device_type}
	json_data, _ := json.Marshal(api_data)
	url := fmt.Sprintf(create_username_url, hue.Hostname)
	response := http_post(url, string(json_data), "application/json")
	var api_response []map[string]map[string]interface{}
	json.Unmarshal(response, &api_response)
	return api_response
}

func (hue *Hue) DeleteUsername(username string) []map[string]interface{} {
	url := fmt.Sprintf(delete_username_url, hue.Hostname, username, username)
	response := http_delete(url)
	var api_response []map[string]interface{}
	json.Unmarshal(response, &api_response)
	return api_response
}

func (hue *Hue) GetConfig() {
	//TODO: To be implemented very soon.
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
	fmt.Println("Searching for Philip Hue Hub ...")
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
