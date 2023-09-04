package file

import "testing"

/*
File name    : file_test.go
Author       : miaoyc
Create Date  : 2022/12/29 11:23
Update Date  : 2022/12/29 11:23
Description  :
*/

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

func TestIsDirExists(t *testing.T) {
	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test folder", args{"/Users/miaoyongchao/code/github/goArsenal/tmp"}, true, false},
		{"test folder not found", args{"/Users/miaoyongchao/code/github/goArsenal/tmp1"}, false, true},
		{"test folder is file", args{"/Users/miaoyongchao/code/github/goArsenal/tmp/1"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDirExists(tt.args.folderPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PathExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClearDir(t *testing.T) {
	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"clear file", args{"/Users/miaoyongchao/code/github/goArsenal/tmp/1"}, true},
		{"clear folder", args{"/Users/miaoyongchao/code/github/goArsenal/tmp"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ClearDir(tt.args.folderPath); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsFileExists(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"fileExist-1", args{"/Users/miaoyongchao/code/github/goArsenal/tmp/1"}, true, false},
		{"fileExist-2", args{"/Users/miaoyongchao/code/github/goArsenal/tmp"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFileExists(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFileExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFileExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteDir(t *testing.T) {
	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"delete-1", args{"/Users/miaoyongchao/code/github/goArsenal/tmp"}, false},
		{"delete-2", args{"/Users/miaoyongchao/code/github/goArsenal/tmp1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteDir(tt.args.folderPath); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
