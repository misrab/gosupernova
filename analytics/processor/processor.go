package processor

import (
	//"fmt"
	"math"
	"strconv"
	"strings"
)

// CUBE START

type Cube struct {
	NumRows int  // depending on how we store data, could omit this = len(Data)
	Labels []string
	Types []string
	Info []string
	// Datapath
	Data [][]string
}


// CUBE END
type Processor struct {
	PreviousLength int // length of Previous row if any
	CurrentCube *Cube // pointer to current cube
	PotentialCubes []*Cube // array of pointers to cubes
}


func (p *Processor) ProcessRow(row []string) {
	rowL := len(row)
	// ignore empty row
	if (rowL == 0) { return }

	if (smallJump(p.PreviousLength, rowL)) {
		continueCurrentCube(p, row)
	} else {
		breakCurrentCube(p, row)
	}


	// update processor length
	p.PreviousLength = rowL
}

// add row to existing cube
func continueCurrentCube(p* Processor, row []string) {
	//fmt.Println("continue")
	p.CurrentCube.Data = append(p.CurrentCube.Data, row)
	p.CurrentCube.NumRows++
}

// start a new cube in the processor
// handle the Previous cube
func breakCurrentCube(p *Processor, row []string) {
	// store previous cube if any
	
	var prevLabels []string
	if p.CurrentCube != nil {
		p.PotentialCubes = append(p.PotentialCubes, p.CurrentCube)
		prevLabels = p.CurrentCube.Labels 
	}

	// set new cube
	p.CurrentCube = new(Cube)
	if potentialLabels(row) {
		p.CurrentCube.Labels = row
		return
	}

	p.CurrentCube.Labels = prevLabels
	p.CurrentCube.Data = append(p.CurrentCube.Data, row)
	p.CurrentCube.NumRows++
}

func potentialLabels(row []string) bool {
	for _, cell := range row {
		if InferType(cell) != "string" { return false }
	}

	return true
}

// try to infer type of a cell
func InferType(cell string) string {
	// trim first
	cell = strings.Trim(cell, " ")

	// empty string
	if cell == "" {
		return "empty"
	}

	// float
	_, err := strconv.ParseFloat(cell, 64)
	if err == nil {
		return "float64"
	}

	// datetime


	return "string"
}


// what we define as a small jump in length
func smallJump(oldLength int, newLength int) bool {
	THRESHOLD := 0.4

	oldL := float64(oldLength)
	newL := float64(newLength)
	//fmt.Printf("Old: %f, New: %f", oldL, newL)
	percentDifference := math.Abs(oldL - newL) / math.Max(oldL, newL)

	return percentDifference < THRESHOLD
}