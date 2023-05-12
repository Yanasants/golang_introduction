package main

import (
	"testing"
)

var initialY02 float64 = 2.0
var initialY12 float64 = -2.0
var initialX2 float64 = 0.0
var finalX2 float64 = 1.0
var functionDomain2 float64 = 5000.0
var h2 float64 = (finalX2 - initialX2) / functionDomain2
var allX2 []float64

func TestEquationsSystem2Index1000(t *testing.T) {
	// All x values
	for i := initialX2; i <= finalX2; i += h2 {
		allX2 = append(allX2, i)
	}

	// Method
	y0 := eullerMethod(allX2, h2, initialY02, initialY12, functionDomain2)[0]
	y1 := eullerMethod(allX2, h2, initialY02, initialY12, functionDomain2)[1]

	y0i0 := y0[1000]
	y1i0 := y1[1000]

	if y0i0 < 6.70 || y0i0 > 6.71 {
		t.Errorf("Value %v out of expected range.", y0i0)
	}
	if y1i0 < 4.02 || y1i0 > 4.03 {
		t.Errorf("Value %v out of expected range.", y1i0)
	}
}

func TestEquationsSystem2Index2000(t *testing.T) {
	// All x values
	for i := initialX2; i <= finalX2; i += h2 {
		allX2 = append(allX2, i)
	}

	// Method
	y0 := eullerMethod(allX2, h2, initialY02, initialY12, functionDomain2)[0]
	y1 := eullerMethod(allX2, h2, initialY02, initialY12, functionDomain2)[1]

	y0i0 := y0[2000]
	y1i0 := y1[2000]

	if y0i0 < 4.49 || y0i0 > 4.50 {
		t.Errorf("Value %v out of expected range.", y0i0)
	}
	if y1i0 < 2.69 || y1i0 > 2.70 {
		t.Errorf("Value %v out of expected range.", y1i0)
	}
}

func TestEquationsSystem2Index3000(t *testing.T) {
	// All x values
	for i := initialX2; i <= finalX2; i += h2 {
		allX2 = append(allX2, i)
	}

	// Method
	y0 := eullerMethod(allX2, h2, initialY02, initialY12, functionDomain2)[0]
	y1 := eullerMethod(allX2, h2, initialY02, initialY12, functionDomain2)[1]

	y0i0 := y0[3000]
	y1i0 := y1[3000]

	if y0i0 < 3.0 || y0i0 > 3.02 {
		t.Errorf("Value %v out of expected range.", y0i0)
	}
	if y1i0 < 1.80 || y1i0 > 1.81 {
		t.Errorf("Value %v out of expected range.", y1i0)
	}
}

func TestEquationsSystem2Index4000(t *testing.T) {
	// All x values
	for i := initialX2; i <= finalX2; i += h2 {
		allX2 = append(allX2, i)
	}

	// Method
	y0 := eullerMethod(allX2, h2, initialY02, initialY12, functionDomain2)[0]
	y1 := eullerMethod(allX2, h2, initialY02, initialY12, functionDomain2)[1]

	y0i0 := y0[4000]
	y1i0 := y1[4000]

	if y0i0 < 2.01 || y0i0 > 2.02 {
		t.Errorf("Value %v out of expected range.", y0i0)
	}
	if y1i0 < 1.21 || y1i0 > 1.22 {
		t.Errorf("Value %v out of expected range.", y1i0)
	}
}
