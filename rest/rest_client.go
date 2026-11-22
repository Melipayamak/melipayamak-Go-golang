package melipayamak

import (
	"strconv"
)

func SendSMS(username string, password string, to string, from string, text string, isFlash bool) (string, error) {

	jsonData := map[string]any{
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

	jsonData := map[string]any{
		"username": username,
		"password": password,
		"text":     text,
		"to":       to,
		"bodyId":   strconv.FormatInt(bodyId, 10),
	}
	return invoke(jsonData, RESTSendUrl, "BaseServiceNumber")
}

func GetDeliveries2(username string, password string, recID int64) (string, error) {
	jsonData := map[string]any{
		"username": username,
		"password": password,
		"recID":    strconv.FormatInt(recID, 10),
	}
	return invoke(jsonData, RESTSendUrl, "GetDeliveries2")
}

func GetMessages(username string, password string, location int, from string, index string, count bool) (string, error) {
	jsonData := map[string]any{
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
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
	}
	return invoke(jsonData, RESTSendUrl, "GetCredit")
}

func GetBasePrice(username string, password string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
	}
	return invoke(jsonData, RESTSendUrl, "GetBasePrice")
}

func GetUserNumbers(username string, password string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
	}
	return invoke(jsonData, RESTSendUrl, "GetUserNumbers")
}

// v2
func SendOTP(username string, password string, to, from, code string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"to":       to,
		"from":     from,
		"code":     code,
	}
	return invoke(jsonData, RESTSendUrl, "SendOTP")
}

func SendMultipleSMS(username string, password string, from string, to []string, text []string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"from":     from,
		"to":       to,
		"text":     text,
	}
	return invoke(jsonData, RESTSendUrl, "SendMultipleSMS")
}

func SendSchedule(username string, password string, text, to, from, scheduleDate string, period int) (string, error) {
	jsonData := map[string]any{
		"UserName":     username,
		"PassWord":     password,
		"from":         from,
		"to":           to,
		"text":         text,
		"scheduledate": scheduleDate,
		"period":       period,
	}
	return invoke(jsonData, RESTSendUrl, "SendSchedule")
}

func GetSchedule(username string, password string, scheduleId int) (string, error) {
	jsonData := map[string]any{
		"UserName":   username,
		"PassWord":   password,
		"scheduleId": scheduleId,
	}
	return invoke(jsonData, RESTSendUrl, "GetSchedule")
}

func GetDeliveries(username string, password string, recID string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recID":    recID,
	}
	return invoke(jsonData, RESTSendUrl, "GetDeliveries")
}

func GetNewDeliveries(username string, password string, recID []string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recID":    recID,
	}
	return invoke(jsonData, RESTSendUrl, "GetNewDeliveries")
}

func GetInboxCount(username string, password string, isRead bool) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"isRead":   isRead,
	}
	return invoke(jsonData, RESTSendUrl, "GetInboxCount")
}

func GetCredit2(username string, password string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
	}
	return invoke(jsonData, RESTSendUrl, "GetCredit2")
}

func GetSmsPrice(username string, password string, irancellCount, mciCount int, from, text string) (string, error) {
	jsonData := map[string]any{
		"UserName":      username,
		"PassWord":      password,
		"irancellCount": irancellCount,
		"mciCount":      mciCount,
		"from":          from,
		"text":          text,
	}
	return invoke(jsonData, RESTSendUrl, "GetSmsPrice")
}

func BaseServiceNumber(username string, password string, text, to string, bodyId int) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"bodyId":   bodyId,
		"to":       to,
		"text":     text,
	}
	return invoke(jsonData, RESTSendUrl, "BaseServiceNumber")
}

func BaseServiceNumberSchedule(username string, password string, text, to string, bodyId int, scheduleDate string) (string, error) {
	jsonData := map[string]any{
		"UserName":     username,
		"PassWord":     password,
		"bodyId":       bodyId,
		"to":           to,
		"text":         text,
		"scheduleDate": scheduleDate,
	}
	return invoke(jsonData, RESTSendUrl, "BaseServiceNumberSchedule")
}

func SendSmartSms(username string, password string, to, text, from, fromSupportOne, fromSupportTwo string) (string, error) {
	jsonData := map[string]any{
		"UserName":       username,
		"PassWord":       password,
		"from":           from,
		"to":             to,
		"text":           text,
		"fromSupportOne": fromSupportOne,
		"fromSupportTwo": fromSupportTwo,
	}
	return invoke(jsonData, RESTSmartUrl, "Send")
}

func SendMultipleSmartSms(username string, password string, to, text []string, from, fromSupportOne, fromSupportTwo string) (string, error) {
	jsonData := map[string]any{
		"UserName":       username,
		"PassWord":       password,
		"from":           from,
		"to":             to,
		"text":           text,
		"fromSupportOne": fromSupportOne,
		"fromSupportTwo": fromSupportTwo,
	}
	return invoke(jsonData, RESTSmartUrl, "SendMultiple")
}

