package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomEncoder_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    map[string]interface{}
		want    string
		wantErr bool
	}{
		{
			name: "基本测试",
			data: map[string]interface{}{
				"name": "John",
				"age":  30,
			},
			want:    `{"name":"John","age":30}`,
			wantErr: false,
		},
		{
			name: "包含HTML字符",
			data: map[string]interface{}{
				"html": "<script>alert('XSS')</script>",
			},
			want:    `{"html":"<script>alert('XSS')</script>"}`,
			wantErr: false,
		},
		{
			name: "包含特殊字符",
			data: map[string]interface{}{
				"special": "& < > ' \"",
			},
			want:    `{"special":"& < > ' \""}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CustomEncoder{
				Data: tt.data,
			}
			got, err := c.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomEncoder.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.JSONEq(t, tt.want, string(got))
		})
	}
}
