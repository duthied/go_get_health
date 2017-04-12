// mogo_health.go

package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
)

var url = "" // some url

type Resource struct {
	url string
	errCount int
	body string
}

type Message struct {
	Name string `json:"name"`
	Status string `json:"status"`
	ResponseTime int `json:"response_time_ms"`
	Error string `json:"error"`
}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	jsonData := []Message{}
	err = json.Unmarshal([]byte(body), &jsonData)
	if err != nil {
		panic(err)
	}

	// test struct data
	for j := range jsonData {
		var buffer bytes.Buffer
		buffer.WriteString(jsonData[j].Status)
		if len(jsonData[j].Error) != 0 {
			buffer.WriteString("error:")
			buffer.WriteString(jsonData[j].Error)
		}

		fmt.Println("{", jsonData[j].ResponseTime, "}", jsonData[j].Name, "-", buffer.String())
	}
	
}
