package melipayamak

import "context"

func (c *SOAPClient) SendSimpleSMS2Async(ctx context.Context, to, from, text string, flash bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendSimpleSMS2(ctx, to, from, text, flash) })
}

func (c *SOAPClient) SendSimpleSMSAsync(ctx context.Context, to []string, from string, text string, flash bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendSimpleSMS(ctx, to, from, text, flash) })
}

func (c *SOAPClient) SendSmsAsync(ctx context.Context, to []string, from string, text string, flash bool, udh string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendSms(ctx, to, from, text, flash, udh) })
}

func (c *SOAPClient) SendWithDomainAsync(ctx context.Context, to string, from string, text string, flash bool, domain string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendWithDomain(ctx, to, from, text, flash, domain) })
}

func (c *SOAPClient) SendByBaseNumberAsync(ctx context.Context, text []string, to string, bodyId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendByBaseNumber(ctx, text, to, bodyId) })
}

func (c *SOAPClient) SendByBaseNumber2Async(ctx context.Context, text string, to string, bodyId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendByBaseNumber2(ctx, text, to, bodyId) })
}

func (c *SOAPClient) getMessagesAsync(ctx context.Context, location int, from string, index int, count int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.getMessages(ctx, location, from, index, count) })
}

func (c *SOAPClient) GetSmsPriceAsync(ctx context.Context, irancellCount int, mtnCount int, from string, text string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetSmsPrice(ctx, irancellCount, mtnCount, from, text) })
}

func (c *SOAPClient) GetMultiDelivery2Async(ctx context.Context, recId string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetMultiDelivery2(ctx, recId) })
}

func (c *SOAPClient) GetMultiDeliveryAsync(ctx context.Context, recId string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetMultiDelivery(ctx, recId) })
}

func (c *SOAPClient) GetInboxCountAsync(ctx context.Context, isRead bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetInboxCount(ctx, isRead) })

}

func (c *SOAPClient) GetDelivery2Async(ctx context.Context, recId string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetDelivery2(ctx, recId) })
}

func (c *SOAPClient) GetDeliveryAsync(ctx context.Context, recId string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetDelivery(ctx, recId) })
}

func (c *SOAPClient) GetDeliveries3Async(ctx context.Context, recIds []string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetDeliveries3(ctx, recIds) })
}

func (c *SOAPClient) GetDeliveries2Async(ctx context.Context, recId string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetDeliveries2(ctx, recId) })
}

func (c *SOAPClient) GetDeliveriesAsync(ctx context.Context, recIds []int64) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetDeliveries(ctx, recIds) })
}

func (c *SOAPClient) GetCreditAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetCredit(ctx) })
}

// Receive

func (c *SOAPClient) RemoveMessages2Async(ctx context.Context, location int, msgIds string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveMessages2(ctx, location, msgIds) })
}

func (c *SOAPClient) RemoveMessagesAsync(ctx context.Context, location string, msgIds string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveMessages(ctx, location, msgIds) })
}

func (c *SOAPClient) GetUsersMessagesByDateAsync(ctx context.Context, location int, from string, index int, count int, dateFrom string, dateTo string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.GetUsersMessagesByDate(ctx, location, from, index, count, dateFrom, dateTo)
	})
}

func (c *SOAPClient) GetOutBoxCountAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetOutBoxCount(ctx) })
}

func (c *SOAPClient) GetMessagesWithChangeIsReadAsync(ctx context.Context, location int, from string, index int, count int, isRead bool, changeIsRead bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.GetMessagesWithChangeIsRead(ctx, location, from, index, count, isRead, changeIsRead)
	})
}

func (c *SOAPClient) GetMessagesReceptionsAsync(ctx context.Context, msgId int, fromRows int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetMessagesReceptions(ctx, msgId, fromRows) })
}

func (c *SOAPClient) GetMessagesFilterByDateAsync(ctx context.Context, location int, from string, index int, count int, dateFrom string, dateTo string, isRead bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.GetMessagesFilterByDate(ctx, location, from, index, count, dateFrom, dateTo, isRead)
	})
}

