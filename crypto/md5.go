package crypto

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

/*
File name    : md5.go
Author       : miaoyc
Create date  : 2021/12/2 4:34 下午
Description  : 加密相关
*/

var (
	readFileBufferSize = 65536
)

// Md5 md5加密，input: 需加密的字符串, return：加密之后的32个字符的数据
func Md5(data string) (md5str string) {
	h := md5.New()
	h.Write([]byte(data))
	md5str = fmt.Sprintf("%x", h.Sum(nil))
	return md5str
}

// MD5sumFromFile 计算文件的md5
func MD5sumFromFile(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	for buf, reader := make([]byte, readFileBufferSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		hash.Write(buf[:n])
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}
