package melipayamak

import (
	"context"
)

// request methods

func (c *SOAPClient) SendSimpleSMS2(ctx context.Context, to string, from string, text string, flash bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "to", Value: to},
		{Name: "from", Value: from},
		{Name: "text", Value: text},
		{Name: "isflash", Value: flash},
	}

	body, err := buildSOAPEnvelope("SendSimpleSMS2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "SendSimpleSMS2", body)
}

func (c *SOAPClient) SendSimpleSMS(ctx context.Context, to []string, from string, text string, flash bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("to", "string", to),
		{Name: "from", Value: from},
		{Name: "text", Value: text},
		{Name: "isflash", Value: flash},
	}

	body, err := buildSOAPEnvelope("SendSimpleSMS", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "SendSimpleSMS", body)
}

func (c *SOAPClient) SendSms(ctx context.Context, to []string, from string, text string, flash bool, udh string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("to", "string", to),
		{Name: "from", Value: from},
		{Name: "text", Value: text},
		{Name: "isflash", Value: flash},
		{Name: "udh", Value: udh},
	}

	body, err := buildSOAPEnvelope("SendSms", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "SendSms", body)
}

func (c *SOAPClient) SendWithDomain(ctx context.Context, to string, from string, text string, flash bool, domain string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "to", Value: to},
		{Name: "from", Value: from},
		{Name: "text", Value: text},
		{Name: "isflash", Value: flash},
		{Name: "domain", Value: domain},
	}

	body, err := buildSOAPEnvelope("SendWithDomain", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "SendWithDomain", body)
}

func (c *SOAPClient) SendByBaseNumber(ctx context.Context, text []string, to string, bodyId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("text", "string", text),
		{Name: "to", Value: to},
		{Name: "bodyId", Value: bodyId},
	}

	body, err := buildSOAPEnvelope("SendByBaseNumber", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "SendByBaseNumber", body)
}

func (c *SOAPClient) SendByBaseNumber2(ctx context.Context, text string, to string, bodyId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "to", Value: to},
		{Name: "text", Value: text},
		{Name: "bodyId", Value: bodyId},
	}

	body, err := buildSOAPEnvelope("SendByBaseNumber2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "SendByBaseNumber2", body)
}

func (c *SOAPClient) getMessages(ctx context.Context, location int, from string, index int, count int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "location", Value: location},
		{Name: "from", Value: from},
		{Name: "index", Value: index},
		{Name: "count", Value: count},
	}

	body, err := buildSOAPEnvelope("getMessages", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "getMessages", body)
}

func (c *SOAPClient) GetSmsPrice(ctx context.Context, irancellCount int, mtnCount int, from string, text string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "irancellCount", Value: irancellCount},
		{Name: "mtnCount", Value: mtnCount},
		{Name: "from", Value: from},
		{Name: "text", Value: text},
	}

	body, err := buildSOAPEnvelope("GetSmsPrice", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetSmsPrice", body)
}

func (c *SOAPClient) GetMultiDelivery2(ctx context.Context, recId string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "recId", Value: recId},
	}

	body, err := buildSOAPEnvelope("GetMultiDelivery2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetMultiDelivery2", body)
}

func (c *SOAPClient) GetMultiDelivery(ctx context.Context, recId string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "recId", Value: recId},
	}

	body, err := buildSOAPEnvelope("GetMultiDelivery", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetMultiDelivery", body)
}

func (c *SOAPClient) GetInboxCount(ctx context.Context, isRead bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "isRead", Value: isRead},
	}

	body, err := buildSOAPEnvelope("GetInboxCount", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetInboxCount", body)
}

func (c *SOAPClient) GetDelivery2(ctx context.Context, recId string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "recId", Value: recId},
	}

	body, err := buildSOAPEnvelope("GetDelivery2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetDelivery2", body)
}

func (c *SOAPClient) GetDelivery(ctx context.Context, recId string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "recId", Value: recId},
	}

	body, err := buildSOAPEnvelope("GetDelivery", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetDelivery", body)
}

func (c *SOAPClient) GetDeliveries3(ctx context.Context, recIds []string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("recIds", "string", recIds),
	}

	body, err := buildSOAPEnvelope("GetDeliveries3", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetDeliveries3", body)
}

func (c *SOAPClient) GetDeliveries2(ctx context.Context, recId string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "recId", Value: recId},
	}

	body, err := buildSOAPEnvelope("GetDeliveries2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetDeliveries2", body)
}

func (c *SOAPClient) GetDeliveries(ctx context.Context, recIds []int64) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("recIds", "long", recIds),
	}

	body, err := buildSOAPEnvelope("GetDeliveries", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetDeliveries", body)
}

func (c *SOAPClient) GetCredit(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetCredit", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPSendUrl, "GetCredit", body)
}

// Receive

