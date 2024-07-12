package crypto

import "encoding/base64"

/*
File name    : base64.go
Author       : miaoyc
Create time  : 2024/7/12 11:18
Update time  : 2024/7/12 11:18
Description  :
*/

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(encodedData string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodedData)
}
