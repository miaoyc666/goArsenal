package array

import "reflect"

/*
File name    : array.go
Author       : miaoyc
Create date  : 2023/1/11 3:01 下午
Description  :
*/

// RemoveDuplicateElement 元素去重
func RemoveDuplicateElement(values []string) []string {
	result := make([]string, 0, len(values))
	temp := map[string]int32{}
	for _, item := range values {
		if _, ok := temp[item]; !ok {
			temp[item] = 1
			result = append(result, item)
		}
	}
	return result
}

// todo: 合并去重

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