func (c *SOAPClient) RemoveMessages2(ctx context.Context, location int, msgIds string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "location", Value: location},
		{Name: "msgIds", Value: msgIds},
	}

	body, err := buildSOAPEnvelope("RemoveMessages2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "RemoveMessages2", body)
}

func (c *SOAPClient) RemoveMessages(ctx context.Context, location string, msgIds string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "msgIds", Value: msgIds},
		{Name: "location", Value: location},
	}

	body, err := buildSOAPEnvelope("RemoveMessages", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "RemoveMessages", body)
}

func (c *SOAPClient) GetUsersMessagesByDate(ctx context.Context, location int, from string, index int, count int, dateFrom string, dateTo string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "location", Value: location},
		{Name: "index", Value: index},
		{Name: "count", Value: count},
		{Name: "dateFrom", Value: dateFrom},
		{Name: "dateTo", Value: dateTo},
	}

	body, err := buildSOAPEnvelope("GetUsersMessagesByDate", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetUsersMessagesByDate", body)
}

func (c *SOAPClient) GetOutBoxCount(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetOutBoxCount", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetOutBoxCount", body)
}

func (c *SOAPClient) GetMessagesWithChangeIsRead(ctx context.Context, location int, from string, index int, count int, isRead bool, changeIsRead bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "location", Value: location},
		{Name: "index", Value: index},
		{Name: "count", Value: count},
		{Name: "isRead", Value: isRead},
		{Name: "changeIsRead", Value: changeIsRead},
	}

	body, err := buildSOAPEnvelope("GetMessagesWithChangeIsRead", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetMessagesWithChangeIsRead", body)
}

func (c *SOAPClient) GetMessagesReceptions(ctx context.Context, msgId int, fromRows int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "msgId", Value: msgId},
		{Name: "fromRows", Value: fromRows},
	}

	body, err := buildSOAPEnvelope("GetMessagesReceptions", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetMessagesReceptions", body)
}

func (c *SOAPClient) GetMessagesFilterByDate(ctx context.Context, location int, from string, index int, count int, dateFrom string, dateTo string, isRead bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "location", Value: location},
		{Name: "index", Value: index},
		{Name: "count", Value: count},
		{Name: "dateFrom", Value: dateFrom},
		{Name: "dateTo", Value: dateTo},
		{Name: "isRead", Value: isRead},
	}

	body, err := buildSOAPEnvelope("GetMessagesFilterByDate", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetMessagesFilterByDate", body)
}

func (c *SOAPClient) GetMessagesByDate(ctx context.Context, location int, from string, index int, count int, dateFrom string, dateTo string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "location", Value: location},
		{Name: "index", Value: index},
		{Name: "count", Value: count},
		{Name: "dateFrom", Value: dateFrom},
		{Name: "dateTo", Value: dateTo},
	}

	body, err := buildSOAPEnvelope("GetMessagesByDate", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetMessagesByDate", body)
}

func (c *SOAPClient) GetMessagesAfterIDJson(ctx context.Context, location int, from string, count int, msgId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "location", Value: location},
		{Name: "count", Value: count},
		{Name: "msgId", Value: msgId},
	}

	body, err := buildSOAPEnvelope("GetMessagesAfterIDJson", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetMessagesAfterIDJson", body)
}

func (c *SOAPClient) GetMessagesAfterID(ctx context.Context, location int, from string, count int, msgId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "location", Value: location},
		{Name: "count", Value: count},
		{Name: "msgId", Value: msgId},
	}

	body, err := buildSOAPEnvelope("GetMessagesAfterID", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetMessagesAfterID", body)
}

func (c *SOAPClient) GetMessages(ctx context.Context, location int, from string, index int, count int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "location", Value: location},
		{Name: "index", Value: index},
		{Name: "count", Value: count},
	}

	body, err := buildSOAPEnvelope("GetMessages", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetMessages", body)
}

func (c *SOAPClient) GetMessageStr(ctx context.Context, location int, from string, index int, count int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "location", Value: location},
		{Name: "index", Value: index},
		{Name: "count", Value: count},
	}

	body, err := buildSOAPEnvelope("GetMessageStr", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "GetMessageStr", body)
}

func (c *SOAPClient) ChangeMessageIsRead(ctx context.Context, msgIds string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "msgIds", Value: msgIds},
	}

	body, err := buildSOAPEnvelope("ChangeMessageIsRead", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPReceiveUrl, "ChangeMessageIsRead", body)
}

// Users
func (c *SOAPClient) AddPayment(ctx context.Context, name string, family string, bankName string, code string, amount int, cardNumber string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "name", Value: name},
		{Name: "family", Value: family},
		{Name: "bankName", Value: bankName},
		{Name: "code", Value: code},
		{Name: "amount", Value: amount},
		{Name: "cardNumber", Value: cardNumber},
	}

	body, err := buildSOAPEnvelope("AddPayment", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "AddPayment", body)
}

