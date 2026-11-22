package melipayamak

import (
	"strconv"
)

func SendSMSAsync(username string, password string, to string, from string, text string, isFlash bool) <-chan Result[string] {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"to":       to,
		"from":     from,
		"text":     text,
		"isFlash":  strconv.FormatBool(isFlash),
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "SendSMS")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendByBaseNumberAsync(username string, password string, text string, to string, bodyId int64) <-chan Result[string] {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"text":     text,
		"to":       to,
		"bodyId":   strconv.FormatInt(bodyId, 10),
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "BaseServiceNumber")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetDeliveries2Async(username string, password string, recID int64) <-chan Result[string] {
	jsonData := map[string]string{
		"username": username,
		"password": password,
		"recID":    strconv.FormatInt(recID, 10),
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetDeliveries2")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetMessagesAsync(username string, password string, location int, from string, index string, count bool) <-chan Result[string] {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
		"Location": strconv.Itoa(location),
		"From":     from,
		"Index":    index,
		"Count":    strconv.FormatBool(count),
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetMessages")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetCreditAsync(username string, password string) <-chan Result[string] {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetCredit")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetBasePriceAsync(username string, password string) <-chan Result[string] {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetBasePrice")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetUserNumbersAsync(username string, password string) <-chan Result[string] {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetUserNumbers")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendMultipleSMSAsync(username string, password string, from string, to []string, text []string) <-chan Result[string] {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
		"From":     from,
		"to":       RESTArray(to),
		"text":     RESTArray(text),
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "SendMultipleSMS")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}
