package main

import (
	"log"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func equationsSystem1(t float64, y0 float64, y1 float64) []float64 {

	dxdt := 998*y0 + 1998*y1
	dydt := -999*y0 - 1999*y1

	return []float64{dxdt, dydt}

}

func eullerMethod(allX []float64, h float64, initialY0 float64, initialY1 float64, functionDomain float64) [][]float64 {

	// All y values
	allYEuller0 := []float64{initialY0}
	allYEuller1 := []float64{initialY1}

	// Euller Method
	for i := 0; i < len(allX)-1; i++ {
		x := allX[i]
		y0 := allYEuller0[i]
		y1 := allYEuller1[i]
		k := equationsSystem1(x, y0, y1)
		allYEuller0 = append(allYEuller0, h*k[0]+allYEuller0[i])
		allYEuller1 = append(allYEuller1, h*k[1]+allYEuller1[i])
	}

	return [][]float64{allYEuller0, allYEuller1}
}

func plotGraph(xs, y0, y1 []float64, graphic_name string) {
	p := plot.New()

	lineY0 := make(plotter.XYs, len(xs))
	for i := range lineY0 {
		lineY0[i].X = xs[i]
		lineY0[i].Y = y0[i]
	}

	lineY1 := make(plotter.XYs, len(xs))
	for i := range lineY1 {
		lineY1[i].X = xs[i]
		lineY1[i].Y = y1[i]
	}

	s0, err := plotter.NewLine(lineY0)
	if err != nil {
		log.Fatal(err)
	}

	s1, err := plotter.NewLine(lineY1)
	if err != nil {
		log.Fatal(err)
	}

	format := ".png"
	png_name := strings.Join([]string{graphic_name, format}, "")

	p.Add(s0, s1)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, png_name); err != nil {
		log.Fatal(err)
	}
}
