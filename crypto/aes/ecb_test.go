package aes

import (
	"crypto/aes"
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
File name    : ecb_test.go
Author       : miaoyc
Create time  : 2024/7/11 12:51
Update time  : 2024/7/11 12:51
Description  :
*/

func TestEcbEncrypt(t *testing.T) {
	key := "miaoyc-key"
	source := "miaoyc"
	pkey := ZeroPadding([]byte(key), aes.BlockSize)
	s1, _ := EcbEncrypt([]byte(source), pkey)
	s2 := base64.StdEncoding.EncodeToString(s1)
	assert.Equal(t, "1jKSOChj2uqwgw0M/Oa2fw==", s2)
}

func TestEcbDecrypt(t *testing.T) {
	key := "miaoyc-key"
	pkey := ZeroPadding([]byte(key), aes.BlockSize)
	source := "1jKSOChj2uqwgw0M/Oa2fw=="
	s1, _ := base64.StdEncoding.DecodeString(source)
	s2, _ := EcbDecrypt(s1, pkey)
	assert.Equal(t, "miaoyc", string(s2))
}
