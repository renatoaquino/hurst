package main

//OLS calculates the Ordinary Least Square
func OLS(x, y []float64) (m float64, b float64) {
	var sx, sy, sxy, sxp2 float64
	for i := range x {
		sx += x[i]
		sy += y[i]
		sxy += x[i] * y[i]
		sxp2 += x[i] * x[i]
	}
	n := float64(len(x))

	m = ((n * sxy) - (sx * sy)) / ((n * sxp2) - (sx * sx))
	b = (sy - m*sx) / n

	return m, b
}
