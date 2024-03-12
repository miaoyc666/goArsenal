package crypto

import "testing"

/*
File name    : sha1_test.go
Author       : miaoyc
Create Date  : 2024/3/12 23:51
Update Date  : 2024/3/12 23:51
Description  :
*/

func TestSha1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name        string
		args        args
		wantSha1str string
	}{
		// TODO: Add test cases.
		{"sha1-1", args{""}, "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"sha1-2", args{"miaoyc"}, "820a3d7b99f7b460c3c24e70ddbd16fd1fb9eb5d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSha1str := Sha1(tt.args.str); gotSha1str != tt.wantSha1str {
				t.Errorf("Sha1() = %v, want %v", gotSha1str, tt.wantSha1str)
			}
		})
	}
}
