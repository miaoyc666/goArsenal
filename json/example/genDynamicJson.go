package main

import (
	"encoding/json"
	"fmt"
)

/*
File name    : genDynamicJson.go
Author       : miaoyc
Create date  : 2022/2/7 12:03 下午
Update date  : 2024/7/4 20:03 下午
Description  : 生成动态json数据
*/

func main() {
	data := make(map[string]interface{})
	data["name"] = "Eve"
	data["age"] = 22
	data["skills"] = []string{"Go", "Python", "JavaScript"}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
