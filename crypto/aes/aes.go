package aes

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"strings"
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
func EcbDecrypt(data, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}
	return PKCS7UnPadding(decrypted), nil
}

// EcbEncrypt ecb加密
func EcbEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}
	return decrypted, nil
}

func B64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func B64Decode(src string) ([]byte, error) {
	s, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return s, err
	} else {
		return s, nil
	}
}

func main() {
	key := "miaoyc-key"
	r := keyPad(key)
	fmt.Println(r)

	//source := "miaoyc"
	//fmt.Println("source data: ", source)
	//fmt.Println("### ecb mode")
	//s1, _ := EcbEncrypt([]byte(source), []byte(r))
	//fmt.Println("Encrypt: ", string(s1))
	//b2, _ := EcbDecrypt(s1, []byte(r))
	//fmt.Println("Decrypt: ", string(b2))

	// base64
	a := "ymD7cJgkR8DP6dwxb4Ztag=="
	data, err := B64Decode(a)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	b3, _ := EcbDecrypt(data, []byte(r))
	fmt.Println("Decrypt: ", string(b3))
}