func (c *SOAPClient) GetMessagesByDateAsync(ctx context.Context, location int, from string, index int, count int, dateFrom string, dateTo string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.GetMessagesByDate(ctx, location, from, index, count, dateFrom, dateTo)
	})
}

func (c *SOAPClient) GetMessagesAfterIDJsonAsync(ctx context.Context, location int, from string, count int, msgId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetMessagesAfterIDJson(ctx, location, from, count, msgId) })
}

func (c *SOAPClient) GetMessagesAfterIDAsync(ctx context.Context, location int, from string, count int, msgId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetMessagesAfterID(ctx, location, from, count, msgId) })
}

func (c *SOAPClient) GetMessagesAsync(ctx context.Context, location int, from string, index int, count int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetMessages(ctx, location, from, index, count) })
}

func (c *SOAPClient) GetMessageStrAsync(ctx context.Context, location int, from string, index int, count int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetMessageStr(ctx, location, from, index, count) })
}

func (c *SOAPClient) ChangeMessageIsReadAsync(ctx context.Context, msgIds string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.ChangeMessageIsRead(ctx, msgIds) })
}

// Users
func (c *SOAPClient) AddPaymentAsync(ctx context.Context, name string, family string, bankName string, code string, amount int, cardNumber string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.AddPayment(ctx, name, family, bankName, code, amount, cardNumber) })
}

func (c *SOAPClient) AddUserAsync(ctx context.Context, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddUser(ctx, productId, descriptions, mobileNumber, emailAddress, nationalCode, name, family, corporation, phone, fax, address, postalCode, certificateNumber)
	})
}

func (c *SOAPClient) AddUserCompleteAsync(ctx context.Context, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string, country int, province int, city int, howFindUs string, commercialCode string, saleId string, recommanderId string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddUserComplete(ctx, productId, descriptions, mobileNumber, emailAddress, nationalCode, name, family, corporation, phone, fax, address, postalCode, certificateNumber, country, province, city, howFindUs, commercialCode, saleId, recommanderId)
	})
}

func (c *SOAPClient) AddUserWithLocationAsync(ctx context.Context, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string, country int, province int, city int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddUserWithLocation(ctx, productId, descriptions, mobileNumber, emailAddress, nationalCode, name, family, corporation, phone, fax, address, postalCode, certificateNumber, country, province, city)
	})
}

func (c *SOAPClient) AddUserWithMobileNumberAsync(ctx context.Context, productId int, mobileNumber string, firstName string, lastName string, email string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddUserWithMobileNumber(ctx, productId, mobileNumber, firstName, lastName, email)
	})
}

func (c *SOAPClient) AddUserWithUserNameAndPassAsync(ctx context.Context, targetUserName string, targetUserPassword string, productId int, descriptions string, mobileNumber string, emailAddress string, nationalCode string, name string, family string, corporation string, phone string, fax string, address string, postalCode string, certificateNumber string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddUserWithUserNameAndPass(ctx, targetUserName, targetUserPassword, productId, descriptions, mobileNumber, emailAddress, nationalCode, name, family, corporation, phone, fax, address, postalCode, certificateNumber)
	})
}

func (c *SOAPClient) AuthenticateUserAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.AuthenticateUser(ctx) })
}

func (c *SOAPClient) ChangeUserCreditAsync(ctx context.Context, amount int, description string, targetUsername string, GetTax bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.ChangeUserCredit(ctx, amount, description, targetUsername, GetTax) })
}

func (c *SOAPClient) ChangeUserCredit2Async(ctx context.Context, amount string, description string, targetUsername string, GetTax bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.ChangeUserCredit2(ctx, amount, description, targetUsername, GetTax) })
}

func (c *SOAPClient) ChangeUserCreditBySeccondPassAsync(ctx context.Context, username string, seccondPassword string, amount int, description string, targetUsername string, GetTax bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.ChangeUserCreditBySeccondPass(ctx, username, seccondPassword, amount, description, targetUsername, GetTax)
	})
}

func (c *SOAPClient) DeductUserCreditAsync(ctx context.Context, username string, password string, amount int, description string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.DeductUserCredit(ctx, username, password, amount, description) })
}

