package melipayamak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	RESTSendUrl = "https://rest.payamak-panel.com/api/SendSMS/"
)

func invoke(jsonData map[string]any, endpoint, op string) (string, error) {
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	response, err := http.Post(endpoint+op, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return string(data), fmt.Errorf("HTTP error: %s", response.Status)
	}

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
