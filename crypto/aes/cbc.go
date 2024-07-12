package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

/*
File name    : cbc.go
Create time  : 2024/7/11 12:47
Update time  : 2024/7/11 12:47
Author       : miaoyc
Description  :
*/

// CbcEncrypt cbc加密，padFunc表示填充函数，可传入PKCS7Padding、ZeroPadding或其他符合格式的填充函数
func CbcEncrypt(plaintext, key, iv []byte, padFunc PadFunc) (ciphertext []byte, err error) {
	block, chipErr := aes.NewCipher(key)
	if chipErr != nil {
		err = chipErr
		return
	}
	plaintextBytes := padFunc(plaintext, aes.BlockSize)
	ciphertext = make([]byte, len(plaintextBytes))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintextBytes)
	return
}

// CbcDecrypt cbc解密，unPadFunc表示去填充函数，可传入PKCS7UnPadding、ZeroUnPadding或其他符合格式的去填充函数
func CbcDecrypt(ciphertext, key, iv []byte, unPadFunc UnPadFunc) (plaintext []byte, err error) {
	block, chipErr := aes.NewCipher(key)
	if chipErr != nil {
		err = chipErr
		return
	}
	plaintextBytes := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintextBytes, ciphertext)
	plaintext = unPadFunc(plaintextBytes)
	return
}
