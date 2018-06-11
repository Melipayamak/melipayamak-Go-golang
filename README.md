# melipayamak Go (golang)

<div dir='rtl'>

### معرفی وب سرویس ملی پیامک
ملی پیامک یک وب سرویس کامل برای ارسال و دریافت پیامک و پیامک صوتی و مدیریت کامل خدمات دیگر است که براحتی میتوانید از آن استفاده کنید.

<hr>

### نصب

<p>قبل از نصب نیاز به ثبت نام در سایت ملی پیامک دارید.</p>

[**ثبت نام به همراه دریافت 200 پیامک هدیه جهت تست وبسرویس**](http://www.melipayamak.com/)

</div>

<div dir='rtl'>
  
#### نحوه استفاده

نمونه کد برای ارسال پیامک

</div>


```js
username := "username"
password := "password"
from := "5000..."
to := "09123456789"
text := "تست وب سرویس ملی پیامک"
isFlash := false
soapClient := SoapClient(username, password)
soapClient.SendSimpleSMS2(to: to, from: from, msg: text, isFlash: isFlash)
//یا برای ارسال به مجموعه ای از مخاطبین
soapClient.SendSimpleSMS([]string, from, text, isFlash)
```

<div dir='rtl'>

از آنجا که وب سرویس ملی پیامک تنها محدود به ارسال پیامک نیست شما از طریق زیر میتوانید به وب سرویس ها دسترسی کامل داشته باشید:
</div>

```js
// وب سرویس پیامک
restClient := RestClient(username, password)
soapClient := SoapClient(username, password)
```

<div dir='rtl'>
  
#### تفاوت های وب سرویس پیامک rest و soap

از آنجا که ملی پیامک وب سرویس کاملی رو در اختیار توسعه دهندگان میگزارد برای راحتی کار با وب سرویس پیامک علاوه بر وب سرویس اصلی soap وب سرویس rest رو هم در اختیار توسعه دهندگان گزاشته شده تا راحتتر بتوانند با وب سرویس کار کنند. تفاوت اصلی این دو در تعداد متد هاییست که میتوانید با آن کار کنید. برای کار های پایه میتوان از وب سرویس rest استفاده کرد برای دسترسی بیشتر و استفاده پیشرفته تر نیز باید از وب سرویس باید از وب سرویس soap استفاده کرد. جهت مطالعه بیشتر وب سرویس ها به قسمت وب سرویس پنل خود مراجعه کنید.

<hr/>


###  اطلاعات بیشتر

برای مطالعه بیشتر و دریافت راهنمای وب سرویس ها و آشنایی با پارامتر های ورودی و خروجی وب سرویس به صفحه معرفی
[وب سرویس ملی پیامک](https://github.com/Melipayamak/Webservices)
مراجعه نمایید .


<hr/>

</div>


<div dir='rtl'>

### وب سرویس پیامک

متدهای وب سرویس:

</div>

#### ارسال

```js
restClient.Send(to: to, from: from, msg: text, isFlash: isFlash)
soapClient.SendSimpleSMS(array: [to1, to2], from: from, msg: text, isFlash: isFlash)
```
<div dir='rtl'>
  در آرگومان سوم روش soap میتوانید از هر تعداد مخاطب به عنوان آرایه استفاده کنید
</div>

#### دریافت وضعیت ارسال
```js
restClient.GetDelivery(id: recId)
soapClient.GetDelivery(id: recId)
soapClient.GetDeliveries(array: [id1, id2])
```

#### لیست پیامک ها

```js
restClient.GetMessages(loc: location, indx: index, cnt: count, from: from)
soapClient.getMessages(loc: location, from: from, indx: index, cnt: count)
// جهت دریافت به صورت رشته ای
soapClient.GetMessagesByDate(loc: location, from: from, indx: index, cnt: count, dfrom: dateFrom, dto: dateTo)
//جهت دریافت بر اساس تاریخ
soapClient.GetUsersMessagesByDate(loc: location, from: from, indx: index, cnt: count, dfrom: dateFrom, dto: dateTo)
// جهت دریافت پیام های کاربران بر اساس تاریخ 
```

#### موجودی
```js
restClient.GetCredit()
soapClient.GetCredit()
```

#### تعرفه پایه / دریافت قیمت قبل از ارسال
```js
restClient.GetBasePrice()
soapClient.GetSmsPrice(irancellCnt: irancellCount, mtnCnt: mtnCount, from: from, msg: text)
```
#### لیست شماره اختصاصی
```js
soapClient.GetUserNumbers()
```

#### بررسی تعداد پیامک های دریافتی
```js
soapClient.GetInboxCount(isRead: isRead)
//پیش فرض خوانده نشده 
```

#### ارسال پیامک پیشرفته
```js
soapClient.SendSms(arrayOfTo: [to1, to2], from: from, msg: text, isFlash: isflash, uhd: udh, arrayOfRecId: [id1, id2], arrayOfStatus: [st1, st2])
```

#### مشاهده مشخصات پیام
```js
soapClient.GetMessagesReceptions(msgId: msgId, fromRows: fromRows)
```


#### حذف پیام دریافتی
```js
soapClient.RemoveMessages2(loc: location, msgIds: msgIds)
```


#### ارسال زماندار
```js
soapClient.AddSchedule(to: to, from: from, msg: text, isFlash: isflash, sdt: scheduleDateTime, prd: period)
```

#### ارسال زماندار متناظر
```js
soapClient.AddMultipleSchedule(array: [to1, to2], from: from, arrayOfMsg: [msg1, msg2], isFlash: isflash, arrayOfSt: [sche1, sche2], prd: period)
```


#### ارسال سررسید
```js
soapClient.AddNewUsance(to :to, from: from, msg: text, isFlash: isflash, ss: scheduleStartDateTime, cntrpt: countRepeat, se: scheduleEndDateTime, prdType: periodType)
```

#### مشاهده وضعیت ارسال زماندار
```js
soapClient.GetScheduleStatus(id: schId)
```

#### حذف پیامک زماندار
```js
soapClient.RemoveSchedule(id: schId)
```


####  ارسال پیامک همراه با تماس صوتی
```js
soapClient.SendSMSWithSpeechText(body: smsBody, spch: speechBody, from: from, to: to)
```

####  ارسال پیامک همراه با تماس صوتی به صورت زمانبندی
```js
soapClient.SendSMSWithSpeechTextBySchduleDate(body: smsBody, spch: speechBody, from: from, to: to, sche: scheduleDate)
```

####  دریافت وضعیت پیامک همراه با تماس صوتی 
```js
soapClient.GetSendSMSWithSpeechTextStatus(id: recId)
```
<div dir='rtl'>
  
### وب سرویس ارسال انبوه/منطقه ای

</div>

#### دریافت شناسه شاخه های بانک شماره
```js
soapClient.GetBranchs(owner: owner)
```


#### اضافه کردن یک بانک شماره جدید
```js
soapClient.AddBranch(name: branchName, owner: owner)
```

#### اضافه کردن شماره به بانک
```js
soapClient.AddNumber(id: branchId, array: [num1, num2])
```

#### حذف یک بانک
```js
soapClient.RemoveBranch(id: branchId)
```

#### ارسال انبوه از طریق بانک
```js
soapClient.AddBulk(from: from, branch: branch, bulkType: bulkType, title: title, msg: message, rangeF: rangeFrom, rangeT: rangeTo, date: DateToSend, reqCnt: requestCount, rowFrom: rowFrom)
```

#### تعداد شماره های موجود
```js
soapClient.GetBulkCount(branch: branch, rangeF: rangeFrom, rangeT: rangeTo)
```

#### گزارش گیری از ارسال انبوه
```js
soapClient.GetBulkReceptions(id: bulkId, fromRows: fromRows)
```


#### تعیین وضعیت ارسال 
```js
soapClient.GetBulkStatus(id: bulkId, sent: sent, failed: failed, status: status)
```

#### تعداد ارسال های امروز
```js
soapClient.GetTodaySent()
```

#### تعداد ارسال های کل

```js
soapClient.GetTotalSent()
```

#### حذف ارسال منطقه ای
```js
soapClient.RemoveBulk(id: bulkId)
```

#### ارسال متناظر
```js
soapClient.SendMultipleSMS(array: [to1, to2], from: from, arrayOfMsg: [msg1, msg2], isFlash: isflash, uhd: udh, arrayOfRecIds: [id1, id2], status: status)
```

#### نمایش دهنده وضعیت گزارش گیری

```js
soapClient.UpdateBulkDelivery(id: bulkId)
```
<div dir='rtl'>
  
### وب سرویس تیکت

</div>

#### ثبت تیکت جدید
```js
soapClient.AddTicket(title: title, content: content, alert: aletWithSms)
```

#### جستجو و دریافت تیکت ها

```js
soapClient.GetReceivedTickets(owner: ticketOwner, type: ticketType, keyword: keyword)
```

#### دریافت تعداد تیکت های کاربران
```js
soapClient.GetReceivedTicketsCount(type: ticketType)
```

#### دریافت تیکت های ارسال شده
```js
soapClient.GetSentTickets(owner: ticketOwner, type: ticketType, keyword: keyword)
```

#### دریافت تعداد تیکت های ارسال شده
```js
soapClient.GetSentTicketsCount(type: ticketType)
```


#### پاسخگویی به تیکت
```js
soapClient.ResponseTicket(id: ticketId, type: type, content: content, alert: alertWithSms)
```
<div dir='rtl'>
  
### وب سرویس دفترچه تلفن

</div>

#### اضافه کردن گروه جدید
```js
soapClient.AddGroup(name: groupName, desc: Descriptions, show: showToChilds)
```

#### اضافه کردن کاربر جدید
```js
soapClient.AddContact(options)

```

#### بررسی موجود بودن شماره در دفترچه تلفن
```js
soapClient.CheckMobileExistInContact(number: mobileNumber)
```

#### دریافت اطلاعات دفترچه تلفن
```js
soapClient.GetContacts(id: groupId, keyword: keyword, cnt: count)
```
#### دریافت گروه ها
```js
soapClient.GetGroups()
```
#### ویرایش مخاطب
```js
soapClient.ChangeContact(options)
```

#### حذف مخاطب
```js
soapClient.RemoveContact(number: mobilenumber)
```
#### دریافت اطلاعات مناسبت های فرد
```js
soapClient.GetContactEvents(id: contactId)
```

<div dir='rtl'>

### وب سرویس کاربران

</div>

#### ثبت فیش واریزی
```js
soapClient.AddPayment(options)
```

#### اضافه کردن کاربر جدید در سامانه
```js
soapClient.AddUser(options)

```

#### اضافه کردن کاربر جدید در سامانه(کامل)
```js
soapClient.AddUserComplete(options)
```

#### اضافه کردن کاربر جدید در سامانه(WithLocation)
```js
soapClient.AddUserWithLocation(options)
```
#### بدست آوردن ID کاربر
```js
soapClient.AuthenticateUser()
```
#### تغییر اعتبار
```js
soapClient.ChangeUserCredit(amount: amount, desc: description, user: targetUsername, gTax: GetTax)
```

#### فراموشی رمز عبور
```js
soapClient.ForgotPassword(number: mobileNumber, email: emailAddress, user: targetUsername)
```
#### دریافت تعرفه پایه کاربر
```js
soapClient.GetUserBasePrice(user: targetUsername)
```

#### دریافت اعتبار کاربر
```js
soapClient.GetUserCredit(user: targetUsername)
```

#### دریافت مشخصات کاربر
```js
soapClient.GetUserDetails(user: targetUsername)
```

#### دریافت شماره های کاربر
```js
soapClient.GetUserNumbers()
```

#### دریافت تراکنش های کاربر
```js
soapClient.GetUserTransactions(user: targetUsername, creditType: creditType, dateF: dateFrom, dateT: dateTo, keyword: keyword)
```

#### دریافت اطلاعات  کاربران
```js
soapClient.GetUsers()
```


#### دریافت اطلاعات  فیلترینگ
```js
soapClient.HasFilter(txt: text)
```


####  حذف کاربر
```js
soapClient.RemoveUser(user: targetUsername)
```


#### مشاهده استان ها
```js
soapClient.GetProvinces()
```

#### مشاهده کد شهرستان 
```js
soapClient.GetCities(id: provinceId)
```


#### مشاهده تاریخ انقضای کاربر 
```js
soapClient.GetExpireDate()
```