func GetDeliveries2SmartSms(username string, password string, Id int) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"Id":       Id,
	}
	return invoke(jsonData, RESTSmartUrl, "GetDeliveries2")
}

func GetDeliveriesSmartSms(username string, password string, Ids []int) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"Ids":      Ids,
	}
	return invoke(jsonData, RESTSmartUrl, "GetDeliveries")
}

func AddBody(username string, password string, title, body string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"title":    title,
		"body":     body,
	}
	return invoke(jsonData, RESTSharedUrl, "AddBody")
}

func EditBody(username string, password string, bodyId int, title, body string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"bodyId":   bodyId,
		"title":    title,
		"body":     body,
	}
	return invoke(jsonData, RESTSharedUrl, "EditBody")
}

func DeleteBody(username string, password string, bodyId int) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"bodyId":   bodyId,
	}
	return invoke(jsonData, RESTSharedUrl, "DeleteBody")
}

func GetBodies(username string, password string, keyword string, status int) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"keyword":  keyword,
		"status":   status,
	}
	return invoke(jsonData, RESTSharedUrl, "GetBodies")
}

func SendSpeechText(username string, password string, text, to string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"to":       to,
		"text":     text,
	}
	return invoke(jsonData, RESTVoiceUrl, "SendSpeechText")
}

func SendVoiceOtp(username string, password string, to, code string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"code":     code,
		"to":       to,
	}
	return invoke(jsonData, RESTVoiceUrl, "SendVoiceOtp")
}

func SendSMSWithSpeechText(username string, password string, smsBody, speechBody, from, to string) (string, error) {
	jsonData := map[string]any{
		"UserName":   username,
		"PassWord":   password,
		"from":       from,
		"to":         to,
		"speechBody": speechBody,
		"smsBody":    smsBody,
	}
	return invoke(jsonData, RESTVoiceUrl, "SendSMSWithSpeechText")
}

func SendSMSWithSpeechTextBySchduleDate(username string, password string, smsBody, speechBody, scheduleDate, from, to string) (string, error) {
	jsonData := map[string]any{
		"UserName":     username,
		"PassWord":     password,
		"from":         from,
		"to":           to,
		"scheduleDate": scheduleDate,
		"smsBody":      smsBody,
		"speechBody":   speechBody,
	}
	return invoke(jsonData, RESTVoiceUrl, "SendSMSWithSpeechTextBySchduleDate")
}

func GetSendSpeechTextStatus(username string, password string, recId string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recId":    recId,
	}
	return invoke(jsonData, RESTVoiceUrl, "GetSendSpeechTextStatus")
}

func GetSendSpeechTextStatuses(username string, password string, recIds []string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recIds":   recIds,
	}
	return invoke(jsonData, RESTVoiceUrl, "GetSendSpeechTextStatuses")
}

func GetSendSMSWithSpeechTextStatus(username string, password string, recId string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"recId":    recId,
	}
	return invoke(jsonData, RESTVoiceUrl, "GetSendSMSWithSpeechTextStatus")
}

func SendBulkSpeechText(username string, password string, title, body, receivers, DateToSend string, RepeatCount int) (string, error) {
	jsonData := map[string]any{
		"UserName":    username,
		"PassWord":    password,
		"title":       title,
		"body":        body,
		"receivers":   receivers,
		"DateToSend":  DateToSend,
		"RepeatCount": RepeatCount,
	}
	return invoke(jsonData, RESTVoiceUrl, "SendBulkSpeechText")
}

func SendBulkVoiceSMS(username string, password string, title, receivers, DateToSend string, voiceFileId, repeatCount int) (string, error) {
	jsonData := map[string]any{
		"UserName":    username,
		"PassWord":    password,
		"title":       title,
		"receivers":   receivers,
		"DateToSend":  DateToSend,
		"voiceFileId": voiceFileId,
		"repeatCount": repeatCount,
	}
	return invoke(jsonData, RESTVoiceUrl, "SendBulkVoiceSMS")
}

func UploadVoiceFile(username string, password string, title, base64StringFile string) (string, error) {
	jsonData := map[string]any{
		"UserName":         username,
		"PassWord":         password,
		"base64StringFile": base64StringFile,
		"title":            title,
	}
	return invoke(jsonData, RESTVoiceUrl, "UploadVoiceFile")
}

func UploadVoiceFileByte(username string, password string, title, file string) (string, error) {
	jsonData := map[string]any{
		"UserName": username,
		"PassWord": password,
		"title":    title,
		"file":     file,
	}
	return invoke(jsonData, RESTVoiceUrl, "UploadVoiceFileByte")
}