func (c *SOAPClient) AddUser(ctx context.Context, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "productId", Value: productId},
		{Name: "descriptions", Value: descriptions},
		{Name: "mobileNumber", Value: mobileNumber},
		{Name: "emailAddress", Value: emailAddress},
		{Name: "nationalCode", Value: nationalCode},
		{Name: "name", Value: name},
		{Name: "family", Value: family},
		{Name: "corporation", Value: corporation},
		{Name: "phone", Value: phone},
		{Name: "fax", Value: fax},
		{Name: "address", Value: address},
		{Name: "postalCode", Value: postalCode},
		{Name: "certificateNumber", Value: certificateNumber},
	}

	body, err := buildSOAPEnvelope("AddUser", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "AddUser", body)
}

func (c *SOAPClient) AddUserComplete(ctx context.Context, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string, country int, province int, city int, howFindUs string, commercialCode string, saleId string, recommanderId string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "productId", Value: productId},
		{Name: "descriptions", Value: descriptions},
		{Name: "mobileNumber", Value: mobileNumber},
		{Name: "emailAddress", Value: emailAddress},
		{Name: "nationalCode", Value: nationalCode},
		{Name: "name", Value: name},
		{Name: "family", Value: family},
		{Name: "corporation", Value: corporation},
		{Name: "phone", Value: phone},
		{Name: "fax", Value: fax},
		{Name: "address", Value: address},
		{Name: "postalCode", Value: postalCode},
		{Name: "certificateNumber", Value: certificateNumber},
		{Name: "country", Value: country},
		{Name: "province", Value: province},
		{Name: "city", Value: city},
		{Name: "howFindUs", Value: howFindUs},
		{Name: "commercialCode", Value: commercialCode},
		{Name: "saleId", Value: saleId},
		{Name: "recommanderId", Value: recommanderId},
	}

	body, err := buildSOAPEnvelope("AddUserComplete", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "AddUserComplete", body)
}

func (c *SOAPClient) AddUserWithLocation(ctx context.Context, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string, country int, province int, city int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "productId", Value: productId},
		{Name: "descriptions", Value: descriptions},
		{Name: "mobileNumber", Value: mobileNumber},
		{Name: "emailAddress", Value: emailAddress},
		{Name: "nationalCode", Value: nationalCode},
		{Name: "name", Value: name},
		{Name: "family", Value: family},
		{Name: "corporation", Value: corporation},
		{Name: "phone", Value: phone},
		{Name: "fax", Value: fax},
		{Name: "address", Value: address},
		{Name: "postalCode", Value: postalCode},
		{Name: "certificateNumber", Value: certificateNumber},
		{Name: "country", Value: country},
		{Name: "province", Value: province},
		{Name: "city", Value: city},
	}

	body, err := buildSOAPEnvelope("AddUserWithLocation", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "AddUserWithLocation", body)
}

func (c *SOAPClient) AddUserWithMobileNumber(ctx context.Context, productId int, mobileNumber string, firstName string, lastName string, email string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "productId", Value: productId},
		{Name: "mobileNumber", Value: mobileNumber},
		{Name: "firstName", Value: firstName},
		{Name: "lastName", Value: lastName},
		{Name: "email", Value: email},
	}

	body, err := buildSOAPEnvelope("AddUserWithMobileNumber", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "AddUserWithMobileNumber", body)
}

func (c *SOAPClient) AddUserWithUserNameAndPass(ctx context.Context, targetUserName string, targetUserPassword string, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "productId", Value: productId},
		{Name: "descriptions", Value: descriptions},
		{Name: "mobileNumber", Value: mobileNumber},
		{Name: "emailAddress", Value: emailAddress},
		{Name: "nationalCode", Value: nationalCode},
		{Name: "name", Value: name},
		{Name: "family", Value: family},
		{Name: "corporation", Value: corporation},
		{Name: "phone", Value: phone},
		{Name: "fax", Value: fax},
		{Name: "address", Value: address},
		{Name: "postalCode", Value: postalCode},
		{Name: "certificateNumber", Value: certificateNumber},
		{Name: "targetUserName", Value: targetUserName},
		{Name: "targetUserPassword", Value: targetUserPassword},
	}

	body, err := buildSOAPEnvelope("AddUserWithUserNameAndPass", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "AddUserWithUserNameAndPass", body)
}

func (c *SOAPClient) AuthenticateUser(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("AuthenticateUser", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "AuthenticateUser", body)
}

func (c *SOAPClient) ChangeUserCredit(ctx context.Context, amount int, description string, targetUsername string, GetTax bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "amount", Value: amount},
		{Name: "description", Value: description},
		{Name: "targetUsername", Value: targetUsername},
		{Name: "GetTax", Value: GetTax},
	}

	body, err := buildSOAPEnvelope("ChangeUserCredit", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "ChangeUserCredit", body)
}

