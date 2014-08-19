package key

import (
	"io/ioutil"
	"log"
	"fmt"
	"encoding/json"
	"github.com/heatxsink/go-hue/src/util"
)

var (
	instantiated *Key = nil
)

func New(filename string) *Key {
	if instantiated == nil {
		instantiated = LoadFromFile(filename)
	}
	return instantiated
}

type Key struct {
	Username string `json:"username"`
}

func LoadFromFile(filename string) *Key {
	exists, _ := util.FileExists(filename)
	if !exists {
		message := fmt.Sprintf("Filename %s does not exist.", filename)
		log.Fatal(message)
	}
	content, error_msg := ioutil.ReadFile(filename)
	if error_msg != nil {
		log.Fatal(error_msg)
	}
	var key Key
	json.Unmarshal(content, &key)
	return &key
}

func (k *Key) Dump() {
	fmt.Println("\tUsername:      ", k.Username)

}
