package main

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/stat"
)

//Mean calculates the mean of a float64 slice
func Mean(values []float64) float64 {
	var res float64
	for i := range values {
		res += values[i]
	}
	return res / float64(len(values))
}

//Deviations calculates the deviation of each number in a series against a mean value
func Deviations(values []float64, mean float64) []float64 {
	res := make([]float64, len(values)+1)
	for i := 1; i <= len(values); i++ {
		res[i] = res[i-1] + values[i-1] - mean
	}
	return res[1:]
}

//Sum summates all the elements of a float64 slice
func Sum(values []float64) float64 {
	var res float64
	for i := range values {
		res += values[i]
	}
	return res
}

//Max returns the maximum value of the float64 slice
func Max(values []float64) float64 {
	var res float64
	for i := range values {
		res = math.Max(values[i], res)
	}
	return res
}

//Min returns the minimum value of the float64 slice
func Min(values []float64) float64 {
	var res float64
	for i := range values {
		res = math.Min(values[i], res)
	}
	return res
}

//StdDev returns the standard deviation for a float64 slice
func StdDev(values []float64, mean float64) float64 {
	var res float64
	for i := range values {
		res += math.Pow(values[i]-mean, 2)
	}
	return math.Sqrt(1.0/float64(len(values))) * res
}

//Range calculates the trustable range of the float64 slice
func Range(values []float64) float64 {
	mean := stat.Mean(values, nil)
	deviations := Deviations(values, mean)
	return (Max(deviations) - Min(deviations)) / StdDev(values, mean)
}

//Split separates the values into "almost" equal parts
func Split(values []float64, parts int) ([][]float64, error) {
	var s, e, l, k int
	var res [][]float64
	l = len(values)
	if l < 2*parts {
		return nil, errors.New("unable to split the values, number of parts greater then length of the items")
	}

	e = l / parts
	res = make([][]float64, parts)
	k = e
	for i := 0; i < parts-1; i++ {
		res[i] = values[s:k]
		s = k
		k += e
	}

	res[parts-1] = values[s:]
	return res, nil
}

func Partition(values []float64) ([][][]float64, error) {
	var res [][][]float64
	var err error
	l := float64(len(values))

	//Finds the ideal amount of groups to work with
	//We want to have 1/64 groups of at least 2*divisor
	var groups int
	for i := float64(6); i > 0; i-- {
		if l >= math.Pow(2, i) {
			groups = int(i) + 1
			break
		}
	}

	res = make([][][]float64, groups)
	for i := 0; i < groups; i++ {
		res[i], err = Split(values, int(math.Pow(2, float64(i))))
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

//Hurst returns the Hurst Exponent for a give series
func Hurst(values []float64) (float64, error) {
	var res [][]float64
	groups, err := Partition(values)
	if err != nil {
		return 0.0, err
	}
	res = make([][]float64, len(groups))
	for i := range groups {
		for j := range groups[i] {
			res[i] = append(res[i], Range(groups[i][j]))
		}
	}

	var h float64

	//A lista de pontos res deve ser plotada
	//como o Y do grafico, X será 1,2,3 (log2(numero de itens do grupo))
	//O valor do hurst será a altura formada entre X e Y
	//Calcular o metodo dos minimos quadrados.

	return h, nil
}

func main() {

}
