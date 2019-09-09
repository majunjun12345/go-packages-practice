package main

import (
	"fmt"

	"github.com/Unknwon/com"
	"github.com/tealeg/xlsx"
)

type Student struct {
	Name   string
	Gender string
	Age    int
	School string
}

func main() {
	// Write()
	Read()
}

func Read() {
	f, err := xlsx.OpenFile("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, sheet := range f.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Println(cell.String())
			}
		}
	}
}

func Write() {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("test_sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	titles := []string{"name", "gender", "age", "school"}
	row := sheet.AddRow()
	for _, title := range titles {
		cell := row.AddCell()
		cell.Value = title
	}

	students := GetStudents()
	for _, student := range students {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = student.Name

		cell = row.AddCell()
		cell.Value = student.Gender

		cell = row.AddCell()
		cell.Value = com.ToStr(student.Age)

		cell = row.AddCell()
		cell.Value = student.School
	}

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func GetStudents() []Student {
	students := []Student{
		Student{
			Name:   "zhangsan",
			Gender: "male",
			Age:    21,
			School: "s1",
		},
		Student{
			Name:   "lisi",
			Gender: "female",
			Age:    19,
			School: "s2",
		},
		Student{
			Name:   "wangwu",
			Gender: "female",
			Age:    18,
			School: "s3",
		},
	}
	return students
}
