package main

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"strings"
    "encoding/base64"
)

/*
ECB加密模式的简单示例
*/

// keyPad 加密key填充算法, 填充方式，加密内容必须为16字节的倍数，若不足则使用0进行填充
func keyPad(text string) string {
	count := 16 - len(text)%16
	pad := strings.Repeat("0", count)
	return text + pad
}

// PKCS7Padding 使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding PKCS7去填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// EcbDecrypt ecb解密
func EcbDecrypt(data, key []byte) string {
    decodeData, _ := base64.StdEncoding.DecodeString(string(data))
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(decodeData))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(decodeData); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return string(PKCS7UnPadding(decrypted))
}

// EcbEncrypt ecb加密
func EcbEncrypt(data, key []byte) string {
	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}
    return base64.StdEncoding.EncodeToString(decrypted)
}

func main() {
	key := "miaoyc-key"
	r := keyPad(key)
	fmt.Println(r)

	source := "this is data"
	fmt.Println("source data: ", source)
	fmt.Println("### ecb mode")
	s := EcbEncrypt([]byte(source), []byte(r))
	fmt.Println("Encrypt: ", s)
	a := EcbDecrypt([]byte(s), []byte(r))
	fmt.Println("Decrypt", string(a))
}
