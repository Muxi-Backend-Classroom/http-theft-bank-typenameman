package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	_ "io"
	"io/ioutil"
	_ "mime/multipart"
	"net/http"
	_ "os"
	"strings"
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
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTIwIiwiaWF0IjoxNjM3MTQ2NDM1LCJuYmYiOjE2MzcxNDY0MzV9.mlggHuQMg4eooV1KBB9scFQE-J7018S5RLpXl-boWX4"
	client := &http.Client{}
	imfor := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	im, _ := base64.StdEncoding.DecodeString(imfor)
	im = im //纯粹为了把它用掉，golang很蠢的一个地方
	//fmt.Println(string(im))，得到两个字符串
	secret_key := "MuxiStudio203304"
	error_code := "for {go func(){time.Sleep(1*time.Hour)}()}"
	//加密error_code

	plaintext := []byte(error_code)
	key := []byte(secret_key)
	ivT := make([]byte, aes.BlockSize+len(plaintext))
	iv := ivT[:aes.BlockSize]
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padtext...)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	error := base64.StdEncoding.EncodeToString(crypted)
	//获得加密的error_code，并破门而入！

	type Body struct {
		Content string
	}
	body := Body{
		Content: error,
	}
	by, _ := json.Marshal((body))
	payload := strings.NewReader(string(by))
	req, _ := http.NewRequest("PUT", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate", payload)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ := client.Do(req)
	//fmt.Println(res.Header)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
