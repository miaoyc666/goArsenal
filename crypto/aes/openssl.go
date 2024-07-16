package aes

import (
	"crypto/rand"
	"crypto/sha256"
	"hash"
	"io"
)

/*
File name    : openssl.go
Author       : miaoyc
Create time  : 2024/7/11 18:58
Update time  : 2024/7/16 12:23
Description  : 兼容openssl命令的加解密函数
*/

const (
	saltSize = 8
	keyLen   = 32 // 256 bits
	ivLen    = 16 // 128 bits
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

func OpensslCbcEncrypt(plaintext []byte, password string) (finalCiphertext []byte, err error) {
	// 生成随机salt
	salt, saltErr := generateSalt()
	if saltErr != nil {
		err = saltErr
		return
	}
	// 生成key和iv
	calKey, iv, genErr := EVPBytesToKey([]byte(password), salt, keyLen, ivLen, sha256.New)
	if genErr != nil {
		err = genErr
		return
	}
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

// OpensslCbcDecrypt 兼容openssl命令的解密方法
// 相当于执行命令： echo -n "ciphertext" | openssl enc -aes-256-cbc -k xxxx -salt
func OpensslCbcDecrypt(ciphertext []byte, password string) (plaintext []byte, err error) {
	// 获取salt
	salt := ciphertext[8:16]
	// 生成key和iv
	calKey, iv, genErr := EVPBytesToKey([]byte(password), salt, keyLen, ivLen, sha256.New)
	if genErr != nil {
		err = genErr
		return
	}
	// 去除前缀与salt
	ciphertext = ciphertext[16:]
	// 解密
	plaintext, err = CbcDecrypt(ciphertext, calKey, iv, PKCS7UnPadding)
	if err != nil {
		return
	}
	return
}

// EVPBytesToKey 兼容openssl的EVP_BytesToKey方法，用于生成加密密钥和IV
// hash算法说明：openssl默认的hash函数，从OpenSSL 1.1开始由 MD5 变为 SHA256
func EVPBytesToKey(password, salt []byte, keyLen, ivLen int, h func() hash.Hash) ([]byte, []byte, error) {
	var key, iv []byte
	var data []byte

	// Calculate the number of hash iterations needed
	dLen := keyLen + ivLen
	for len(data) < dLen {
		h := h()
		if len(data) > 0 {
			h.Write(data)
		}
		h.Write(password)
		h.Write(salt)
		data = h.Sum(data)
	}

	key = data[:keyLen]
	iv = data[keyLen : keyLen+ivLen]
	return key, iv, nil
}