func (c *SOAPClient) ForgotPasswordAsync(ctx context.Context, mobileNumber string, emailAddress string, targetUsername string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.ForgotPassword(ctx, mobileNumber, emailAddress, targetUsername) })
}

func (c *SOAPClient) GetCitiesAsync(ctx context.Context, provinceId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetCities(ctx, provinceId) })
}

func (c *SOAPClient) GetEnExpireDateAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetEnExpireDate(ctx) })

}

func (c *SOAPClient) GetExpireDateAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetExpireDate(ctx) })
}

func (c *SOAPClient) GetProvincesAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetProvinces(ctx) })
}

func (c *SOAPClient) GetUserBasePriceAsync(ctx context.Context, targetUsername string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUserBasePrice(ctx, targetUsername) })
}

func (c *SOAPClient) GetUserByUserIDAsync(ctx context.Context, pass string, userId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUserByUserID(ctx, pass, userId) })
}

func (c *SOAPClient) GetUserCreditAsync(ctx context.Context, targetUsername string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUserCredit(ctx, targetUsername) })
}

func (c *SOAPClient) GetUserCreditBySecondPassAsync(ctx context.Context, secondPassword string, targetUsername string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUserCreditBySecondPass(ctx, secondPassword, targetUsername) })
}

func (c *SOAPClient) GetUserDetailsAsync(ctx context.Context, targetUsername string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUserDetails(ctx, targetUsername) })
}

func (c *SOAPClient) GetUserIsExistAsync(ctx context.Context, targetUsername string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUserIsExist(ctx, targetUsername) })
}

func (c *SOAPClient) GetUserNumbersAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUserNumbers(ctx) })
}

func (c *SOAPClient) GetUserTransactionsAsync(ctx context.Context, targetUsername string, creditType string, dateFrom string, dateTo string, keyword string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.GetUserTransactions(ctx, targetUsername, creditType, dateFrom, dateTo, keyword)
	})
}

func (c *SOAPClient) GetUserWalletAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUserWallet(ctx) })
}

func (c *SOAPClient) GetUserWalletTransactionAsync(ctx context.Context, dateFrom string, dateTo string, count int, startIndex int, payType string, PayLoc string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.GetUserWalletTransaction(ctx, dateFrom, dateTo, count, startIndex, payType, PayLoc)
	})
}

func (c *SOAPClient) GetUsersAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetUsers(ctx) })
}

func (c *SOAPClient) HasFilterAsync(ctx context.Context, text string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.HasFilter(ctx, text) })
}

func (c *SOAPClient) RemoveUserAsync(ctx context.Context, targetUsername string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveUser(ctx, targetUsername) })
}

func (c *SOAPClient) RevivalUserAsync(ctx context.Context, targetUsername string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RevivalUser(ctx, targetUsername) })
}

// Contacts

func (c *SOAPClient) AddContactAsync(ctx context.Context, groupIds string, firstname string, lastname string, nickname string, corporation string, mobilenumber string, phone string, fax string, birthdate string, email string, gender int, province int, city int, address string, postalCode string, additionaldate string, additionaltext string, descriptions string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddContact(ctx, groupIds, firstname, lastname, nickname, corporation, mobilenumber, phone, fax, birthdate, email, gender, province, city, address, postalCode, additionaldate, additionaltext, descriptions)
	})
}

func (c *SOAPClient) AddContactEventsAsync(ctx context.Context, contactId int, eventName string, eventType int, eventDate string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.AddContactEvents(ctx, contactId, eventName, eventType, eventDate) })
}

func (c *SOAPClient) AddGroupAsync(ctx context.Context, groupName string, Descriptions string, showToChilds bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.AddGroup(ctx, groupName, Descriptions, showToChilds) })
}

func (c *SOAPClient) ChangeContactAsync(ctx context.Context, contactId int, mobilenumber string, firstname string, lastname string, nickname string, corporation string, phone string, fax string, email string, gender int, province int, city int, address string, postalCode string, additionaltext string, descriptions string, contactStatus int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.ChangeContact(ctx, contactId, mobilenumber, firstname, lastname, nickname, corporation, phone, fax, email, gender, province, city, address, postalCode, additionaltext, descriptions, contactStatus)
	})
}

