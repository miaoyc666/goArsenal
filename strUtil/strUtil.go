package strUtil

import (
	"regexp"
	"strings"
)

/*
File name    : strUtil.go
Author       : miaoyc
Create Date  : 2024/5/19 23:58
Update Date  : 2024/6/30 16:10
Description  :
*/

var (
	md5Pattern  = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)
	uuidPattern = regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}$`)
	namePattern = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
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
