package qps

import (
	"testing"
	"unsafe"
)

/*
File name    : qps_test.go
Author       : miaoyc1989@hotmail.com
Create date  : 2020/5/9 3:47 下午
Description  :
*/

func TestGetData(t *testing.T) {
	type args struct {
		fileName string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetTestFunc(t *testing.T) {
	var tests []struct {
		name string
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTestFunc(); got != tt.want {
				t.Errorf("GetTestFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimer(t *testing.T) {
	type args struct {
		qpsLimit   int
		totalCount int
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		data unsafe.Pointer
	}
	var tests []struct {
		name string
		args args
		want int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unmarshal(tt.args.data); got != tt.want {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
