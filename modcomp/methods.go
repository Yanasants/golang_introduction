package main

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
