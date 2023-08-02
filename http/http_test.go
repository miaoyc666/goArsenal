package http

import (
	"net/http"
	"reflect"
	"testing"
)

/*
File name    : http_test.go
Author       : miaoyc
Create Date  : 2023/8/2 18:21
Update Date  : 2023/8/2 18:21
Description  :
*/

func TestNewHttpClient(t *testing.T) {
	type args struct {
		proxy               string
		insecureSkipVerify_ bool
	}
	tests := []struct {
		name string
		args args
		want *http.Client
	}{
		// TODO: Add test cases.
		{"new client", args{"192.168.88.100:7890", true}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpClient(tt.args.proxy, tt.args.insecureSkipVerify_); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
