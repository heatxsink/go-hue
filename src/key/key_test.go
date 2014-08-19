package key

import (
	"testing"
	"fmt"
)

func TestNewAndDump(t *testing.T) {
	fmt.Println("key.New")
	filename := "../../conf/example.conf"
	zzz := New(filename)
	fmt.Println("configuration.Dump")
	zzz.Dump()
}
