package aes

import (
	"crypto/md5"
)

/*
Description  :
*/
/*
File name    : openssl_compat.go
Author       : miaoyc
Create time  : 2025/9/28
Update time  : 2025/9/28
Description  : 兼容CryptoJS的加解密函数和密钥生成算法
*/

// generateKeyAndIV 与CryptoJS生成key和iv算法实现
func generateKeyAndIV(salt []byte, key string) (iv, calKey []byte) {
	hash1 := md5.Sum([]byte(key + string(salt)))
	hash2 := md5.Sum(append(hash1[:], []byte(key+string(salt))...))
	hash3 := md5.Sum(append(hash2[:], []byte(key+string(salt))...))
	calKey = append(hash1[:], hash2[:]...)
	iv = hash3[:]
	return
}

// CryptoJsCbcEncrypt 兼容CryptoJS的默认加密方法,  CryptoJS.AES.encrypt
func CryptoJsCbcEncrypt(plaintext []byte, password string) (finalCiphertext []byte, err error) {
	// 生成随机salt
	salt, saltErr := generateSalt()
	if saltErr != nil {
		err = saltErr
		return
	}
	// 生成key和iv
	iv, calKey := generateKeyAndIV(salt, password)
	// 加密
	ciphertext, encErr := CbcEncrypt(plaintext, calKey, iv, PKCS7Padding)
	if encErr != nil {
		err = encErr
		return
	}
	// 添加前缀
	finalCiphertext = append([]byte(prefix), salt...)
	finalCiphertext = append(finalCiphertext, ciphertext...)
	return
}

// CryptoJsCbcDecrypt 兼容CryptoJS的默认解密方法, CryptoJS.AES.decrypt
func CryptoJsCbcDecrypt(ciphertext []byte, password string) (plaintext []byte, err error) {
	// 获取salt
	salt := ciphertext[8:16]
	// 获取key和iv
	iv, calKey := generateKeyAndIV(salt, password)
	// 去除前缀与salt
	ciphertext = ciphertext[16:]
	// 解密
	plaintext, err = CbcDecrypt(ciphertext, calKey, iv, PKCS7UnPadding)
	if err != nil {
		return
	}
	return
}
