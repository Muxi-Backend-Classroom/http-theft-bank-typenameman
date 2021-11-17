package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
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

	fmt.Println(req)

	// write your code below
	// ...*/
	client := &http.Client{}
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTIwIiwiaWF0IjoxNjM3MTQ2NDM1LCJuYmYiOjE2MzcxNDY0MzV9.mlggHuQMg4eooV1KBB9scFQE-J7018S5RLpXl-boWX4"
	req, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key", nil)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ := client.Do(req)
	//fmt.Println(res.Header)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
