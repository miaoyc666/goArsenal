package network

import "testing"

/*
File name    : network_test.go
Author       : miaoyc
Create date  : 2022/10/29 20:31
Description  :
*/

func TestIsIp(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
		{"ipv4-true", args{"192.168.88.1"}, true, "ipv4"},
		{"ipv4-false", args{"192.168.88.666"}, false, ""},
		{"ipv6-true", args{"2606:4700:90:4fb:aeec:54b4:a3ad:6688"}, true, "ipv6"},
		{"ipv6-false", args{"[2606:4700:90:4fb:aeec:54b4:a3ad:6688]:1000"}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsIp(tt.args.ip)
			if got != tt.want {
				t.Errorf("IsIp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsIp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
