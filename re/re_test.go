package re

import "testing"

/*
File name    : re_test.go
Author       : miaoyc
Create date  : 2022/10/31 14:23
Description  :
*/

func BenchmarkExampleParseAddress(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ExampleParseAddress("9e9d7@1732205139ab.com.")
	}
}

func TestExampleParseAddress(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"valid", args{"9e9d7@1732205139ab.com"}, "", "9e9d7@1732205139ab.com", false},
		{"invalid1", args{"9e9d7@1732205139ab.com."}, "", "", true},
		{"invalid2", args{".9e9d7@1732205139ab.com"}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ExampleParseAddress(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExampleParseAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExampleParseAddress() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExampleParseAddress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestVerifyEmailFormat(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"valid", args{"9e9d7@1732205139ab.com"}, true},
		{"invalid1", args{"9e9d7@1732205139ab.com."}, false},
		{"invalid2", args{".9e9d7@1732205139ab.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyEmailFormat(tt.args.email); got != tt.want {
				t.Errorf("VerifyEmailFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkVerifyEmailFormat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		VerifyEmailFormat("9e9d7@1732205139ab.com.")
	}
}

func TestReg2VerifyEmailFormat(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"valid", args{"9e9d7@1732205139ab.com"}, true},
		{"invalid1", args{"9e9d7@1732205139ab.com."}, false},
		{"invalid2", args{".9e9d7@1732205139ab.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reg2VerifyEmailFormat(tt.args.email); got != tt.want {
				t.Errorf("Reg2VerifyEmailFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReg2VerifyEmailFormat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		VerifyEmailFormat("9e9d7@1732205139ab.com.")
	}
}
