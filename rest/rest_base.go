package melipayamak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	RESTSendUrl = "https://rest.payamak-panel.com/api/SendSMS/"
)

func invoke(jsonData map[string]string, endpoint, op string) {

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(endpoint+op, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("HTTP request failed: %s\n", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	fmt.Println(string(data))
}

func RESTArray[T any](values []T) string {
	data, err := json.Marshal(values)
	if err != nil {
		return "[]"
	}
	return string(data)
}
