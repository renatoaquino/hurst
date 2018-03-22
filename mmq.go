package main

import "math/big"

//OLS calculates the Ordinary Least Square
//The underlying method used was selected
//to be SumOLS based on the benchmarks.
//If you will use extraordinary large numbers,
//use SumOLSBig instead
//The order of the return values is m and b
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

//MeanOLS calculates the Ordinary Least Square
//os a given set of points and gives the
// values os m and b of the line equation
// y = mx + b
//This function uses the version os the OLS that
//is based on means of x and y
func MeanOLS(x, y []float64) (float64, float64) {
	var sx, sy, sxy, sxp2 float64
	for i := range x {
		sx += x[i]
		sy += y[i]
		sxy += x[i] * y[i]
		sxp2 += x[i] * x[i]
	}
	n := float64(len(x))
	my := sy / n
	mx := sx / n
	xp2m := sxp2 / n
	mxy := sxy / n

	m := ((mx * my) - mxy) / ((mx * mx) - xp2m)
	b := my - m*mx

	return m, b
}

//SumOLS calculates the Ordinary Least Square
//os a given set of points and gives the
// values os m and b of the line equation
// y = mx + b
//This function uses the version os the OLS that
//is based on sums of x and y
//The order of the return values is m and b
func SumOLS(x, y []float64) (float64, float64) {
	var sx, sy, sxy, sxp2 float64
	for i := range x {
		sx += x[i]
		sy += y[i]
		sxy += x[i] * y[i]
		sxp2 += x[i] * x[i]
	}
	n := float64(len(x))

	m := ((n * sxy) - (sx * sy)) / ((n * sxp2) - (sx * sx))
	b := (sy - m*sx) / n

	return m, b
}

//SumOLSBig calculates the Ordinary Least Square
//os a given set of points and gives the
// values os m and b of the line equation
// y = mx + b
//This function uses the version os the OLS that
//is based on sums of x and y
func SumOLSBig(x, y []float64) (float64, float64) {
	var sx = big.NewFloat(0)
	var sy = big.NewFloat(0)
	var sxy = big.NewFloat(0)
	var sxp2 = big.NewFloat(0)
	for i := range x {
		sx.Add(sx, big.NewFloat(x[i]))
		sy.Add(sy, big.NewFloat(y[i]))
		sxy.Add(sxy, big.NewFloat(x[i]*y[i]))
		sxp2.Add(sxp2, big.NewFloat(x[i]*x[i]))
	}
	n := big.NewFloat(float64(len(x)))

	var dividend = big.NewFloat(0)
	dividend.Mul(n, sxy)

	var sxsy = big.NewFloat(0)
	sxsy.Mul(sx, sy)

	dividend.Sub(dividend, sxsy)

	var divisor = big.NewFloat(0)
	divisor.Mul(n, sxp2)

	var sxsx = big.NewFloat(0)
	sxsx.Mul(sx, sx)

	divisor.Sub(divisor, sxsx)

	var m = new(big.Float).Quo(dividend, divisor)
	var b = new(big.Float).Mul(m, sx)
	b.Sub(sy, b)
	b.Quo(b, n)

	mf, _ := m.Float64()
	mb, _ := b.Float64()
	return mf, mb
}
