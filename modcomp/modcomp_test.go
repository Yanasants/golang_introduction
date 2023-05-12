package main

import (
	"testing"
)

var initialY0 float64 = 1.0
var initialY1 float64 = 0.0
var initialX float64 = 0.0
var finalX float64 = 1.0
var functionDomain float64 = 1000.0
var h float64 = (finalX - initialX) / functionDomain
var allX []float64

func TestEquationsSystem1Index0(t *testing.T) {
	// All x values
	for i := initialX; i <= finalX; i += h {
		allX = append(allX, i)
	}

	// Method
	y0 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[0]
	y1 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[1]

	y0i0 := y0[200]
	y1i0 := y1[200]

	if y0i0 < 1.63 || y0i0 > 1.64 {
		t.Errorf("Value out of expected range.")
	}
	if y1i0 < -0.82 || y1i0 > -0.81 {
		t.Errorf("Value out of expected range.")
	}

}
