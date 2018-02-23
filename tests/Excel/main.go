package main

import (
	"fmt"
    "github.com/tealeg/xlsx"
)

func ReadXLS(){
	excelFileName := "./Data.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
		fmt.Println(err.Error())
		return
	}
	
    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
                text := cell.String()
                fmt.Printf("%s\n", text)
            }
        }
    }
}

func WriteXLS(){
	var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    var err error

    file = xlsx.NewFile()
    sheet, err = file.AddSheet("Sheet1")
    if err != nil {
        fmt.Printf(err.Error())
    }
    row = sheet.AddRow()
    cell = row.AddCell()
    cell.Value = "I am a cell!"
    err = file.Save("MyXLSXFile.xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
}

func ReadAndWrite(){
    var cell *xlsx.Cell
    var err error

	excelFileName := "./Data.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
        fmt.Printf(err.Error())
	}

	for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
			cell = row.AddCell()
			//cell.Value = "I am a cell!"
			cell.SetString("AA2")
        }
	}
	
	err = xlFile.Save(excelFileName);
    if err != nil {
        fmt.Printf(err.Error())
    }
}

func main() {
	ReadAndWrite()
}