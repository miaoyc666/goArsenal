package json

import "testing"

/*
File name    : formatPrintJson_test.go
Author       : miaoyc
Create Date  : 2023/5/1 17:37
Update Date  : 2023/5/1 17:37
Description  :
*/

func TestFormatPrintJson(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"format-1", args{`{"name":"miaoyc","age":18}`}, "{\n\t\"name\": \"miaoyc\",\n\t\"age\": 18\n}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatPrintJson(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatPrintJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatPrintJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
