package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
File name    : formatPrintJson.go
Author       : miaoyc
Create Date  : 2022/11/9 15:01
Update Date  : 2022/11/9 15:01
Description  :
*/

func main() {
	s := os.Args[1]
	s1 := strings.Replace(s, `\"`, `"`, -1)
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(s1), "", "\t")
	if err != nil {
		fmt.Println("JSON parse error: ", err)
		fmt.Println(s1)
		return
	}
	fmt.Println(string(prettyJSON.Bytes()))
}

