package json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
File name    : stdjson.go
Author       : miaoyc
Create date  : 2021/12/28 11:27 上午
Description  :
*/

// {"category": "b", "value": 6}

type TestStruct struct {
	Category string `json:"category"`
	Value    int    `json:"value"`
}

func RunStdJson() {
	fileName := "/home/miaoyongchao/test/case.txt"
	var testStruct TestStruct
	f, _ := os.Open(fileName)
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimSpace(line)
		err = json.Unmarshal([]byte(line), &testStruct)
		fmt.Println(testStruct.Category, testStruct.Value)
	}
}
