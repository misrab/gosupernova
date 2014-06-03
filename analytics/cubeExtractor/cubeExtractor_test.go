package cubeExtractor

import (
	"testing"
)


func TestProcessFiles(t *testing.T) {
	// Two xlsx and one csv file for testing
	/*
	paths := []string{"/Users/Misrab/go/src/github.com/Misrab/gosupernova/testing/Fleet performance.xlsx", 
		"/Users/Misrab/go/src/github.com/Misrab/gosupernova/testing/Shipments performed in 2012.xlsx", 
		"/Users/Misrab/go/src/github.com/Misrab/gosupernova/testing/us-500.csv"}
	*/
	paths := []string{"/Users/Misrab/go/src/github.com/Misrab/gosupernova/testing/us-500.csv"}

	ProcessFiles(paths)
}