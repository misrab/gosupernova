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
	c := make(chan []*processor.Cube, N)

	for _, path := range paths {
		go processFileType(path, c)
	}

	// block untill all done
	var results []*processor.Cube
	for i := 0; i < N; i++ { 
		results = append(results, <- c...)
	}

	// print results
	for _, cubePointer := range results {
		fmt.Println("========= Potential Cube =========")
		fmt.Printf("ROWS: %d \n", cubePointer.NumRows)
		fmt.Printf("LABELS: %s \n", cubePointer.Labels)
	}
}



// determines how to handle file based on extension
func processFileType(path string, c chan []*processor.Cube) {
	extension := filepath.Ext(path)
	switch extension {
	//case ".xlsx":
	//	processExcelFile(path, c)
	case ".csv":
		processCsvFile(path, c)
	default:
		// unsupported
		// TODO error handling
		fmt.Println("File type unsupported")
	}
}


//func processExcelFile(path string, c chan int) {
//	c <- nil
//}

func processCsvFile(path string, c chan []*processor.Cube) {
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

	// add last cube to potentials
	p.PotentialCubes = append(p.PotentialCubes, p.CurrentCube)

	//fmt.Println(p.CurrentCube.Labels)
	/*
	for i := 0; i < len(p.CurrentCube.Data); i++ {
		fmt.Println(p.CurrentCube.Data[i])
	}*/

	// pass cubes found to channel
	c <- p.PotentialCubes
}