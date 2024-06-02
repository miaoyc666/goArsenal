package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	s := Md5("miaoyc")
	assert.Equal(t, "ece751dcb8e591181c83e121689cc6ba", s)
}

func TestMD5sumFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"clac file md5", args{"./md5.go"}, "34173e9e3d1bd8952b2f8beba2083d91", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MD5sumFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("MD5sumFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MD5sumFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
