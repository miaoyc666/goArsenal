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
