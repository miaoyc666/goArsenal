package archive

import "testing"

func TestZipCSVBytesWithPassword(t *testing.T) {
	// 示例 csv 字节流
	csvBytes := []byte("id,name,age\n1,Alice,22\n2,Bob,30\n")

	err := ZipBytesWithPassword(csvBytes, "output.zip", "info.csv", "123456")
	if err != nil {
		panic(err)
	}

	println("ZIP 加密压缩完成：output.zip")
}
