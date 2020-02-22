package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string) (response []byte) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}
	response, _ = ioutil.ReadAll(resp.Body)
	return
}

func Post(url string, data interface{}, contentType string) (content []byte) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	content, _ = ioutil.ReadAll(resp.Body)
	return
}
