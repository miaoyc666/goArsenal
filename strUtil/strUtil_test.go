package strUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
File name    : strUtil_test.go
Create time  : 2024/5/20 01:20
Update time  : 2024/5/20 01:20
Description  :
*/

func TestIsUUID(t *testing.T) {
	s := IsUUID("e397692f-1f01-4d22-bba8-60f5bca607aa")
	assert.Equal(t, s, true)
	s = IsUUID("e397692f-1f01-4d22")
	assert.Equal(t, s, false)
}

func TestIsMD5(t *testing.T) {
	s := IsMD5("35a075221d7cbbcc43842cdbbbd767f8")
	assert.Equal(t, s, true)
	s = IsMD5("35a075221d7cbbcc43")
	assert.Equal(t, s, false)
	s = IsMD5("35a075221d7cbbcc43842cdbbbd767f8x")
	assert.Equal(t, s, false)
}

func TestConcat(t *testing.T) {
	str1, str2 := "xxx", "yyy"
	assert.Equal(t, "xxxyyy", Concat(str1, str2))
}
