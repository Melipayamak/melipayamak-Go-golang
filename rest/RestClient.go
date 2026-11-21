package melipayamak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func invoke(jsonData map[string]string, op string) {

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://rest.payamak-panel.com/api/SendSMS/"+op, "application/json", bytes.NewBuffer(jsonValue))
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

func SendSMS(username string, password string, to string, from string, text string, isFlash bool) {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"to":       to,
		"from":     from,
		"text":     text,
		"isFlash":  strconv.FormatBool(isFlash),
	}
	invoke(jsonData, "SendSMS")
}

func SendByBaseNumber(username string, password string, text string, to string, bodyId int64) {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"text":     text,
		"to":       to,
		"bodyId":   strconv.FormatInt(bodyId, 10),
	}
	invoke(jsonData, "BaseServiceNumber")
}

func GetDeliveries2(username string, password string, recID int64) {
	jsonData := map[string]string{
		"username": username,
		"password": password,
		"recID":    strconv.FormatInt(recID, 10),
	}
	invoke(jsonData, "GetDeliveries2")
}

func GetMessages(username string, password string, location int, from string, index string, count bool) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
		"Location": strconv.Itoa(location),
		"From":     from,
		"Index":    index,
		"Count":    strconv.FormatBool(count),
	}
	invoke(jsonData, "GetMessages")
}

func GetCredit(username string, password string) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	invoke(jsonData, "GetCredit")
}

func GetBasePrice(username string, password string) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	invoke(jsonData, "GetBasePrice")
}

func GetUserNumbers(username string, password string) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	invoke(jsonData, "GetUserNumbers")
}
