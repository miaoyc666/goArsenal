package re

import (
	"github.com/dlclark/regexp2"
	"net/mail"
	"regexp"
)

/*
File name    : re.go
Author       : miaoyc
Create date  : 2022/10/31 14:21
Description  :
*/

var mailReg *regexp.Regexp
var mailReg2 *regexp2.Regexp

func init() {
	mailPattern := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	mailReg = regexp.MustCompile(mailPattern)
	mailReg2, _ = regexp2.Compile(mailPattern, 0)
}

func ExampleParseAddress(email string) (string, string, error) {
	e, err := mail.ParseAddress(email)
	if err != nil {
		return "", "", err
	}
	return e.Name, e.Address, nil
}

func VerifyEmailFormat(email string) bool {
	return mailReg.MatchString(email)
}

func Reg2VerifyEmailFormat(email string) bool {
	isMatch, _ := mailReg2.MatchString(email)
	return isMatch
}
