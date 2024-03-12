package crypto

import (
	"crypto/sha1"
	"fmt"
)

/*
File name    : sha1.go
Author       : miaoyc
Create Date  : 2024/3/12 23:49
Update Date  : 2024/3/12 23:49
Description  :
*/

func Sha1(str string) (sha1str string) {
	h := sha1.New()
	_, _ = h.Write([]byte(str))
	sha1str = fmt.Sprintf("%x", h.Sum(nil))
	return sha1str
}
