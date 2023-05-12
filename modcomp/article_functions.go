package main

import (
	"math"
)

func equationsSystem1(t float64, y0 float64, y1 float64) []float64 {

	dxdt := 998*y0 + 1998*y1
	dydt := -999*y0 - 1999*y1

	return []float64{dxdt, dydt}

}

func analyticsFunction1(t float64, y0 float64, y1 float64) []float64 {

	xt := 2*math.Exp(-t) - math.Exp(-1000*t)
	yt := math.Exp(-1000*t) - math.Exp(-t)

	return []float64{xt, yt}
}

func equationsSystem2(t float64, y0 float64, y1 float64) []float64 {
	Ut := y0
	Vt := y1

	dUdt := 1195*Ut - 1995*Vt
	dVdt := 1197*Ut - 1997*Vt

	return []float64{dUdt, dVdt}
}
