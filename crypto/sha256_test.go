package crypto

import (
	"fmt"
	"testing"
)

func Test_Sha256Encrypt(t *testing.T) {
	inputString := "Hello, SHA-256!"
	encryptedString := Sha256Encrypt(inputString)
	fmt.Printf("原始字符串: %s\n", inputString)
	fmt.Printf("SHA-256加密后: %s\n", encryptedString)
}
