package main

import "testing"

func TestAvg(t *testing.T) {
	type avgtest struct {
		data   []float64
		answer float64
	}
	avgTestFigures := []avgtest{
		avgtest{[]float64{1, 2}, 1.5},
		avgtest{[]float64{1, 1, 2, 2, 3, 3}, 2},
		avgtest{[]float64{1.2, 1.4, 1.6, 1.8, 2.0, 2.2}, 1.7},
	}
	for _, v := range avgTestFigures {
		a := avg(v.data)
		if a != v.answer {
			t.Error("Expected", v.answer, "got", a)
		}

	}
}
