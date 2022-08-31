package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
    "bytes"
    "encoding/base64"
	"errors"
	"fmt"
	"io"
)

// NewCipher creates and returns a new cipher.Block.
// The key argument should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
var (
	commonkey = []byte("miaoyc@@aes@@key")
	commonkey1 = "miaoyc@@aes@@key"
)

func AesCFBEncrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(commonkey)
	if err != nil {
		return "", err

	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    fmt.Println(ciphertext)
	iv := ciphertext[:aes.BlockSize]

    fmt.Println(iv)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err

	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:],
		[]byte(plaintext))
	return hex.EncodeToString(ciphertext), nil

}

func AesCFBDecrypt(d string) (string, error) {
	ciphertext, err := hex.DecodeString(d)
	if err != nil {
		return "", err

	}
	block, err := aes.NewCipher(commonkey)
	if err != nil {
		return "", err

	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")

	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil

}

// CBCEncrypt AES-CBC 加密
// key 必须是 16(AES-128)、24(AES-192) 或 32(AES-256) 字节的 AES 密钥；
// 初始化向量 iv 为随机的 16 位字符串 (必须是16位)，
// 解密需要用到这个相同的 iv，因此将它包含在密文的开头。
func AesCBCEncrypt(plaintext string, key string) string {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("cbc decrypt err:", err)
        }
    }()

    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return ""
    }

    blockSize := len(key)
    padding := blockSize - len(plaintext)%blockSize // 填充字节
    if padding == 0 {
        padding = blockSize
    }

    // 填充 padding 个 byte(padding) 到 plaintext
    plaintext += string(bytes.Repeat([]byte{byte(padding)}, padding))
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err = rand.Read(iv); err != nil { // 将同时写到 ciphertext 的开头
        return ""
    }

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(plaintext))

    return base64.StdEncoding.EncodeToString(ciphertext)
}

// CBCDecrypt AES-CBC 解密
func AesCBCDecrypt(ciphertext string, key string) string {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("cbc decrypt err:", err)
        }
    }()

    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return ""
    }

    ciphercode, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return ""
    }

    iv := ciphercode[:aes.BlockSize]        // 密文的前 16 个字节为 iv
    ciphercode = ciphercode[aes.BlockSize:] // 正式密文

    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(ciphercode, ciphercode)

    plaintext := string(ciphercode) // ↓ 减去 padding
    return plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]
}

func main() {
    source := "this is data"
    fmt.Println("source data: ", source)
    fmt.Println("### cfb mode")
	s, _ := AesCFBEncrypt(source)
    fmt.Println("Encrypt: ", s)
	a, _ := AesCFBDecrypt(s)
    fmt.Println("Decrypt", a)

    /**
    fmt.Println("### cbc mode")
	s1 := AesCBCEncrypt(source, commonkey1)
    fmt.Println("Encrypt: ", s1)
	a1 := AesCBCDecrypt(s1, commonkey1)
    fmt.Println("Decrypt: ", a1)
    **/
}
