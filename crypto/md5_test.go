package crypto

import "testing"

func TestMd5(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name       string
		args       args
		wantMd5str string
	}{
		// TODO: Add test cases.
		{"md5", args{"miaoyc"}, "ece751dcb8e591181c83e121689cc6ba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMd5str := Md5(tt.args.data); gotMd5str != tt.wantMd5str {
				t.Errorf("Md5() = %v, want %v", gotMd5str, tt.wantMd5str)
			}
		})
	}
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
