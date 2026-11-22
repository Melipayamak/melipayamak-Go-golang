package melipayamak

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	RESTSendUrl = "https://rest.payamak-panel.com/api/SendSMS/"
)

// func invoke(jsonData map[string]any, endpoint, op string) (string, error) {
// 	jsonValue, err := json.Marshal(jsonData)
// 	if err != nil {
// 		return "", err
// 	}

// 	response, err := http.Post(endpoint+op, "application/json", bytes.NewBuffer(jsonValue))
// 	if err != nil {
// 		return "", err
// 	}

// 	defer response.Body.Close()

// 	data, err := io.ReadAll(response.Body)

// 	if err != nil {
// 		return "", err
// 	}

// 	if response.StatusCode < 200 || response.StatusCode >= 300 {
// 		return string(data), fmt.Errorf("HTTP error: %s", response.Status)
// 	}

// 	return string(data), nil
// }

func invoke(jsonData map[string]any, endpoint, op string) (string, error) {
	url := endpoint + op

	log.Println("========== MELIPAYAMAK DEBUG ==========")
	log.Println("URL:", url)

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	log.Println("Request:", string(jsonValue))
	log.Println("Sending request...")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Post(
		url,
		"application/json; charset=utf-8",
		bytes.NewBuffer(jsonValue),
	)

	if err != nil {
		log.Println("REQUEST ERROR:", err)
		return "", err
	}

	defer response.Body.Close()

	log.Println("HTTP STATUS:", response.StatusCode)

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("RESPONSE READ ERROR:", err)
		return "", err
	}

	log.Println("RESPONSE BODY:", string(data))
	log.Println("=======================================")

	return string(data), nil
}

// func RESTArray[T any](values []T) string {
// 	data, err := json.Marshal(values)
// 	if err != nil {
// 		return "[]"
// 	}
// 	return string(data)
// }

type Result[T any] struct {
	Value T
	Err   error
}
