package main

import (
	"fmt"

	"github.com/xuri/excelize"
)

func main() {
	// Read()

	Create()
}

// read
func Read() {
	xlsxFile, err := excelize.OpenFile("test_write.xlsx")
	Checkerr(err)

	cell, err := xlsxFile.GetCellValue("sheet1", "B1")
	Checkerr(err)
	fmt.Println(cell)

	rows, err := xlsxFile.GetRows("sheet1")
	Checkerr(err)

	for _, row := range rows {
		for _, cell := range row {
			fmt.Println(cell)
		}
	}
}

// create
func Create() {
	xlsxFile := excelize.NewFile()
	index := xlsxFile.NewSheet("sheet1")

	xlsxFile.SetCellValue("sheet1", "A1", "姓名")
	xlsxFile.SetCellValue("sheet1", "B1", "年龄")
	xlsxFile.SetCellValue("sheet1", "A2", "zhangsan")
	xlsxFile.SetCellValue("sheet1", "B2", 21)

	xlsxFile.SetActiveSheet(index)

	err := xlsxFile.SaveAs("test_excel.xlsx")
	Checkerr(err)

}

func Checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
