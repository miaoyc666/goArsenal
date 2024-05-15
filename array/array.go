package array

import (
	"reflect"
)

/*
File name    : array.go
Author       : miaoyc
Create date  : 2023/1/11 3:01 下午
Update date  : 2023/1/11 3:01 下午
Description  : array自定义函数，使用反射的函数有一定性能损耗（非核心高性能场景可以使用）
*/

// RemoveDuplicateElement 元素去重
func RemoveDuplicateElement(input interface{}) interface{} {
	inputVal := reflect.ValueOf(input)
	if inputVal.Kind() != reflect.Slice {
		panic("RemoveDuplicateElement: input is not a slice")
	}

	resultVal := reflect.MakeSlice(inputVal.Type(), 0, inputVal.Len())
	temp := map[interface{}]struct{}{}
	for i := 0; i < inputVal.Len(); i++ {
		item := inputVal.Index(i).Interface()
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			resultVal = reflect.Append(resultVal, inputVal.Index(i))
		}
	}

	return resultVal.Interface()
}

// MergeAndDeduplicate 合并两个数组并去除重复项
func MergeAndDeduplicate[T comparable](a, b []T) []T {
	unique := make(map[T]bool)
	for _, item := range a {
		unique[item] = true
	}
	for _, item := range b {
		unique[item] = true
	}
	result := make([]T, 0, len(unique))
	for key := range unique {
		result = append(result, key)
	}
	return result
}

// In 类python的in操作符
func In[T comparable](slice []T, elem T) bool {
	for _, item := range slice {
		if item == elem {
			return true
		}
	}
	return false
}

// SliceEqual 判断两个切片是否相等
func SliceEqual(a, b interface{}) bool {
	countElements := func(slice interface{}) map[interface{}]int {
		count := make(map[interface{}]int)
		switch slice := slice.(type) {
		case []int:
			for _, v := range slice {
				count[v]++
			}
		case []string:
			for _, v := range slice {
				count[v]++
			}
			// 添加其他类型的判断，如 []float64 等
		}
		return count
	}

	mapEqual := func(a, b map[interface{}]int) bool {
		return reflect.DeepEqual(a, b)
	}

	switch a := a.(type) {
	case []int:
		if b, ok := b.([]int); ok {
			return mapEqual(countElements(a), countElements(b))
		}
	case []string:
		if b, ok := b.([]string); ok {
			return mapEqual(countElements(a), countElements(b))
		}
		// 添加其他类型的判断，如 []float64 等
	}
	return false
}
