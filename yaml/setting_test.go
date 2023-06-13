package setting

import (
	"fmt"
	"testing"
)

func TestSetup(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"load config", args{"./test.yaml"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Setup(tt.args.configFile)
			fmt.Println(GlobalConf)
		})
	}
}
