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
username := "username";
password := "password";
from := "5000...";
to := "09123456789";
text := "تست وب سرویس ملی پیامک";
isFlash := false;
SendSimpleSMS2(username, password, to, from, text, isFlash);
//یا برای ارسال به مجموعه ای از مخاطبین
SendSimpleSMS(username, password, to []string, from, text, isFlash);
```

<div dir='rtl'>

از آنجا که وب سرویس ملی پیامک تنها محدود به ارسال پیامک نیست شما از طریق متدهای زیر میتوانید به وب سرویس ها دسترسی کامل داشته باشید

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
Send(to, from, text, isFlash);
SendSimpleSMS(to []string, from, text, isFlash);
```
<div dir='rtl'>
  در آرگومان سوم روش soap میتوانید از هر تعداد مخاطب به عنوان آرایه استفاده کنید
</div>

ticketSoapClient.ResponseTicket(username, password, ticketId, type, content, alertWithSms);
```
<div dir='rtl'>
  
### وب سرویس دفترچه تلفن

</div>

#### اضافه کردن گروه جدید
```js
AddGroup(username, password, groupName, Descriptions, showToChilds);
```

#### اضافه کردن کاربر جدید
```js
AddContact(username, password, options);

```

#### بررسی موجود بودن شماره در دفترچه تلفن
```js
CheckMobileExistInContact(username, password, mobileNumber);
```

#### دریافت اطلاعات دفترچه تلفن
```js
GetContacts(username, password, groupId, keyword, count);
```
#### دریافت گروه ها
```js
GetGroups(username, password);
```
#### ویرایش مخاطب
```js
ChangeContact(username, password, options);
```

#### حذف مخاطب
```js
RemoveContact(username, password, mobilenumber);
```
#### دریافت اطلاعات مناسبت های فرد
```js
GetContactEvents(username, password, contactId);
```

<div dir='rtl'>

### وب سرویس کاربران

</div>

#### ثبت فیش واریزی
```js
AddPayment(username, password, options);
```

#### اضافه کردن کاربر جدید در سامانه
```js
AddUser(username, password, options);

```

#### اضافه کردن کاربر جدید در سامانه(کامل)
```js
AddUserComplete(username, password, options);
```

#### اضافه کردن کاربر جدید در سامانه(WithLocation)
```js
AddUserWithLocation(username, password, options);
```
#### بدست آوردن ID کاربر
```js
AuthenticateUser(username, password);
```
#### تغییر اعتبار
```js
ChangeUserCredit(username, password, amount, description, targetUsername, GetTax);
```

#### فراموشی رمز عبور
```js
ForgotPassword(username, password, mobileNumber, emailAddress, targetUsername);
```
#### دریافت تعرفه پایه کاربر
```js
GetUserBasePrice(username, password, targetUsername);
```

#### دریافت اعتبار کاربر
```js
GetUserCredit(username, password, targetUsername);
```

#### دریافت مشخصات کاربر
```js
GetUserDetails(username, password, targetUsername);
```

#### دریافت شماره های کاربر
```js
GetUserNumbers(username, password);
```

#### دریافت تراکنش های کاربر
```js
GetUserTransactions(username, password, targetUsername, creditType, dateFrom, dateTo, keyword);
```

#### دریافت اطلاعات  کاربران
```js
GetUsers(username, password);
```


#### دریافت اطلاعات  فیلترینگ
```js
HasFilter(username, password, text);
```


####  حذف کاربر
```js
RemoveUser(username, password, targetUsername);
```


#### مشاهده استان ها
```js
GetProvinces(username, password);
```

#### مشاهده کد شهرستان 
```js
GetCities(username, password, provinceId);
```


#### مشاهده تاریخ انقضای کاربر 
```js
GetExpireDate(username, password);
```
