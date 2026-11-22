package melipayamak

import (
	"strconv"
)

func SendSMSAsync(username string, password string, to string, from string, text string, isFlash bool) <-chan Result[string] {

	jsonData := map[string]any{
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

	jsonData := map[string]any{
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
	jsonData := map[string]any{
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
	jsonData := map[string]any{
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
	jsonData := map[string]any{
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
	jsonData := map[string]any{
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
	jsonData := map[string]any{
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
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"From":     from,
		"to":       to,
		"text":     text,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "SendMultipleSMS")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendScheduleAsync(username string, password string, text, to, from, scheduleDate string, period int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":     username,
		"PassWord":     password,
		"from":         from,
		"to":           to,
		"text":         text,
		"scheduledate": scheduleDate,
		"period":       period,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "SendSchedule")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetScheduleAsync(username string, password string, scheduleId int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":   username,
		"PassWord":   password,
		"scheduleId": scheduleId,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetSchedule")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetDeliveriesAsync(username string, password string, recID string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recID":    recID,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetDeliveries")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetNewDeliveriesAsync(username string, password string, recID []string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recID":    recID,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetNewDeliveries")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetInboxCountAsync(username string, password string, isRead bool) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"isRead":   isRead,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetInboxCount")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetCredit2Async(username string, password string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetCredit2")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetSmsPriceAsync(username string, password string, irancellCount, mciCount int, from, text string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":      username,
		"PassWord":      password,
		"irancellCount": irancellCount,
		"mciCount":      mciCount,
		"from":          from,
		"text":          text,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "GetSmsPrice")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func BaseServiceNumberAsync(username string, password string, text, to string, bodyId int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"bodyId":   bodyId,
		"to":       to,
		"text":     text,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "BaseServiceNumber")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func BaseServiceNumberScheduleAsync(username string, password string, text, to string, bodyId int, scheduleDate string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":     username,
		"PassWord":     password,
		"bodyId":       bodyId,
		"to":           to,
		"text":         text,
		"scheduleDate": scheduleDate,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSendUrl, "BaseServiceNumberSchedule")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendSmartSmsAsync(username string, password string, to, text, from, fromSupportOne, fromSupportTwo string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":       username,
		"PassWord":       password,
		"from":           from,
		"to":             to,
		"text":           text,
		"fromSupportOne": fromSupportOne,
		"fromSupportTwo": fromSupportTwo,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSmartUrl, "Send")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendMultipleSmartSmsAsync(username string, password string, to, text []string, from, fromSupportOne, fromSupportTwo string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":       username,
		"PassWord":       password,
		"from":           from,
		"to":             to,
		"text":           text,
		"fromSupportOne": fromSupportOne,
		"fromSupportTwo": fromSupportTwo,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSmartUrl, "SendMultiple")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetDeliveries2SmartSmsAsync(username string, password string, Id int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"Id":       Id,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSmartUrl, "GetDeliveries2")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetDeliveriesSmartSmsAsync(username string, password string, Ids []int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"Ids":      Ids,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSmartUrl, "GetDeliveries")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func AddBodyAsync(username string, password string, title, body string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"title":    title,
		"body":     body,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSharedUrl, "AddBody")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func EditBodyAsync(username string, password string, bodyId int, title, body string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"bodyId":   bodyId,
		"title":    title,
		"body":     body,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSharedUrl, "EditBody")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func DeleteBodyAsync(username string, password string, bodyId int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"bodyId":   bodyId,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSharedUrl, "DeleteBody")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetBodiesAsync(username string, password string, keyword string, status int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"keyword":  keyword,
		"status":   status,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTSharedUrl, "GetBodies")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendSpeechTextAsync(username string, password string, text, to string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"to":       to,
		"text":     text,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "SendSpeechText")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendVoiceOtpAsync(username string, password string, to, code string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"code":     code,
		"to":       to,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "SendVoiceOtp")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendSMSWithSpeechTextAsync(username string, password string, smsBody, speechBody, from, to string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":   username,
		"PassWord":   password,
		"from":       from,
		"to":         to,
		"speechBody": speechBody,
		"smsBody":    smsBody,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "SendSMSWithSpeechText")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendSMSWithSpeechTextBySchduleDateAsync(username string, password string, smsBody, speechBody, scheduleDate, from, to string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":     username,
		"PassWord":     password,
		"from":         from,
		"to":           to,
		"scheduleDate": scheduleDate,
		"smsBody":      smsBody,
		"speechBody":   speechBody,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "SendSMSWithSpeechTextBySchduleDate")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetSendSpeechTextStatusAsync(username string, password string, recId string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recId":    recId,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "GetSendSpeechTextStatus")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetSendSpeechTextStatusesAsync(username string, password string, recIds []string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recIds":   recIds,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "GetSendSpeechTextStatuses")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func GetSendSMSWithSpeechTextStatusAsync(username string, password string, recId string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recId":    recId,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "GetSendSMSWithSpeechTextStatus")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendBulkSpeechTextAsync(username string, password string, title, body, receivers, DateToSend string, RepeatCount int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":    username,
		"PassWord":    password,
		"title":       title,
		"body":        body,
		"receivers":   receivers,
		"DateToSend":  DateToSend,
		"RepeatCount": RepeatCount,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "SendBulkSpeechText")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func SendBulkVoiceSMSAsync(username string, password string, title, receivers, DateToSend string, voiceFileId, repeatCount int) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":    username,
		"PassWord":    password,
		"title":       title,
		"receivers":   receivers,
		"DateToSend":  DateToSend,
		"voiceFileId": voiceFileId,
		"repeatCount": repeatCount,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "SendBulkVoiceSMS")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func UploadVoiceFileAsync(username string, password string, title, base64StringFile string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName":         username,
		"PassWord":         password,
		"base64StringFile": base64StringFile,
		"title":            title,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "UploadVoiceFile")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}

func UploadVoiceFileByteAsync(username string, password string, title, file string) <-chan Result[string] {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"title":    title,
		"file":     file,
	}
	result := make(chan Result[string], 1)
	go func() {
		response, err := invoke(jsonData, RESTVoiceUrl, "UploadVoiceFileByte")
		result <- Result[string]{Value: response, Err: err}
		close(result)
	}()
	return result
}
