package crypto

import (
	"fmt"
	"testing"
)

func Test_sha256Encrypt(t *testing.T) {
	inputString := "Hello, SHA-256!"
	encryptedString := sha256Encrypt(inputString)
	fmt.Printf("原始字符串: %s\n", inputString)
	fmt.Printf("SHA-256加密后: %s\n", encryptedString)
}