func (c *SOAPClient) ChangeGroupAsync(ctx context.Context, groupId int, groupName string, Descriptions string, showToChilds bool, groupStatus int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.ChangeGroup(ctx, groupId, groupName, Descriptions, showToChilds, groupStatus)
	})
}

func (c *SOAPClient) CheckMobileExistInContactAsync(ctx context.Context, mobileNumber string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.CheckMobileExistInContact(ctx, mobileNumber) })
}

func (c *SOAPClient) GetContactEventsAsync(ctx context.Context, contactId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetContactEvents(ctx, contactId) })
}

func (c *SOAPClient) GetContactsAsync(ctx context.Context, groupId int, keyword string, from int, count int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetContacts(ctx, groupId, keyword, from, count) })
}

func (c *SOAPClient) GetContactsByIDAsync(ctx context.Context, contactId int, status int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetContactsByID(ctx, contactId, status) })
}

func (c *SOAPClient) GetGroupsAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetGroups(ctx) })
}

func (c *SOAPClient) MergeGroupsAsync(ctx context.Context, originGroupId int, destinationGroupId int, deleteOriginGroup bool) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.MergeGroups(ctx, originGroupId, destinationGroupId, deleteOriginGroup)
	})
}

func (c *SOAPClient) RemoveContactAsync(ctx context.Context, mobilenumber string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveContact(ctx, mobilenumber) })
}

func (c *SOAPClient) RemoveContactByContactIDAsync(ctx context.Context, contactId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveContactByContactID(ctx, contactId) })
}

func (c *SOAPClient) RemoveGroupAsync(ctx context.Context, groupId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveGroup(ctx, groupId) })
}

// ACtions

func (c *SOAPClient) AddBranchAsync(ctx context.Context, branchName string, owner int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.AddBranch(ctx, branchName, owner) })
}

func (c *SOAPClient) AddBulkAsync(ctx context.Context, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddBulk(ctx, from, branch, bulkType, title, message, rangeFrom, rangeTo, DateToSend, requestCount, rowFrom)
	})
}

func (c *SOAPClient) AddBulk2Async(ctx context.Context, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddBulk2(ctx, from, branch, bulkType, title, message, rangeFrom, rangeTo, DateToSend, requestCount, rowFrom)
	})
}

func (c *SOAPClient) AddNewBulkAsync(ctx context.Context, from string, branch int, bulkType int, title string, message string, rangeFrom string, rangeTo string, DateToSend string, requestCount int, rowFrom int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddNewBulk(ctx, from, branch, bulkType, title, message, rangeFrom, rangeTo, DateToSend, requestCount, rowFrom)
	})
}

func (c *SOAPClient) AddNumberAsync(ctx context.Context, branchId int, mobileNumbers []string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.AddNumber(ctx, branchId, mobileNumbers) })
}

func (c *SOAPClient) GetBranchsAsync(ctx context.Context, owner int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetBranchs(ctx, owner) })
}

func (c *SOAPClient) GetBulkAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetBulk(ctx) })
}

func (c *SOAPClient) GetBulkCountAsync(ctx context.Context, branch int, rangeFrom string, rangeTo string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetBulkCount(ctx, branch, rangeFrom, rangeTo) })
}

func (c *SOAPClient) GetBulkReceptionsAsync(ctx context.Context, bulkId int, fromRows int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetBulkReceptions(ctx, bulkId, fromRows) })
}

func (c *SOAPClient) GetBulkStatusAsync(ctx context.Context, bulkId int, sent int, failed int, status int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetBulkStatus(ctx, bulkId, sent, failed, status) })
}

func (c *SOAPClient) GetMobileCountAsync(ctx context.Context, branch int, rangeFrom string, rangeTo string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetMobileCount(ctx, branch, rangeFrom, rangeTo) })
}

func (c *SOAPClient) GetSendBulkAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetSendBulk(ctx) })
}

func (c *SOAPClient) GetTodaySentAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetTodaySent(ctx) })
}

func (c *SOAPClient) GetTotalSentAsync(ctx context.Context) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetTotalSent(ctx) })
}

