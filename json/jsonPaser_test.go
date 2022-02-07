package json

import (
	"testing"
)

/*
File name    : jsonPaser_test.go
Author       : miaoyc
Create date  : 2022/1/14 5:44 下午
Description  :
*/

var fileName string

func init() {
	fileName = "/Users/miaoyc/code/github/goArsenal/json/case.txt"
}

func TestFastjson(t *testing.T) {
	fastjsonParser(fileName)
}

func TestGJson(t *testing.T) {
	gjsonParser(fileName)
}

func TestStdJson(t *testing.T) {
	parseStdJson(fileName)
	genJson()
}
