package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "test string 24235514313#$%#$^&%&"
	encode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encode)
	decode, _ := base64.StdEncoding.DecodeString(encode)
	fmt.Println(string(decode))
}
