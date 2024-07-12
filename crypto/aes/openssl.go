package aes

import (
	"crypto/md5"
	"crypto/rand"
	"io"
)

/*
File name    : openssl.go
Author       : miaoyc
Create time  : 2024/7/11 18:58
Update time  : 2024/7/11 18:58
Description  : 兼容openssl命令的加解密函数
*/

const (
	saltSize = 8
	prefix   = "Salted__"
)

// generateSalt 生成随机盐
func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// generateKeyAndIV 与openssl生成key和iv算法相同的实现
func generateKeyAndIV(salt []byte, key string) (iv, calKey []byte) {
	hash1 := md5.Sum([]byte(key + string(salt)))
	hash2 := md5.Sum(append(hash1[:], []byte(key+string(salt))...))
	hash3 := md5.Sum(append(hash2[:], []byte(key+string(salt))...))
	calKey = append(hash1[:], hash2[:]...)
	iv = hash3[:]
	return
}

// getKeyAndIv 与openssl获取key和iv算法相同的实现
func getKeyAndIv(ciphertext []byte, key string) (iv []byte, calKey []byte) {
	salt := ciphertext[8:16]
	hash1 := md5.Sum([]byte(key + string(salt)))
	hash2 := md5.Sum(append(hash1[:], []byte(key+string(salt))...))
	hash3 := md5.Sum(append(hash2[:], []byte(key+string(salt))...))
	calKey = append(hash1[:], hash2[:]...)
	iv = hash3[:]
	return
}

func OpensslCbcEncrypt(plaintext []byte, password string) (finalCiphertext []byte, err error) {
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

func OpensslCbcDecrypt(ciphertext []byte, password string) (plaintext []byte, err error) {
	// 获取key和iv
	iv, calKey := getKeyAndIv(ciphertext, password)
	// 去除前缀与salt
	ciphertext = ciphertext[16:]
	// 解密
	plaintext, err = CbcDecrypt(ciphertext, calKey, iv, PKCS7UnPadding)
	if err != nil {
		return
	}
	return
}
