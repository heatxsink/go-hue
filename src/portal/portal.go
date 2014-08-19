package portal

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"encoding/json"
	"github.com/heatxsink/go-hue/src/util"
)

var (
	portal_url = "http://www.meethue.com/api/nupnp"
)

type Portal struct {
	Id string `json:"id"`
	InternalIpAddress string `json:"internalipaddress"`
	MacAddress string `json:"macaddress"`
}

func GetPortal() []Portal {
	response := util.HttpGet(portal_url)
	var api_response []Portal
	json.Unmarshal(response, &api_response)
	return api_response
}

func Discover() string {
	service := "239.255.255.250:1900"
	mac_address, err := net.ResolveUDPAddr("udp4", service)
	util.CheckError(err)
	send, err := net.DialUDP("udp4", nil, mac_address)
	util.CheckError(err)
	defer send.Close()
	// Send SSDP Message
	ssdp_discovery_message := []byte("M-SEARCH * HTTP/1.1\r\nHOST: 239.255.255.250:1900\r\nMAN: ssdp:discover\r\nMX: 10\r\nST: \"ssdp:all\"\r\n\r\n")
	_, error := send.Write(ssdp_discovery_message)
	util.CheckError(error)
	fmt.Println("Searching for Philip Hue Hub (Could take up to 30 secs)...")
	// Listen for SSDP/HTTP NOTIFY over UDP
	listen, err := net.ListenMulticastUDP("udp4", nil, mac_address)
	util.CheckError(err)
	defer listen.Close()
	description_url := ""
	for {
		b := make([]byte, 256)
		_, _, err := listen.ReadFromUDP(b)
		util.CheckError(err)
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
	util.CheckError(err)
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
