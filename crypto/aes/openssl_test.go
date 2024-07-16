package aes

import (
	"fmt"
	"testing"

	"github.com/miaoyc666/goArsenal/crypto"
	"github.com/stretchr/testify/assert"
)

/*
File name    : openssl_test.go
Create time  : 2024/7/12 11:17
Update time  : 2024/7/12 11:17
Description  :
*/

func TestOpensslCbcEncrypt(t *testing.T) {
	// 对于加密命令：echo -n "test" | openssl enc -aes-256-cbc -k xxxx -e -base64 -salt -p
	// 对应解密命令：echo "U2FsdGVkX18pHpnXSGlH1+yiVZUX80NDDfWbZMkhk9s=" | openssl enc -aes-256-cbc -k xxxx -d -base64 -salt -p
	fmt.Println("Encode")
	key := "xxxx"
	got, _ := OpensslCbcEncrypt([]byte("test"), key)
	got1 := crypto.Base64Encode(got)
	fmt.Println("Decode")
	got2, _ := crypto.Base64Decode(got1)
	got3, _ := OpensslCbcDecrypt(got2, key)
	fmt.Println(string(got3))
	assert.Equal(t, "test", string(got3))
}
