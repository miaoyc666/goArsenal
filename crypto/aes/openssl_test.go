package aes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/miaoyc666/goArsenal/crypto"
)

/*
File name    : openssl_test.go.go
Create time  : 2024/7/12 11:17
Update time  : 2024/7/12 11:17
Description  :
*/

func TestOpensslCbcEncrypt(t *testing.T) {
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
