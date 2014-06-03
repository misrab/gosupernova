package cubeExtractor

import (
	"testing"
	//"fmt"
	//"github.com/Misrab/gosupernova/analytics/cubeExtractor"
)


func TestReadFileStream(t *testing.T) {
	excel_path := "/Users/Misrab/go/src/github.com/Misrab/gosupernova/testing/Fleet performance.xlsx"
	go ReadFileStream(excel_path)
}