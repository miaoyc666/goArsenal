package array

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
File name    : array_test.go
Author       : miaoyc
Create Date  : 2024/4/23 01:36
Update Date  : 2025/4/3 19:44
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

func TestRemoveDuplicateElement(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 1, 2}
	assert.Equal(t, []int{0, 1, 2, 3}, RemoveDuplicateElement(numbers))
	strs := []string{"0", "1", "2", "3", "1", "2"}
	assert.Equal(t, []string{"0", "1", "2", "3"}, RemoveDuplicateElement(strs))
}

func TestMergeAndDeduplicate(t *testing.T) {
	numbers := []int{0, 1, 2, 3, 1, 2}
	numbers1 := []int{0, 1, 2, 3, 1, 2}
	x := MergeAndDeduplicate(numbers, numbers1)
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))
	assert.Equal(t, true, SliceEqual([]int{0, 1, 2, 3}, MergeAndDeduplicate(numbers, numbers1)))
	strs := []string{"0", "1", "2", "3"}
	strs1 := []string{"3", "5", "4"}
	y := MergeAndDeduplicate(strs, strs1)
	fmt.Println(y)
	fmt.Println(reflect.TypeOf(y))
	fmt.Println([]string{"0", "1", "2", "3", "4", "5"})
	fmt.Println(reflect.TypeOf([]string{"0", "1", "2", "3", "4", "5"}))
	assert.Equal(t, true, SliceEqual([]string{"0", "1", "2", "3", "4", "5"}, MergeAndDeduplicate(strs, strs1)))
}

func TestSliceEqual(t *testing.T) {
	numbers := []int{0, 1, 2, 3}
	numbers1 := []int{0, 1, 2, 3}
	numbers2 := []int{0, 2, 3}
	assert.Equal(t, true, SliceEqual(numbers, numbers1))
	assert.Equal(t, false, SliceEqual(numbers, numbers2))
	strs := []string{"0", "1", "2", "3"}
	strs1 := []string{"0", "1", "2", "3"}
	strs2 := []string{"0", "1", "2"}
	assert.Equal(t, true, SliceEqual(strs, strs1))
	assert.Equal(t, false, SliceEqual(strs, strs2))
}

func TestPaginateItemsInt(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := PaginateItems(items, 3, 2)
	expected := []int{5, 6}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("PaginateItems() = %v, want %v", got, expected)
	}

}

func TestPaginateItemsString(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}

	got := PaginateItems(items, 2, 2)
	expected := []string{"c", "d"}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("PaginateItems() = %v, want %v", got, expected)
	}
}