func (c *SOAPClient) ChangeUserCredit2(ctx context.Context, amount string, description string, targetUsername string, GetTax bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "amount", Value: amount},
		{Name: "targetUsername", Value: targetUsername},
		{Name: "description", Value: description},
		{Name: "GetTax", Value: GetTax},
	}

	body, err := buildSOAPEnvelope("ChangeUserCredit2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "ChangeUserCredit2", body)
}

func (c *SOAPClient) ChangeUserCreditBySeccondPass(ctx context.Context, username string, seccondPassword string, amount int, description string, targetUsername string, GetTax bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "username", Value: username},
		{Name: "seccondPassword", Value: seccondPassword},
		{Name: "amount", Value: amount},
		{Name: "targetUsername", Value: targetUsername},
		{Name: "description", Value: description},
		{Name: "GetTax", Value: GetTax},
	}

	body, err := buildSOAPEnvelope("ChangeUserCreditBySeccondPass", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "ChangeUserCreditBySeccondPass", body)
}

func (c *SOAPClient) DeductUserCredit(ctx context.Context, username string, password string, amount int, description string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "username", Value: username},
		{Name: "password", Value: password},
		{Name: "amount", Value: amount},
		{Name: "description", Value: description},
	}

	body, err := buildSOAPEnvelope("DeductUserCredit", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "DeductUserCredit", body)
}

func (c *SOAPClient) ForgotPassword(ctx context.Context, mobileNumber string, emailAddress string, targetUsername string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "mobileNumber", Value: mobileNumber},
		{Name: "emailAddress", Value: emailAddress},
		{Name: "targetUsername", Value: targetUsername},
	}

	body, err := buildSOAPEnvelope("ForgotPassword", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "ForgotPassword", body)
}

func (c *SOAPClient) GetCities(ctx context.Context, provinceId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "provinceId", Value: provinceId},
	}

	body, err := buildSOAPEnvelope("GetCities", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetCities", body)
}

func (c *SOAPClient) GetEnExpireDate(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetEnExpireDate", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetEnExpireDate", body)
}

func (c *SOAPClient) GetExpireDate(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetExpireDate", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetExpireDate", body)
}

func (c *SOAPClient) GetProvinces(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetProvinces", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetProvinces", body)
}

func (c *SOAPClient) GetUserBasePrice(ctx context.Context, targetUsername string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "targetUsername", Value: targetUsername},
	}

	body, err := buildSOAPEnvelope("GetUserBasePrice", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserBasePrice", body)
}

func (c *SOAPClient) GetUserByUserID(ctx context.Context, pass string, userId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "pass", Value: pass},
		{Name: "userId", Value: userId},
	}

	body, err := buildSOAPEnvelope("GetUserByUserID", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserByUserID", body)
}

func (c *SOAPClient) GetUserCredit(ctx context.Context, targetUsername string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "targetUsername", Value: targetUsername},
	}

	body, err := buildSOAPEnvelope("GetUserCredit", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserCredit", body)
}

func (c *SOAPClient) GetUserCreditBySecondPass(ctx context.Context, secondPassword string, targetUsername string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "secondPassword", Value: secondPassword},
		{Name: "targetUsername", Value: targetUsername},
	}

	body, err := buildSOAPEnvelope("GetUserCreditBySecondPass", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserCreditBySecondPass", body)
}

func (c *SOAPClient) GetUserDetails(ctx context.Context, targetUsername string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "targetUsername", Value: targetUsername},
	}

	body, err := buildSOAPEnvelope("GetUserDetails", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserDetails", body)
}

func (c *SOAPClient) GetUserIsExist(ctx context.Context, targetUsername string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "targetUsername", Value: targetUsername},
	}

	body, err := buildSOAPEnvelope("GetUserIsExist", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserIsExist", body)
}

func (c *SOAPClient) GetUserNumbers(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetUserNumbers", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserNumbers", body)
}

func (c *SOAPClient) GetUserTransactions(ctx context.Context, targetUsername string, creditType string, dateFrom string, dateTo string, keyword string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "targetUsername", Value: targetUsername},
		{Name: "creditType", Value: creditType},
		{Name: "dateFrom", Value: dateFrom},
		{Name: "dateTo", Value: dateTo},
		{Name: "keyword", Value: keyword},
	}

	body, err := buildSOAPEnvelope("GetUserTransactions", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserTransactions", body)
}

func (c *SOAPClient) GetUserWallet(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetUserWallet", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserWallet", body)
}

