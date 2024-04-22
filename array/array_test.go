package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
File name    : array_test.go
Author       : miaoyc
Create Date  : 2024/4/23 01:36
Update Date  : 2024/4/23 01:36
Description  :
*/

func TestIn(t *testing.T) {
	numbers := []int{0, 1, 2, 3}
	assert.Equal(t, true, In(numbers, 1))
	assert.Equal(t, false, In(numbers, 10))
	strs := []string{"0", "1", "2", "3"}
	assert.Equal(t, true, In(strs, "1"))
	assert.Equal(t, false, In(strs, "10"))
}
