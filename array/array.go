package array


/*
File name    : array.go
Author       : miaoyc
Create date  : 2023/1/11 3:01 下午
Description  : 
*/


// removeDuplicateElement 元素去重
func removeDuplicateElement(values []string) []string {
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
