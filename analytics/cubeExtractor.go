package cubeExtractor

import (
	"fmt"
	"path/filepath"
	"github.com/tealeg/xlsx"
)


func ReadFileStream(path string) {
	extension := filepath.Ext(path)
	fmt.Println("Running ReadFileStream from path: " + path + " with extension: " + extension)

	switch extension {
	case ".xls":
		readExcelFile(path)
	case ".xlsx":
		readExcelFile(path)
	case ".csv":
		readCsvFile(path)
	default:
		panic("Unsupported file type")
	}
}



func readExcelFile(path string) {
	fmt.Println("Reading excel file...")

	xlFile, error := xlsx.OpenFile(path)
	if error != nil {
    	panic(error)
    }
    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
        	//row_slice := make()
        	fmt.Printf("%d\n", len(row.Cells))
        	/*
            for _, cell := range row.Cells {
                fmt.Printf("%s\n", cell.String())
            }
            */
        }
    }

}
func readCsvFile(path string) {
	fmt.Println("Reading csv file...")
}
