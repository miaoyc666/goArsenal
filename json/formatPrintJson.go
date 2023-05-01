package json

import (
	"bytes"
	"encoding/json"
	"strings"
)

/*
File name    : formatPrintJson.go
Author       : miaoyc
Create Date  : 2022/11/9 15:01
Update Date  : 2022/11/9 15:01
Description  :
*/

func FormatPrintJson(s string) (string, error) {
	s1 := strings.Replace(s, `\"`, `"`, -1)
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(s1), "", "\t")
	if err != nil {
		return "", err
	}
	formatJson := string(prettyJSON.Bytes())
	return formatJson, nil
}
