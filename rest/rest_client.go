package melipayamak

import (
	"strconv"
)

func SendSMS(username string, password string, to string, from string, text string, isFlash bool) (string, error) {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"to":       to,
		"from":     from,
		"text":     text,
		"isFlash":  strconv.FormatBool(isFlash),
	}
	return invoke(jsonData, RESTSendUrl, "SendSMS")
}

func SendByBaseNumber(username string, password string, text string, to string, bodyId int64) (string, error) {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"text":     text,
		"to":       to,
		"bodyId":   strconv.FormatInt(bodyId, 10),
	}
	return invoke(jsonData, RESTSendUrl, "BaseServiceNumber")
}

func GetDeliveries2(username string, password string, recID int64) (string, error) {
	jsonData := map[string]string{
		"username": username,
		"password": password,
		"recID":    strconv.FormatInt(recID, 10),
	}
	return invoke(jsonData, RESTSendUrl, "GetDeliveries2")
}

func GetMessages(username string, password string, location int, from string, index string, count bool) (string, error) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
		"Location": strconv.Itoa(location),
		"From":     from,
		"Index":    index,
		"Count":    strconv.FormatBool(count),
	}
	return invoke(jsonData, RESTSendUrl, "GetMessages")
}

func GetCredit(username string, password string) (string, error) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	return invoke(jsonData, RESTSendUrl, "GetCredit")
}

func GetBasePrice(username string, password string) (string, error) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	return invoke(jsonData, RESTSendUrl, "GetBasePrice")
}

func GetUserNumbers(username string, password string) (string, error) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
	}
	return invoke(jsonData, RESTSendUrl, "GetUserNumbers")
}

// v2
func SendOTP(username string, password string, to, from, code string) (string, error) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
		"to":       to,
		"From":     from,
		"code":     code,
	}
	return invoke(jsonData, RESTSendUrl, "SendOTP")
}

func SendMultipleSMS(username string, password string, from string, to []string, text []string) (string, error) {
	jsonData := map[string]string{
		"UserName": username,
		"PassWord": password,
		"From":     from,
		"to":       RESTArray(to),
		"text":     RESTArray(text),
	}
	return invoke(jsonData, RESTSendUrl, "SendMultipleSMS")
}

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }

// func GetMessages(username string, password string) {
// 	jsonData := map[string]string{
// 		"UserName": username,
// 		"PassWord": password,
// 		"Location": strconv.Itoa(location),
// 		"From":     from,
// 		"Index":    index,
// 		"Count":    strconv.FormatBool(count),
// 	}
// 	invoke(jsonData, RESTSendUrl, "GetMessages")
// }
