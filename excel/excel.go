package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	// 打开 Excel 文件
	f, err := excelize.OpenFile("~/1.xlsx")
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}

	// 获取所有工作表名称
	sheetList := f.GetSheetList()
	fmt.Println("Sheets:", sheetList)

	// 读取 "Sheet1" 中的所有单元格数据
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
