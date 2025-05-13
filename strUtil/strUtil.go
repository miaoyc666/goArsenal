package strUtil

import (
	"regexp"
	"strings"
	"unicode"
)

/*
File name    : strUtil.go
Author       : miaoyc
Create Date  : 2024/5/19 23:58
Update Date  : 2024/8/13 23:49
Description  :
*/

var (
	md5Pattern   = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
	uuidPattern  = regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}$`)
	namePattern  = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	emailPattern = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

func Concat(str1, str2 string, strs ...string) string {
	var builder strings.Builder
	builder.WriteString(str1)
	builder.WriteString(str2)
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

func IsUUID(uuid string) bool {
	return uuidPattern.MatchString(uuid)
}

func IsMD5(str string) bool {
	return md5Pattern.MatchString(str)
}

// IsValidName 是否是有效的名称, 匹配数字、字母、中划线和下划线
func IsValidName(s string) bool {
	return namePattern.MatchString(s)
}

func IsNumeric(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsMobile(phoneNumber string) bool {
	if len(phoneNumber) != 11 {
		return false
	}
	if phoneNumber[0] != '1' {
		return false
	}
	return IsNumeric(phoneNumber)
}

func IsEmail(email string) bool {
	return emailPattern.MatchString(email)
}

// IsAllInvisibleOrSpace 判断字符串是否全部由不可见字符或空白字符组成
func IsAllInvisibleOrSpace(s string) bool {
	if len(s) == 0 {
		return true
	}

	for _, r := range s {
		// 如果存在可见字符（非空白且非控制字符），返回 false
		if !unicode.IsSpace(r) && !unicode.IsControl(r) {
			return false
		}
	}
	return true
}
