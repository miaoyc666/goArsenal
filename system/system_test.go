package system

import (
	"testing"
)

/*
File name    : system_test.go.go
Author       : miaoyc
Create date  : 2021/12/2 5:07 下午
Description  :
*/

func Test_RunSystemShell(t *testing.T) {
	type args struct {
		name  string
		args  string
		shell string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{"case1", args{"sh", "-c", "pwd | wc -l"}, "1\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := RunSystemShell(tt.args.name, tt.args.args, tt.args.shell)
			if (err != nil) != tt.wantErr {
				t.Errorf("runSystemShell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("runSystemShell() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}