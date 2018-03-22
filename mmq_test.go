package main

import "testing"

func TestCompareOLS(t *testing.T) {
	var y = make([]float64, len(brown))
	for i := range brown {
		y[i] = float64(i + 1)
	}
	mm, mb := MeanOLS(brown, y)
	sm, sb := SumOLS(brown, y)
	bm, bb := SumOLSBig(brown, y)
	if sm != mm || mb != sb {
		t.Fatalf("Expecting equations to be equal:\ny = %f * X + %f\ny = %f * X + %f\n", mm, mb, sm, sb)
	}
	if sm != bm || bb != sb {
		t.Fatalf("Expecting big and sum equations to be equal:\ny = %f * X + %f\ny = %f * X + %f\n", mm, mb, sm, sb)
	}
}

var testx = []float64{
	1.0,
	1.584,
	2.0,
	2.321,
	2.584,
	2.807,
	3.0,
	3.169,
	3.321,
	3.321,
	3.584,
	3.7,
	3.807,
	3.906,
	4.0,
	4.087,
	4.169,
	4.247,
	4.321,
	4.392,
	4.459,
	4.523,
	4.584,
	4.643,
	4.7,
	4.725,
	4.807,
	4.857,
	4.906,
	4.954,
	5.0,
}

var testy = []float64{
	0.0,
	0.292,
	0.839,
	1.084,
	1.397,
	1.626,
	1.749,
	1.926,
	1.955,
	1.955,
	2.027,
	2.114,
	2.208,
	2.284,
	2.376,
	2.540,
	2.607,
	2.732,
	2.861,
	2.952,
	2.987,
	2.834,
	2.795,
	2.834,
	2.912,
	2.994,
	3.080,
	3.157,
	3.230,
	3.297,
	3.237,
}

func TestOLS(t *testing.T) {
	m, b := OLS(testx, testy)
	if m != 0.812 || b != -0.798 {
		t.Fatalf("Expecting: 0.812 & -0.798, got %f %f\n", m, b)
	}
}

var rm, rb float64

func BenchmarkOLS(b *testing.B) {
	var y = make([]float64, len(brown))
	for i := range brown {
		y[i] = float64(i + 1)
	}
	b.Run("SumOLS", func(b *testing.B) {
		var sm, sb float64
		for n := 0; n < b.N; n++ {
			sm, sb = SumOLS(brown, y)
		}
		rm = sm
		rb = sb
	})
	b.Run("MeanOLS", func(b *testing.B) {
		var mm, mb float64
		for n := 0; n < b.N; n++ {
			mm, mb = MeanOLS(brown, y)
		}
		rm = mm
		rb = mb
	})
	b.Run("SumOLSBig", func(b *testing.B) {
		var sm, sb float64
		for n := 0; n < b.N; n++ {
			sm, sb = SumOLSBig(brown, y)
		}
		rm = sm
		rb = sb
	})
}
