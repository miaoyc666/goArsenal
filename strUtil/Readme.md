strUtil
=

```go
// 自动检测字符串编码
package main

import (
    "fmt"
    "github.com/saintfish/chardet"
)

func main() {
    str := []byte("你好, 世界")
    detector := chardet.NewTextDetector()
    result, err := detector.DetectBest(str)
    if err != nil {
        fmt.Println("无法检测字符串编码:", err)
    } else {
        fmt.Println("检测到的编码:", result.Charset)
    }
}
```

```go

// DetectEncoding 自动检测字符串的字符编码
func DetectEncoding(input []byte) (string, error) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(input)
	if err != nil {
		return "", err
	}
	return result.Charset, nil
}

// ConvertToUTF8 将给定的字符串从指定的源编码转换为UTF-8
func ConvertToUTF8(src, srcEncoding string) (string, error) {
	// 创建一个转换器，将源编码转换为UTF-8
	encoding, _ := charset.Lookup(srcEncoding)
	if encoding == nil {
		return "", errors.New("未知的源编码")
	}
	// 使用转换器
	reader := transform.NewReader(strings.NewReader(src), encoding.NewDecoder())
	// 读取转换后的字符串
	utf8Str, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(utf8Str), nil
}

// getEncoding 根据编码名称返回对应的编码器和解码器
func getEncoding(name string) (*encoding.Decoder, *encoding.Encoder, error) {
	switch name {
	case "ISO-8859-1":
		return charmap.ISO8859_1.NewDecoder(), charmap.ISO8859_1.NewEncoder(), nil
	case "UTF-8":
		return encoding.Nop.NewDecoder(), encoding.Nop.NewEncoder(), nil
	case "UTF-16":
		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder(), unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewEncoder(), nil
	default:
		return nil, nil, fmt.Errorf("不支持的编码: %s", name)
	}
}

// ConvertEncoding 将输入字符串从 srcEncoding 转换为 dstEncoding
func ConvertEncoding(input string, srcEncoding, dstEncoding string) (string, error) {
	srcDecoder, _, err := getEncoding(srcEncoding)
	if err != nil {
		return "", err
	}
	_, dstEncoder, err := getEncoding(dstEncoding)
	if err != nil {
		return "", err
	}

	// 解码输入字符串
	decoded, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(input), srcDecoder))
	if err != nil {
		return "", fmt.Errorf("解码失败: %v", err)
	}

	// 编码为目标编码
	encoded, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(string(decoded)), dstEncoder))
	if err != nil {
		return "", fmt.Errorf("编码失败: %v", err)
	}

	return string(encoded), nil
}
```
