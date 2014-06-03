package processor

import (
	"fmt"
	"math"
)


type Processor struct {
	previousLength int // length of previous row if any
}


func (p *Processor) ProcessRow(row []string) {
	rowL := len(row)
	// ignore empty row
	if (rowL == 0) { return }

	if (smallJump(p.previousLength, rowL)) {
		continueCurrentCube(p, row)
	} else {
		breakCurrentCube(p, row)
	}


	// update processor length
	p.previousLength = rowL
}

// add row to existing cube
func continueCurrentCube(p* Processor, row []string) {

}

// start a new cube in the processor
// handle the previous cube
func breakCurrentCube(p *Processor, row []string) {

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