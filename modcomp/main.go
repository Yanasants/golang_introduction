package main

func main() {
	// Initial parameters
	initialY0 := 1.0
	initialY1 := 0.0
	initialX := 0.0
	finalX := 1.0
	functionDomain := 1000.0

	// Calculating the step
	h := (finalX - initialX) / functionDomain

	// All x values
	var allX []float64
	for i := initialX; i <= finalX; i += h {
		allX = append(allX, i)
	}

	// Method
	y0 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[0]
	y1 := eullerMethod(allX, h, initialY0, initialY1, functionDomain)[1]

	// Output
	plotGraph(allX, y0, y1, "X(t), Y(t)")
}
