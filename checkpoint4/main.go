package main

import (
	"bytes"
	_ "crypto/aes"
	_ "crypto/cipher"
	"encoding/base64"
	_ "encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	_ "strings"
)

//"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"

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
	req, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate", nil)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ := client.Do(req)
	//fmt.Println(res.Header)
	content, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(content))
	//图片转码

	req, _ = http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample", nil)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ = client.Do(req)
	//fmt.Println(res.Header)
	content, _ = ioutil.ReadAll(res.Body)
	//fmt.Println(string(content))
	f, _ := os.Open("./t.txt")
	var b = make([]byte, 136344)
	f.Read(b)
	imageBytes, _ := base64.StdEncoding.DecodeString(string(b))
	ioutil.WriteFile("./test.jpg", imageBytes, 0777)
	//输入图片访问

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodywriter, _ := bodyWriter.CreateFormFile("file", "test.jpg")
	file, _ := os.Open("./test.jpg")
	io.Copy(bodywriter, file)
	bodyWriter.Close() //此处特别注意
	req, _ = http.NewRequest("POST", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/iris_recognition_gate", bodyBuf)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	res, _ = client.Do(req)
	//fmt.Println(res.Header)
	content, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
