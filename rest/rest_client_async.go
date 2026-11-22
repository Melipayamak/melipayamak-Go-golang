package melipayamak

import (
	"strconv"
)

func SendSMSAsync(username string, password string, to string, from string, text string, isFlash bool) {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"to":       to,
		"from":     from,
		"text":     text,
		"isFlash":  strconv.FormatBool(isFlash),
	}
	go invoke(jsonData, RESTSendUrl, "SendSMS")
}

func SendByBaseNumberAsync(username string, password string, text string, to string, bodyId int64) {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"text":     text,
		"to":       to,
		"bodyId":   strconv.FormatInt(bodyId, 10),
	}
	go invoke(jsonData, RESTSendUrl, "BaseServiceNumber")
}

func GetDeliveries2Async(username string, password string, recID int64) {
	jsonData := map[string]string{
		"username": username,
		"password": password,
		"recID":    strconv.FormatInt(recID, 10),
	}
	go invoke(jsonData, RESTSendUrl, "GetDeliveries2")
}

func GetMessagesAsync(username string, password string, location int, from string, index string, count bool) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
		"Location": strconv.Itoa(location),
		"From":     from,
		"Index":    index,
		"Count":    strconv.FormatBool(count),
	}
	go invoke(jsonData, RESTSendUrl, "GetMessages")
}

func GetCreditAsync(username string, password string) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	go invoke(jsonData, RESTSendUrl, "GetCredit")
}

func GetBasePriceAsync(username string, password string) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	go invoke(jsonData, RESTSendUrl, "GetBasePrice")
}

func GetUserNumbersAsync(username string, password string) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	go invoke(jsonData, RESTSendUrl, "GetUserNumbers")
}
