package array

import (
	"reflect"
	"encoding/json"
)

/*
File name    : array.go
Author       : miaoyc
Create date  : 2023/1/11 15:01
Update date  : 2025/1/20 22:12
Description  : array自定义函数，使用反射的函数有一定性能损耗（非核心高性能场景可以使用）
*/


// StructToJSON 是一个通用函数，可以将任何struct转换成JSON字符串
// 如果转换成功，返回JSON字符串和nil
// 如果转换失败，返回空字符串和错误信息
func StructToJSON(v interface{}) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

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

// UnionStringSlices 合并多个字符串切片并去重
// 参数：slices - 二维字符串切片
// 返回值：合并并去重后的字符串切片
// 返回结果无排序
func UnionStringSlices(slices [][]string) []string {
	set := make(map[string]struct{})
	for _, slice := range slices {
		for _, item := range slice {
			set[item] = struct{}{}
		}
	}

	union := make([]string, 0, len(set))
	for item := range set {
		union = append(union, item)
	}

	return union
}

// MergeStringSlices 合并多个字符串切片并取交集
// 参数：slices - 二维字符串切片
// 返回值：合并并取交集后的字符串切片
func MergeStringSlices(slices [][]string) []string {
	if len(slices) == 0 {
		return []string{}
	}

	// 初始化交集集合为第一个切片的元素
	intersection := make(map[string]struct{})
	for _, item := range slices[0] {
		intersection[item] = struct{}{}
	}

	// 逐个与后续切片取交集
	for _, slice := range slices[1:] {
		tempSet := make(map[string]struct{})
		for _, item := range slice {
			if _, exists := intersection[item]; exists {
				tempSet[item] = struct{}{}
			}
		}
		intersection = tempSet
	}

	// 将交集集合转换为切片
	result := make([]string, 0, len(intersection))
	for item := range intersection {
		result = append(result, item)
	}

	return result
}
