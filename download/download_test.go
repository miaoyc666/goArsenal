package download

import "testing"

/*
File name    : download_test.go
Author       : miaoyc
Create Date  : 2022/12/29 11:17
Update Date  : 2022/12/29 11:17
Description  :
*/

func Test_downloadFile(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"download-1", args{"https://github.com/miaoyc666/goArsenal/blob/master/README.md"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloadFile(tt.args.url)
		})
	}
}
