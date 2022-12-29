package file

import "testing"

/*
File name    : file_test.go
Author       : miaoyc
Create Date  : 2022/12/29 11:23
Update Date  : 2022/12/29 11:23
Description  :
*/

func TestFileExists(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"fileExist-1", args{"./file.go"}, true},
		{"fileExist-2", args{"./file1.go"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExists(tt.args.filePath); got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
