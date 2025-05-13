package strUtil

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_isValidName(t *testing.T) {
	s := IsValidName("hello-world")
	assert.Equal(t, s, true)
	s = IsValidName("hello-world_")
	assert.Equal(t, s, true)
	s = IsValidName("hello123")
	assert.Equal(t, s, true)
	s = IsValidName("hello)123")
	assert.Equal(t, s, false)
	s = IsValidName("hello 123")
	assert.Equal(t, s, false)
}

func TestIsMobile(t *testing.T) {
	s := IsMobile("13800138000")
	assert.Equal(t, s, true)
	s = IsMobile("23800138000")
	assert.Equal(t, s, false)
	s = IsMobile("1380013800")
	assert.Equal(t, s, false)
	s = IsMobile("1380013800a")
	assert.Equal(t, s, false)
	s = IsMobile("")
	assert.Equal(t, s, false)
}

func TestIsEmail(t *testing.T) {
	s := IsEmail("example@example.com")
	assert.Equal(t, s, true)
	s = IsEmail("example.example.com")
	assert.Equal(t, s, false)
	s = IsEmail("example@.com")
	assert.Equal(t, s, false)
	s = IsEmail("@example.com")
	assert.Equal(t, s, false)
	s = IsEmail("")
	assert.Equal(t, s, false)
}

func TestIsAllInvisibleOrSpace(t *testing.T) {
	s := IsAllInvisibleOrSpace(" ")
	assert.Equal(t, s, true)
	s = IsAllInvisibleOrSpace("")
	assert.Equal(t, s, true)
	s = IsAllInvisibleOrSpace("a")
	assert.Equal(t, s, false)
	s = IsAllInvisibleOrSpace("a ")
	assert.Equal(t, s, false)
}
