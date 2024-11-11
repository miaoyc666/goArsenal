package json

import (
	"bytes"
	"encoding/json"
)

type CustomEncoder struct {
	Data map[string]interface{}
}

// MarshalJSON 自定义JSON编码器, 禁用HTML转义
func (c CustomEncoder) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(c.Data)
	return buffer.Bytes(), err
}
