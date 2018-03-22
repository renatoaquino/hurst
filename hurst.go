package main

import (
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

//CumDev calculates the cumulative deviation of a slice of numbers
func CumDev(values []float64, mean float64) []float64 {
	res := make([]float64, len(values)+1)
	for i := 1; i <= len(values); i++ {
		res[i] = res[i-1] + values[i-1] - mean
	}
	return res[1:]
}

//Range calculates the trustable range of the float64 slice
//It's the equivalent of Log2(R(n)/S(n))
func Range(values []float64) float64 {
	mean, stddev := stat.MeanStdDev(values, nil)
	deviations := CumDev(values, mean)
	return math.Log2((floats.Max(deviations) - floats.Min(deviations)) / stddev)
}

//Hurst returns the Hurst Exponent for a given series
func Hurst(values []float64) float64 {
	l := len(values)
	var y = make([]float64, l-2)
	var x = make([]float64, l-2)

	for i := 2; i < l; i++ {
		y = append(y, Range(values[0:i]))
		x = append(x, math.Log2(float64(i)))
	}

	h, _ := OLS(x, y)

	return math.Abs(h)
}

//Split separates the values into "almost" equal parts
// func Split(values []float64, parts int) ([][]float64, error) {
// 	var s, e, l, k int
// 	var res [][]float64
// 	l = len(values)
// 	if l < parts {
// 		return nil, errors.New("unable to split the values, number of parts greater then length of the items")
// 	}

// 	//fmt.Fprintf(os.Stderr, "Split %d elements into %d parts\n", len(values), parts)
// 	e = l / parts
// 	res = make([][]float64, parts)
// 	k = e
// 	for i := 0; i < parts-1; i++ {
// 		//fmt.Fprintf(os.Stderr, "%d -> %d:%d\n", i, s, k)
// 		res[i] = values[s:k]
// 		s = k
// 		k += e
// 	}

// 	res[parts-1] = values[s:]
// 	return res, nil
// }

// //Partition splits the slice of floats for post processing
// //The logic splits the data into nearly equal len(values)/2 groups
// func Partition(values []float64) ([][][]float64, error) {
// 	var res [][][]float64
// 	var err error
// 	divisions := len(values)

// 	res = make([][][]float64, divisions)
// 	for i := 1; i <= divisions; i++ {
// 		res[i-1], err = Split(values, i)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return res, nil
// }
