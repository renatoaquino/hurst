package main

import (
	"testing"

	"gonum.org/v1/gonum/floats"
)

var olsTests = []struct {
	x []float64
	y []float64
	m float64
	b float64
}{
	{
		x: []float64{
			43, 48, 56, 61, 67, 70,
		},
		y: []float64{
			128, 120, 135, 143, 141, 152,
		},
		m: 0.9644,
		b: 81.0481,
	},
	{
		x: []float64{
			1.0, 1.584, 2.0, 2.321, 2.584, 2.807,
			3.0, 3.169, 3.321, 3.321, 3.584, 3.7,
			3.807, 3.906, 4.0, 4.087, 4.169, 4.247,
			4.321, 4.392, 4.459, 4.523, 4.584, 4.643,
			4.7, 4.725, 4.807, 4.857, 4.906, 4.954,
			5.0,
		},
		y: []float64{
			0.0, 0.292, 0.839, 1.084, 1.397, 1.626,
			1.749, 1.926, 1.955, 1.955, 2.027, 2.114,
			2.208, 2.284, 2.376, 2.540, 2.607, 2.732,
			2.861, 2.952, 2.987, 2.834, 2.795, 2.834,
			2.912, 2.994, 3.080, 3.157, 3.230, 3.297,
			3.237,
		},
		m: 0.8123,
		b: -0.7918,
	},
}

func TestOLS(t *testing.T) {
	for _, test := range olsTests {
		m, b := OLS(test.x, test.y)
		m = floats.Round(m, 4)
		b = floats.Round(b, 4)
		if m != test.m || b != test.b {
			t.Fatalf("Expecting: %f & %f , got %f %f\n", test.m, test.b, m, b)
		}
	}
}

var rm, rb float64

func BenchmarkOLS(b *testing.B) {
	var y = make([]float64, len(brown))
	for i := range brown {
		y[i] = float64(i + 1)
	}
	b.Run("OLS", func(b *testing.B) {
		var sm, sb float64
		for n := 0; n < b.N; n++ {
			sm, sb = OLS(brown, y)
		}
		rm = sm
		rb = sb
	})
}