func (c *SOAPClient) GetUserWalletTransaction(ctx context.Context, dateFrom string, dateTo string, count int, startIndex int, payType string, PayLoc string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "dateFrom", Value: dateFrom},
		{Name: "dateTo", Value: dateTo},
		{Name: "count", Value: count},
		{Name: "startIndex", Value: startIndex},
		{Name: "payType", Value: payType},
		{Name: "PayLoc", Value: PayLoc},
	}

	body, err := buildSOAPEnvelope("GetUserWalletTransaction", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUserWalletTransaction", body)
}

func (c *SOAPClient) GetUsers(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetUsers", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "GetUsers", body)
}

func (c *SOAPClient) HasFilter(ctx context.Context, text string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "text", Value: text},
	}

	body, err := buildSOAPEnvelope("HasFilter", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "HasFilter", body)
}

func (c *SOAPClient) RemoveUser(ctx context.Context, targetUsername string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "targetUsername", Value: targetUsername},
	}

	body, err := buildSOAPEnvelope("RemoveUser", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "RemoveUser", body)
}

func (c *SOAPClient) RevivalUser(ctx context.Context, targetUsername string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "targetUsername", Value: targetUsername},
	}

	body, err := buildSOAPEnvelope("RevivalUser", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPUsersUrl, "RevivalUser", body)
}

// Contacts

func (c *SOAPClient) AddContact(ctx context.Context, groupIds string, firstname string, lastname string, nickname string, corporation string, mobilenumber string, phone string, fax string, birthdate string, email string, gender int, province int, city int, address string, postalCode string, additionaldate string, additionaltext string, descriptions string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "groupIds", Value: groupIds},
		{Name: "firstname", Value: firstname},
		{Name: "lastname", Value: lastname},
		{Name: "nickname", Value: nickname},
		{Name: "corporation", Value: corporation},
		{Name: "mobilenumber", Value: mobilenumber},
		{Name: "phone", Value: phone},
		{Name: "fax", Value: fax},
		{Name: "birthdate", Value: birthdate},
		{Name: "email", Value: email},
		{Name: "gender", Value: gender},
		{Name: "province", Value: province},
		{Name: "city", Value: city},
		{Name: "address", Value: address},
		{Name: "postalCode", Value: postalCode},
		{Name: "additionaldate", Value: additionaldate},
		{Name: "additionaltext", Value: additionaltext},
		{Name: "descriptions", Value: descriptions},
	}

	body, err := buildSOAPEnvelope("AddContact", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "AddContact", body)
}

func (c *SOAPClient) AddContactEvents(ctx context.Context, contactId int, eventName string, eventType int, eventDate string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "contactId", Value: contactId},
		{Name: "eventName", Value: eventName},
		{Name: "eventType", Value: eventType},
		{Name: "eventDate", Value: eventDate},
	}

	body, err := buildSOAPEnvelope("AddContactEvents", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "AddContactEvents", body)
}

func (c *SOAPClient) AddGroup(ctx context.Context, groupName string, Descriptions string, showToChilds bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "groupName", Value: groupName},
		{Name: "Descriptions", Value: Descriptions},
		{Name: "showToChilds", Value: showToChilds},
	}

	body, err := buildSOAPEnvelope("AddGroup", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "AddGroup", body)
}

func (c *SOAPClient) ChangeContact(ctx context.Context, contactId int, mobilenumber string, firstname string, lastname string, nickname string, corporation string, phone string, fax string, email string, gender int, province int, city int, address string, postalCode string, additionaltext string, descriptions string, contactStatus int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "contactId", Value: contactId},
		{Name: "mobilenumber", Value: mobilenumber},
		{Name: "firstname", Value: firstname},
		{Name: "lastname", Value: lastname},
		{Name: "nickname", Value: nickname},
		{Name: "corporation", Value: corporation},
		{Name: "phone", Value: phone},
		{Name: "fax", Value: fax},
		{Name: "email", Value: email},
		{Name: "gender", Value: gender},
		{Name: "province", Value: province},
		{Name: "city", Value: city},
		{Name: "address", Value: address},
		{Name: "postalCode", Value: postalCode},
		{Name: "additionaltext", Value: additionaltext},
		{Name: "descriptions", Value: descriptions},
		{Name: "contactStatus", Value: contactStatus},
	}

	body, err := buildSOAPEnvelope("ChangeContact", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "ChangeContact", body)
}

func (c *SOAPClient) ChangeGroup(ctx context.Context, groupId int, groupName string, Descriptions string, showToChilds bool, groupStatus int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "groupId", Value: groupId},
		{Name: "groupName", Value: groupName},
		{Name: "Descriptions", Value: Descriptions},
		{Name: "showToChilds", Value: showToChilds},
		{Name: "groupStatus", Value: groupStatus},
	}

	body, err := buildSOAPEnvelope("ChangeGroup", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "ChangeGroup", body)
}

func (c *SOAPClient) CheckMobileExistInContact(ctx context.Context, mobileNumber string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "mobileNumber", Value: mobileNumber},
	}

	body, err := buildSOAPEnvelope("CheckMobileExistInContact", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "CheckMobileExistInContact", body)
}

