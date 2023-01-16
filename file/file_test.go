package file

import "testing"

/*
File name    : file_test.go
Author       : miaoyc
Create Date  : 2022/12/29 11:23
Update Date  : 2022/12/29 11:23
Description  :
*/

func TestIsExists(t *testing.T) {
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
			if got := IsExists(tt.args.filePath); got != tt.want {
				t.Errorf("IsExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveFile(t *testing.T) {
	type args struct {
		filePath string
		content  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"save-file-1", args{"./test.md", "miaoyc no.1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveFile(tt.args.filePath, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("SaveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"delete file-1", args{"./test.md"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteFile(tt.args.filePath)
		})
	}
}
