package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	/*req, err := httptool.NewRequest(
		"",
		"",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(req)*/

	// write your code below
	// ...
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code", nil)
	req.Header.Add("Code", "120")
	res, _ := client.Do(req)
	by, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(by))
	/*passport := res.Header["Passport"][0]
	req, _ = http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key", nil)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ = client.Do(req)
	fmt.Println(res.Header)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))*/
}
