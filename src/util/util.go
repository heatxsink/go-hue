package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"os"
)

var (
	user_agent = "hue/1.0.2 CFNetwork/609.1.4 Darwin/13.0.0"
)

type ApiResponse struct {
	Success map[string]interface{} `json:"success,omitempty"`
	Error ApiResponseError `json:"error,omitempty"`
}

type ApiResponseError struct {
	Type uint   `json:"type"`
	Address string `json:"address"`
	Description string `json:"description"`
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FileExists(path string) (bool, error) {
	_, error := os.Stat(path)
	if error == nil { return true, nil }
	if os.IsNotExist(error) { return false, nil }
	return false, error
}

func CreateMd5Hash(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	md5_hash := h.Sum(nil)
	return fmt.Sprintf("%x", md5_hash)
}

func HttpPut(url string, data string, content_type string) []byte {
	post_body := strings.NewReader(data)
	request, err := http.NewRequest("PUT", url, post_body)
	CheckError(err)
	request.Header.Set("Content-Type", content_type)
	client := http.Client{}
	response, err := client.Do(request)
	CheckError(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	return contents
}

func HttpPost(url string, data string, content_type string) []byte {
	post_body := strings.NewReader(data)
	request, err := http.NewRequest("POST", url, post_body)
	CheckError(err)
	request.Header.Set("Content-Type", content_type)
	client := http.Client{}
	response, err := client.Do(request)
	CheckError(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	return contents
}

func HttpGet(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	CheckError(err)
	client := http.Client{}
	response, err := client.Do(request)
	CheckError(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	return contents
}

func HttpDelete(url string) []byte {
	request, err := http.NewRequest("DELETE", url, nil)
	CheckError(err)
	client := http.Client{}
	response, err := client.Do(request)
	CheckError(err)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	return contents
}