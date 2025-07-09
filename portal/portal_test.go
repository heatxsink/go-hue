package portal

import (
	"fmt"
	"testing"
)

func TestGetPortal(t *testing.T) {
	p, err := GetPortal()
	if err != nil {
		t.Log("TestGetPortal Error: ", err)
		t.Fail()
	} else {
		for i := range p {
			fmt.Println(p[i].String())
		}
	}
}
