package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
File name    : sha1_test.go
Author       : miaoyc
Create Date  : 2024/3/12 23:51
Update Date  : 2024/3/12 23:51
Description  :
*/

func TestSha1(t *testing.T) {
	s := Sha1("")
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", s)
	s = Sha1("miaoyc")
	assert.Equal(t, "820a3d7b99f7b460c3c24e70ddbd16fd1fb9eb5d", s)
}
