package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodywriter, _ := bodyWriter.CreateFormFile("file", "permute.go")
	file, _ := os.Open("./permute.go")
	io.Copy(bodywriter, file)
	bodyWriter.Close()
	req, _ := http.NewRequest("POST", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination", bodyBuf)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	res, _ := client.Do(req)
	contents, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(contents))
}
