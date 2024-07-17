package time

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

/*
File name    : time.go
Author       : miaoyc
Create time  : 2024/7/17 17:54
Update time  : 2024/7/17 17:54
Description  :
*/

// ParseAndStandardizeTime 函数接收一个时间字符串，尝试解析它，并将其转换为最常用的时间格式（time.DateTime）
func ParseAndStandardizeTime(timeString string) (string, error) {
	t, err := dateparse.ParseAny(timeString)
	if err != nil {
		return "", fmt.Errorf("无法解析时间字符串 %s: %w", timeString, err)
	}
	// 将时间转换为标准格式
	standardTime := t.Format(time.DateTime)
	return standardTime, nil
}
