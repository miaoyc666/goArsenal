package json

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/valyala/fastjson"
)

/*
File name    : fastjson.go
Author       : miaoyc
Create date  : 2021/12/28 11:27 上午
Description  : fastjson - fast JSON parser and validator for Go
*/

func parser(fileName string) {
	f, _ := os.Open(fileName)
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimSpace(line)
		var p fastjson.Parser
		v, _ := p.Parse(line)
		fmt.Printf("Category=%s\n", v.GetStringBytes("category"))
		fmt.Printf("Value=%d\n", v.GetInt("value"))
	}
}

func main() {
	fileName := "/home/miaoyongchao/test/case.txt"
	parser(fileName)
}
