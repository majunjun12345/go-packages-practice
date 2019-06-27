package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	// ReadXlsx()

	// CreateXlsx()

	UpdateXlsx()
}

// read
func ReadXlsx() {
	excelFileName := "goTest.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	CheckErr(err)

	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}

// create
func CreateXlsx() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1, row2 *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	CheckErr(err)

	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "年龄"

	row1 = sheet.AddRow()
	row1.SetHeightCM(1)
	cell = row1.AddCell()
	cell.Value = "狗子"
	cell = row1.AddCell()
	cell.Value = "18"

	row2 = sheet.AddRow()
	row2.SetHeightCM(1)
	cell = row2.AddCell()
	cell.Value = "蛋子"
	cell = row2.AddCell()
	cell.Value = "28"

	err = file.Save("test_write.xlsx")
	CheckErr(err)
}

// update
func UpdateXlsx() {
	excelFileName := "test_write.xlsx"
	xlsxFile, err := xlsx.OpenFile(excelFileName)
	CheckErr(err)

	sheet := xlsxFile.Sheets[0]
	row := sheet.AddRow()
	row.SetHeightCM(1)

	cell := row.AddCell()
	cell.Value = "张三"
	cell = row.AddCell()
	cell.Value = "21"

	err = xlsxFile.Save("test_write.xlsx")
	CheckErr(err)

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
