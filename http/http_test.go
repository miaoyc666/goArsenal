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
		params TransportParams
	}
	params := NewTransportParams()
	params.Proxy = "https://proxy.happycode.fun:3129"
	params.CaCertFile = "/Users/miaoyc/tmp/proxy.happycode.fun.pem"

	tests := []struct {
		name string
		args args
		want *http.Client
	}{
		// TODO: Add test cases.
		{"new client", args{*params}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpClient(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
