package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "strconv"
    "io/ioutil"
    "net/http"
)

func makeRequest(jsonData map[string]string, op string) {

	jsonValue, _ := json.Marshal(jsonData)
    response, err := http.Post("https://rest.payamak-panel.com/api/SendSMS/" + op, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}

func main() {
	username := "9122088891"
	password := "1234"
	to := "09103155711"
	from := "500010604095"
	text := "Go Rest Test"
	isFlash := false

    jsonData := map[string]string {
    	"username": username,
    	"password": password, 
    	"to" : to, 
    	"from" : from, 
    	"text" : text, 
    	"isFlash" : strconv.FormatBool(isFlash),
    }
    makeRequest(jsonData, "SendSMS")
}

