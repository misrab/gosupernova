package cubeExtractor

import (
	"fmt"
	//"log"
	 "os"
	"io"
	"path/filepath"
	"encoding/csv"

	//"github.com/tealeg/xlsx"

	"github.com/Misrab/gosupernova/analytics/processor"
)

func ProcessFiles(paths []string) {
	// process in paralell
	N := len(paths)
	c := make(chan int, N)

	for _, path := range paths {
		go processFileType(path, c)
	}

	// block untill all done
	//results := []int
	for i := 0; i < N; i++ { <- c }
}



// determines how to handle file based on extension
func processFileType(path string, c chan int) {
	extension := filepath.Ext(path)
	switch extension {
	case ".xlsx":
		processExcelFile(path, c)
	case ".csv":
		processCsvFile(path, c)
	default:
		// unsupported
		// TODO error handling
		fmt.Println("File type unsupported")
	}
}


func processExcelFile(path string, c chan int) {
	c <- 1
}

func processCsvFile(path string, c chan int) {
	file, err := os.Open(path)
	defer file.Close()

    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    reader := csv.NewReader(file)
   	// reader.TrailingComma = true
    //reader.LazyQuotes = true
    reader.FieldsPerRecord = -1 // don't check
    reader.TrimLeadingSpace = true


    p := new(processor.Processor)

    // ! ReadAll avoided to save on memory
    for {
	    row, err := reader.Read()
	    if err == io.EOF {
	        break
	    } else if err != nil {
	        panic(err)
	    }
	    // process row
	    p.ProcessRow(row)
	}

	//fmt.Println(p.CurrentCube)
	/*
	for i := 0; i < len(p.CurrentCube.Data); i++ {
		fmt.Println(p.CurrentCube.Data[i])
	}*/


	c <- 1
}