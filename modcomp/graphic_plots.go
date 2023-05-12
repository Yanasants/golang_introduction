package main

import (
	"log"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

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