func (c *SOAPClient) GetContactEvents(ctx context.Context, contactId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "contactId", Value: contactId},
	}

	body, err := buildSOAPEnvelope("GetContactEvents", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "GetContactEvents", body)
}

func (c *SOAPClient) GetContacts(ctx context.Context, groupId int, keyword string, from int, count int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "groupId", Value: groupId},
		{Name: "from", Value: from},
		{Name: "count", Value: count},
		{Name: "keyword", Value: keyword},
	}

	body, err := buildSOAPEnvelope("GetContacts", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "GetContacts", body)
}

func (c *SOAPClient) GetContactsByID(ctx context.Context, contactId int, status int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "contactId", Value: contactId},
		{Name: "status", Value: status},
	}

	body, err := buildSOAPEnvelope("GetContactsByID", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "GetContactsByID", body)
}

func (c *SOAPClient) GetGroups(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetGroups", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "GetGroups", body)
}

func (c *SOAPClient) MergeGroups(ctx context.Context, originGroupId int, destinationGroupId int, deleteOriginGroup bool) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "originGroupId", Value: originGroupId},
		{Name: "destinationGroupId", Value: destinationGroupId},
		{Name: "deleteOriginGroup", Value: deleteOriginGroup},
	}

	body, err := buildSOAPEnvelope("MergeGroups", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "MergeGroups", body)
}

func (c *SOAPClient) RemoveContact(ctx context.Context, mobilenumber string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "mobilenumber", Value: mobilenumber},
	}

	body, err := buildSOAPEnvelope("RemoveContact", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "RemoveContact", body)
}

func (c *SOAPClient) RemoveContactByContactID(ctx context.Context, contactId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "contactId", Value: contactId},
	}

	body, err := buildSOAPEnvelope("RemoveContactByContactID", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "RemoveContactByContactID", body)
}

func (c *SOAPClient) RemoveGroup(ctx context.Context, groupId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "groupId", Value: groupId},
	}

	body, err := buildSOAPEnvelope("RemoveGroup", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPContactsUrl, "RemoveGroup", body)
}

// ACtions

func (c *SOAPClient) AddBranch(ctx context.Context, branchName string, owner int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "branchName", Value: branchName},
		{Name: "owner", Value: owner},
	}

	body, err := buildSOAPEnvelope("AddBranch", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "AddBranch", body)
}

func (c *SOAPClient) AddBulk(ctx context.Context, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "branch", Value: branch},
		{Name: "bulkType", Value: bulkType},
		{Name: "title", Value: title},
		{Name: "message", Value: message},
		{Name: "rangeFrom", Value: rangeFrom},
		{Name: "rangeTo", Value: rangeTo},
		{Name: "DateToSend", Value: DateToSend},
		{Name: "requestCount", Value: requestCount},
		{Name: "rowFrom", Value: rowFrom},
	}

	body, err := buildSOAPEnvelope("AddBulk", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "AddBulk", body)
}

func (c *SOAPClient) AddBulk2(ctx context.Context, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "branch", Value: branch},
		{Name: "bulkType", Value: bulkType},
		{Name: "title", Value: title},
		{Name: "message", Value: message},
		{Name: "rangeFrom", Value: rangeFrom},
		{Name: "rangeTo", Value: rangeTo},
		{Name: "DateToSend", Value: DateToSend},
		{Name: "requestCount", Value: requestCount},
		{Name: "rowFrom", Value: rowFrom},
	}

	body, err := buildSOAPEnvelope("AddBulk2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "AddBulk2", body)
}

func (c *SOAPClient) AddNewBulk(ctx context.Context, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "from", Value: from},
		{Name: "branch", Value: branch},
		{Name: "bulkType", Value: bulkType},
		{Name: "title", Value: title},
		{Name: "message", Value: message},
		{Name: "rangeFrom", Value: rangeFrom},
		{Name: "rangeTo", Value: rangeTo},
		{Name: "DateToSend", Value: DateToSend},
		{Name: "requestCount", Value: requestCount},
		{Name: "rowFrom", Value: rowFrom},
	}

	body, err := buildSOAPEnvelope("AddNewBulk", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "AddNewBulk", body)
}

func (c *SOAPClient) AddNumber(ctx context.Context, branchId int, mobileNumbers []string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("mobileNumbers", "string", mobileNumbers),
		{Name: "branchId", Value: branchId},
	}

	body, err := buildSOAPEnvelope("AddNumber", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "AddNumber", body)
}

func (c *SOAPClient) GetBranchs(ctx context.Context, owner int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "owner", Value: owner},
	}

	body, err := buildSOAPEnvelope("GetBranchs", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetBranchs", body)
}

func (c *SOAPClient) GetBulk(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetBulk", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetBulk", body)
}

