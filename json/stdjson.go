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

func parseStdJson(fileName string) {
	var testStruct TestStruct
	f, _ := os.Open(fileName)
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			line = strings.TrimSpace(line)
			fmt.Println(testStruct.Category, testStruct.Value)
			err = json.Unmarshal([]byte(line), &testStruct)
			break
		}
		line = strings.TrimSpace(line)
		err = json.Unmarshal([]byte(line), &testStruct)
		fmt.Println(testStruct.Category, testStruct.Value)
	}
}

func genJson() {
	var testStruct TestStruct
	testStruct.Category = "a"
	testStruct.Value = 1000
	line, _ := json.Marshal(testStruct)
	fmt.Println(string(line))
}
