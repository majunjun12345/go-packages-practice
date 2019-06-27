package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

/*
	如需要在 csv 中使用引号，那么将字符串以双引号包围：""first_name is "zahngshan" haha"
	每个格以 , 分隔
*/

func main() {
	// Create()

	Read()
}

// create
func Create() {
	in1 := []string{"first_name", "last_name", "username"}
	in2 := []string{"Ken", "Thompson", "ken"}
	in3 := []string{"Rob", "Pike", "rob"}
	in4 := []string{"Robert", "Griesemer", "gri"}

	fi, err := os.Create("test.csv")
	CheckErr(err)
	csvWriter := csv.NewWriter(fi)

	csvWriter.Write(in1)
	csvWriter.Write(in2)
	csvWriter.Write(in3)
	csvWriter.Write(in4)
	csvWriter.Flush()
}

// read
func Read() {
	fi, err := os.Open("test.csv")
	CheckErr(err)
	csvReader := csv.NewReader(fi)

	records, err := csvReader.ReadAll()
	CheckErr(err)

	for _, record := range records {
		fmt.Println(record)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