func (c *SOAPClient) GetBulkCount(ctx context.Context, branch int, rangeFrom string, rangeTo string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "branch", Value: branch},
		{Name: "rangeFrom", Value: rangeFrom},
		{Name: "rangeTo", Value: rangeTo},
	}

	body, err := buildSOAPEnvelope("GetBulkCount", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetBulkCount", body)
}

func (c *SOAPClient) GetBulkReceptions(ctx context.Context, bulkId int, fromRows int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "bulkId", Value: bulkId},
		{Name: "fromRows", Value: fromRows},
	}

	body, err := buildSOAPEnvelope("GetBulkReceptions", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetBulkReceptions", body)
}

func (c *SOAPClient) GetBulkStatus(ctx context.Context, bulkId int, sent int, failed int, status int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "bulkId", Value: bulkId},
		{Name: "sent", Value: sent},
		{Name: "status", Value: status},
		{Name: "failed", Value: failed},
	}

	body, err := buildSOAPEnvelope("GetBulkStatus", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetBulkStatus", body)
}

func (c *SOAPClient) GetMobileCount(ctx context.Context, branch int, rangeFrom string, rangeTo string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "branch", Value: branch},
		{Name: "rangeFrom", Value: rangeFrom},
		{Name: "rangeTo", Value: rangeTo},
	}

	body, err := buildSOAPEnvelope("GetMobileCount", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetMobileCount", body)
}

func (c *SOAPClient) GetSendBulk(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetSendBulk", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetSendBulk", body)
}

func (c *SOAPClient) GetTodaySent(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetTodaySent", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetTodaySent", body)
}

func (c *SOAPClient) GetTotalSent(ctx context.Context) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
	}

	body, err := buildSOAPEnvelope("GetTotalSent", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "GetTotalSent", body)
}

func (c *SOAPClient) RemoveBranch(ctx context.Context, branchId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "branchId", Value: branchId},
	}

	body, err := buildSOAPEnvelope("RemoveBranch", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "RemoveBranch", body)
}

func (c *SOAPClient) RemoveBulk(ctx context.Context, bulkId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "bulkId", Value: bulkId},
	}

	body, err := buildSOAPEnvelope("RemoveBulk", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "RemoveBulk", body)
}

func (c *SOAPClient) SendMultipleSMS(ctx context.Context, to []string, from string, text []string, isflash bool, udh string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("to", "string", to),
		SOAPArray("text", "string", text),
		{Name: "from", Value: from},
		{Name: "isflash", Value: isflash},
		{Name: "udh", Value: udh},
	}

	body, err := buildSOAPEnvelope("SendMultipleSMS", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "SendMultipleSMS", body)
}

func (c *SOAPClient) SendMultipleSMS2(ctx context.Context, to []string, from []string, text []string, isflash bool, udh string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("to", "string", to),
		SOAPArray("from", "string", from),
		SOAPArray("text", "string", text),
		{Name: "isflash", Value: isflash},
		{Name: "udh", Value: udh},
	}

	body, err := buildSOAPEnvelope("SendMultipleSMS2", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "SendMultipleSMS2", body)
}

func (c *SOAPClient) UpdateBulkDelivery(ctx context.Context, bulkId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "bulkId", Value: bulkId},
	}

	body, err := buildSOAPEnvelope("UpdateBulkDelivery", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPActionsUrl, "UpdateBulkDelivery", body)
}

// Schedule

func (c *SOAPClient) AddMultipleSchedule(ctx context.Context, to []string, from string, text []string, isflash bool, scheduleDateTime []string, period string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("to", "string", to),
		{Name: "from", Value: from},
		SOAPArray("text", "string", text),
		{Name: "isflash", Value: isflash},
		SOAPArray("scheduleDateTime", "string", scheduleDateTime),
		{Name: "period", Value: period},
	}

	body, err := buildSOAPEnvelope("AddMultipleSchedule", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPScheduleUrl, "AddMultipleSchedule", body)
}

func (c *SOAPClient) AddNewUsance(ctx context.Context, to string, from string, text string, isflash bool, scheduleStartDateTime string, countrepeat int, scheduleEndDateTime string, periodType string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "to", Value: to},
		{Name: "from", Value: from},
		{Name: "text", Value: text},
		{Name: "isflash", Value: isflash},
		{Name: "scheduleStartDateTime", Value: scheduleStartDateTime},
		{Name: "countrepeat", Value: countrepeat},
		{Name: "scheduleEndDateTime", Value: scheduleEndDateTime},
		{Name: "periodType", Value: periodType},
	}

	body, err := buildSOAPEnvelope("AddNewUsance", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPScheduleUrl, "AddNewUsance", body)
}

