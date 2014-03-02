package portal

import (
	"testing"
	"fmt"
)

func TestGetPortal(t *testing.T) {
	portal := GetPortal()
	fmt.Println("portal.GetPortal()")
	for i := range portal {
		fmt.Printf("\tId:                  %s\n", portal[i].Id)
		fmt.Printf("\tInternal Ip Address: %s\n", portal[i].InternalIpAddress)
		fmt.Printf("\tMac Address:         %s\n", portal[i].MacAddress)
		t.Log(portal[i].InternalIpAddress)
	}
}
