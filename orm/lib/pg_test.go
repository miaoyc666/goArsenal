package lib

import "testing"

/*
File name    : pg_test.go
Author       : miaoyc
Create Date  : 2023/5/1 00:18
Update Date  : 2023/5/1 00:18
Description  :
*/

func TestInitPGConnection(t *testing.T) {
	type args struct {
		host     string
		port     string
		user     string
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"init", args{
			host:     "",
			port:     "",
			user:     "",
			password: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitPGConnection(tt.args.host, tt.args.port, tt.args.user, tt.args.password)
		})
	}
}
