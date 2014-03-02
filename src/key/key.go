package key

import (
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"strings"
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
	Username string
}

func FileExists(path string) (bool, error) {
	_, error := os.Stat(path)
	if error == nil { return true, nil }
	if os.IsNotExist(error) { return false, nil }
	return false, error
}

func LoadFromFile(filename string) *Key {
	exists, _ := FileExists(filename)
	if !exists {
		message := fmt.Sprintf("Filename %s does not exist.", filename)
		log.Fatal(message)
	}
	content, error_msg := ioutil.ReadFile(filename)
	if error_msg != nil {
		log.Fatal(error_msg)
	}
	var key Key
	tokens := strings.Split(string(content), "\n")
	username_api_key := strings.TrimRight(tokens[0], "\n")
	key.Username = username_api_key
	return &key
}

func (k *Key) Dump() {
	fmt.Println("\tUsername:      ", k.Username)

}
