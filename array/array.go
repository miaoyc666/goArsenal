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
func MergeAndDeduplicate(a, b interface{}) interface{} {
	unique := make(map[interface{}]bool)

	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	for i := 0; i < va.Len(); i++ {
		unique[va.Index(i).Interface()] = true
	}

	for i := 0; i < vb.Len(); i++ {
		unique[vb.Index(i).Interface()] = true
	}

	result := reflect.MakeSlice(va.Type(), 0, len(unique))
	for key := range unique {
		result = reflect.Append(result, reflect.ValueOf(key))
	}

	return result.Interface()
}

// In 类python的in操作符
func In(slice interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(slice)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			if arrV.Index(i).Interface() == elem {
				return true
			}
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
