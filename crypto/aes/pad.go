package aes

import (
	"bytes"
)

/*
File name    : pad.go
Author       : miaoyc
Create time  : 2024/7/11 14:06
Update time  : 2024/7/11 14:06
Description  : 填充算法
*/

// ZeroPadding 对给定的数据进行Zero Padding填充
// data: 需要填充的原始数据
// blockSize: 块的大小，AES的块大小通常是16字节
func ZeroPadding(data []byte, blockSize int) []byte {
	paddingSize := blockSize - len(data)%blockSize
	if paddingSize == blockSize {
		return data
	}
	padding := bytes.Repeat([]byte{0}, paddingSize)
	return append(data, padding...)
}

// PKCS7Padding PKCS7进行填充
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS7UnPadding PKCS7去填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
