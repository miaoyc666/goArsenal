package versionDiff

import "testing"

/*
File name    : version_test.go
Author       : miaoyc
Create date  : 2022/1/14 6:03 下午
Description  :
*/

func TestVersionDiff(t *testing.T) {
	a := "2021.11.23.1537"
	b := "2021.10.24.1537"
	versionDiff(a, b)
}
