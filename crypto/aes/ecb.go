package aes

import (
	"crypto/aes"
)

/*
File name    : ecb.go
Author       : miaoyc
Create date  : 2022/8/31 18:19
Update date  : 2024/7/11 14:09
Description  : ecb加密
*/

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