func (c *SOAPClient) RemoveBranchAsync(ctx context.Context, branchId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveBranch(ctx, branchId) })
}

func (c *SOAPClient) RemoveBulkAsync(ctx context.Context, bulkId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveBulk(ctx, bulkId) })
}

func (c *SOAPClient) SendMultipleSMSAsync(ctx context.Context, to []string, from string, text []string, isflash bool, udh string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendMultipleSMS(ctx, to, from, text, isflash, udh) })
}

func (c *SOAPClient) SendMultipleSMS2Async(ctx context.Context, to []string, from []string, text []string, isflash bool, udh string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendMultipleSMS2(ctx, to, from, text, isflash, udh) })
}

func (c *SOAPClient) UpdateBulkDeliveryAsync(ctx context.Context, bulkId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.UpdateBulkDelivery(ctx, bulkId) })
}

// Schedule

func (c *SOAPClient) AddMultipleScheduleAsync(ctx context.Context, to []string, from string, text []string, isflash bool, scheduleDateTime []string, period string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddMultipleSchedule(ctx, to, from, text, isflash, scheduleDateTime, period)
	})
}

func (c *SOAPClient) AddNewUsanceAsync(ctx context.Context, to string, from string, text string, isflash bool, scheduleStartDateTime string, countrepeat int, scheduleEndDateTime string, periodType string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddNewUsance(ctx, to, from, text, isflash, scheduleStartDateTime, countrepeat, scheduleEndDateTime, periodType)
	})
}

func (c *SOAPClient) AddScheduleAsync(ctx context.Context, to string, from string, text string, isflash bool, scheduleDateTime string, period string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.AddSchedule(ctx, to, from, text, isflash, scheduleDateTime, period) })
}

func (c *SOAPClient) AddUsanceAsync(ctx context.Context, to string, from string, text string, isflash bool, scheduleStartDateTime string, repeatAfterDays int, scheduleEndDateTime string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.AddUsance(ctx, to, from, text, isflash, scheduleStartDateTime, repeatAfterDays, scheduleEndDateTime)
	})
}

func (c *SOAPClient) GetScheduleStatusAsync(ctx context.Context, scheduleId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetScheduleStatus(ctx, scheduleId) })
}

func (c *SOAPClient) RemoveScheduleAsync(ctx context.Context, scheduleId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveSchedule(ctx, scheduleId) })
}

func (c *SOAPClient) RemoveScheduleListAsync(ctx context.Context, scheduleIdList []int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.RemoveScheduleList(ctx, scheduleIdList) })
}

// Voice
func (c *SOAPClient) GetSendSMSWithSpeechTextStatusAsync(ctx context.Context, recId int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.GetSendSMSWithSpeechTextStatus(ctx, recId) })
}

func (c *SOAPClient) SendBulkSpeechTextAsync(ctx context.Context, title string, bbody string, receivers string, dateToSend string, repeatCount int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.SendBulkSpeechText(ctx, title, bbody, receivers, dateToSend, repeatCount)
	})
}

func (c *SOAPClient) SendBulkVoiceSMSAsync(ctx context.Context, title string, voiceFileId int, receivers string, DateToSend string, repeatCount int) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.SendBulkVoiceSMS(ctx, title, voiceFileId, receivers, DateToSend, repeatCount)
	})
}

func (c *SOAPClient) SendSMSWithSpeechTextAsync(ctx context.Context, smsBody string, speechBody string, from string, to string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.SendSMSWithSpeechText(ctx, smsBody, speechBody, from, to) })
}

func (c *SOAPClient) SendSMSWithSpeechTextBySchduleDateAsync(ctx context.Context, smsBody string, speechBody string, from string, to string, scheduleDate string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) {
		return c.SendSMSWithSpeechTextBySchduleDate(ctx, smsBody, speechBody, from, to, scheduleDate)
	})
}

func (c *SOAPClient) UploadVoiceFileAsync(ctx context.Context, title string, base64StringFile string) <-chan Result[[]byte] {
	return Async(func() ([]byte, error) { return c.UploadVoiceFile(ctx, title, base64StringFile) })
}
