package main

import (
	"testing"
)

func TestAnalyticsSystem1Index200(t *testing.T) {
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
		t.Errorf("Value %v out of expected range.", y0i0)
	}
	if y1i0 < -0.82 || y1i0 > -0.81 {
		t.Errorf("Value %v out of expected range.", y1i0)
	}
}

func TestAnalyticsSystem1Index400(t *testing.T) {
	// All x values
	for i := initialX; i <= finalX; i += h {
		allX = append(allX, i)
	}

	// Method
	y0 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[0]
	y1 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[1]

	y0i0 := y0[400]
	y1i0 := y1[400]

	if y0i0 < 1.34 || y0i0 > 1.35 {
		t.Errorf("Value %v out of expected range.", y0i0)
	}
	if y1i0 < -0.68 || y1i0 > -0.67 {
		t.Errorf("Value %v out of expected range.", y1i0)
	}
}

func TestAnalyticsSystem1Index600(t *testing.T) {
	// All x values
	for i := initialX; i <= finalX; i += h {
		allX = append(allX, i)
	}

	// Method
	y0 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[0]
	y1 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[1]

	y0i0 := y0[600]
	y1i0 := y1[600]

	if y0i0 < 1.09 || y0i0 > 1.10 {
		t.Errorf("Value %v out of expected range.", y0i0)
	}
	if y1i0 < -0.55 || y1i0 > -0.54 {
		t.Errorf("Value %v out of expected range.", y1i0)
	}
}

func TestAnalyticsSystem1Index800(t *testing.T) {
	// All x values
	for i := initialX; i <= finalX; i += h {
		allX = append(allX, i)
	}

	// Method
	y0 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[0]
	y1 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[1]

	y0i0 := y0[800]
	y1i0 := y1[800]

	if y0i0 < 0.89 || y0i0 > 0.90 {
		t.Errorf("Value %v out of expected range.", y0i0)
	}
	if y1i0 < -0.45 || y1i0 > -0.44 {
		t.Errorf("Value %v out of expected range.", y1i0)
	}
}
