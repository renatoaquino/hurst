package main

import (
	"math"
	"sync"

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
	return math.Log10((floats.Max(deviations) - floats.Min(deviations)) / stddev)
}

//Hurst returns the Hurst Exponent for a given series
func Hurst(values []float64) float64 {
	l := len(values)
	var y []float64
	var x []float64

	for i := 2; i < l; i++ {
		y = append(y, Range(values[0:i]))
		x = append(x, math.Log10(float64(i)))
	}

	h, _ := OLS(x, y)

	return math.Abs(h)
}

//ConcurrentHurst returns the Hurst Exponent for a given series using concurrency
func ConcurrentHurst(values []float64) float64 {
	l := len(values)
	var y = make([]float64, l-2)
	var x = make([]float64, l-2)

	var wg sync.WaitGroup
	var workers = int(math.Min(float64(l-2), 80))

	type round struct {
		i      int
		values []float64
	}

	ichan := make(chan round, 100)
	quit := make(chan int)

	for i := 0; i < workers; i++ {
		go func() {
			for {
				select {
				case round := <-ichan:
					y[round.i] = Range(round.values)
					wg.Done()
				case <-quit:
					return
				}
			}
		}()
	}

	j := 2
	wg.Add(l - 2)
	for i := 0; i < l-2; i++ {
		ichan <- round{i: i, values: values[0:j]}
		x[i] = math.Log10(float64(j))
		j++
	}
	wg.Wait()
	for i := 0; i < workers; i++ {
		quit <- 0
	}

	h, _ := OLS(x, y)

	return math.Abs(h)
}
