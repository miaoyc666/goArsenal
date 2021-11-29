package json

/*
File name    : json.go
Author       : miaoyc
Create date  : 2021/11/19 11:39 下午
Description  : 解析json示例，非官方库
*/

import (
	"bufio"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"os"
	"strings"
)

/*
原始文件case.json为每行都是json数据的文本
*/

func Run() {
	fileName := "case.text"
	f, _ := os.Open(fileName)
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimSpace(line)
		category := gjson.Get(line, "category")
		value := gjson.Get(line, "value")
		fmt.Println(category, value)
	}
}
