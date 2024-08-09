package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func sha256Encrypt(inputString string) string {
	// 创建一个新的sha256哈希对象
	hasher := sha256.New()
	// 写入要计算的数据
	hasher.Write([]byte(inputString))
	// 计算哈希值，返回一个字节切片
	hashBytes := hasher.Sum(nil)
	// 将字节切片转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