func (c *SOAPClient) AddSchedule(ctx context.Context, to string, from string, text string, isflash bool, scheduleDateTime string, period string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "to", Value: to},
		{Name: "from", Value: from},
		{Name: "text", Value: text},
		{Name: "isflash", Value: isflash},
		{Name: "scheduleDateTime", Value: scheduleDateTime},
		{Name: "period", Value: period},
	}

	body, err := buildSOAPEnvelope("AddSchedule", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPScheduleUrl, "AddSchedule", body)
}

func (c *SOAPClient) AddUsance(ctx context.Context, to string, from string, text string, isflash bool, scheduleStartDateTime string, repeatAfterDays int, scheduleEndDateTime string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "to", Value: to},
		{Name: "from", Value: from},
		{Name: "text", Value: text},
		{Name: "isflash", Value: isflash},
		{Name: "scheduleStartDateTime", Value: scheduleStartDateTime},
		{Name: "repeatAfterDays", Value: repeatAfterDays},
		{Name: "scheduleEndDateTime", Value: scheduleEndDateTime},
	}

	body, err := buildSOAPEnvelope("AddUsance", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPScheduleUrl, "AddUsance", body)
}

func (c *SOAPClient) GetScheduleStatus(ctx context.Context, scheduleId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "scheduleId", Value: scheduleId},
	}

	body, err := buildSOAPEnvelope("GetScheduleStatus", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPScheduleUrl, "GetScheduleStatus", body)
}

func (c *SOAPClient) RemoveSchedule(ctx context.Context, scheduleId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "scheduleId", Value: scheduleId},
	}

	body, err := buildSOAPEnvelope("RemoveSchedule", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPScheduleUrl, "RemoveSchedule", body)
}

func (c *SOAPClient) RemoveScheduleList(ctx context.Context, scheduleIdList []int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		SOAPArray("scheduleIdList", "int", scheduleIdList),
	}

	body, err := buildSOAPEnvelope("RemoveScheduleList", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPScheduleUrl, "RemoveScheduleList", body)
}

// Voice
func (c *SOAPClient) GetSendSMSWithSpeechTextStatus(ctx context.Context, recId int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "recId", Value: recId},
	}

	body, err := buildSOAPEnvelope("GetSendSMSWithSpeechTextStatus", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPVoiceUrl, "GetSendSMSWithSpeechTextStatus", body)
}

func (c *SOAPClient) SendBulkSpeechText(ctx context.Context, title string, bbody string, receivers string, dateToSend string, repeatCount int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "title", Value: title},
		{Name: "body", Value: bbody},
		{Name: "receivers", Value: receivers},
		{Name: "DateToSend", Value: dateToSend},
		{Name: "repeatCount", Value: repeatCount},
	}

	body, err := buildSOAPEnvelope("SendBulkSpeechText", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPVoiceUrl, "SendBulkSpeechText", body)
}

func (c *SOAPClient) SendBulkVoiceSMS(ctx context.Context, title string, voiceFileId int, receivers string, DateToSend string, repeatCount int) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "title", Value: title},
		{Name: "voiceFileId", Value: voiceFileId},
		{Name: "receivers", Value: receivers},
		{Name: "DateToSend", Value: DateToSend},
		{Name: "repeatCount", Value: repeatCount},
	}

	body, err := buildSOAPEnvelope("SendBulkVoiceSMS", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPVoiceUrl, "SendBulkVoiceSMS", body)
}

func (c *SOAPClient) SendSMSWithSpeechText(ctx context.Context, smsBody string, speechBody string, from string, to string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "smsBody", Value: smsBody},
		{Name: "speechBody", Value: speechBody},
		{Name: "from", Value: from},
		{Name: "to", Value: to},
	}

	body, err := buildSOAPEnvelope("SendSMSWithSpeechText", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPVoiceUrl, "SendSMSWithSpeechText", body)
}

func (c *SOAPClient) SendSMSWithSpeechTextBySchduleDate(ctx context.Context, smsBody string, speechBody string, from string, to string, scheduleDate string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "smsBody", Value: smsBody},
		{Name: "from", Value: from},
		{Name: "speechBody", Value: speechBody},
		{Name: "to", Value: to},
		{Name: "scheduleDate", Value: scheduleDate},
	}

	body, err := buildSOAPEnvelope("SendSMSWithSpeechTextBySchduleDate", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPVoiceUrl, "SendSMSWithSpeechTextBySchduleDate", body)
}

func (c *SOAPClient) UploadVoiceFile(ctx context.Context, title string, base64StringFile string) ([]byte, error) {

	params := []SOAPParam{
		{Name: "username", Value: c.username},
		{Name: "password", Value: c.password},
		{Name: "title", Value: title},
		{Name: "base64StringFile", Value: base64StringFile},
	}

	body, err := buildSOAPEnvelope("UploadVoiceFile", params...)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, SOAPVoiceUrl, "UploadVoiceFile", body)
}
