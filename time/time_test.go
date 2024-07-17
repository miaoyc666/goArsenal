package time

import (
	"fmt"
	"testing"
)

/*
File name    : time_test.go
Author       : miaoyc
Create time  : 2024/7/17 17:55
Update time  : 2024/7/17 17:55
Description  :
*/

func TestParseAndStandardizeTime(t *testing.T) {
	timeStrings := []string{
		"2023-04-30",
		"30/04/2023",
		"April 30, 2023",
		"2023-04-30T15:04:05Z",
		"30-04-2023 15:04:05",
	}

	for _, ts := range timeStrings {
		standardTime, err := ParseAndStandardizeTime(ts)
		if err != nil {
			fmt.Printf("解析时间字符串 %s 时出错: %v", ts, err)
		}
		fmt.Printf("原始时间: %s, 标准时间: %s\n", ts, standardTime)
	}
}
