package main

import (
	"math"
	"testing"
)

//TestTranspose testing
func TestTranspose(t *testing.T) {
	a := [][]float64{[]float64{5, 2}, []float64{3, 5}}
	b := [][]float64{[]float64{5, 3}, []float64{2, 5}}
	total := transpose(a)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if total[i][j] != b[i][j] {
				t.Errorf("Transpose was incorrect, got: %d, want: %d.", total, b)
			}
		}
	}
}

//TestSolver testing
func TestSolver1(t *testing.T) {
	a := [][]float64{[]float64{2, 3}, []float64{3, 5}}
	b := []float64{8, 13}
	c := []float64{0.9999999999, 2}
	eps := math.Pow(10, -5)

	total := solver(a, b)
	for i := 0; i < len(total); i++ {
		if math.Abs(total[i]-c[i]) >= eps {
			t.Errorf("Solution is incorrect, got: %d, want: %d.", total, c)
		}
	}
}

func TestSolver2(t *testing.T) {
	a := [][]float64{
		[]float64{81, -45, 45},
		[]float64{-45, 50, -15},
		[]float64{45, -15, 38},
	}
	b := []float64{531, -460, 193}
	c := []float64{6, -5, -4}
	eps := math.Pow(10, -5)

	total := solver(a, b)
	for i := 0; i < len(total); i++ {
		if math.Abs(total[i]-c[i]) >= eps {
			t.Errorf("Solution is incorrect, got: %d, want: %d.", total, c)
		}
	}
}

//TestTranspose testing
func TestGetMatrixU(t *testing.T) {
	a := [][]float64{[]float64{2, 3}, []float64{3, 5}}
	b := [][]float64{[]float64{1.4142135623730951, 2.1213203435596424}, []float64{0, 0.7071067811865481}}
	eps := math.Pow(10, -5)
	total := getMatrixU(a)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if math.Abs(total[i][j]-b[i][j]) >= eps {
				t.Errorf("Decomposition was incorrect, got: %d, want: %d.", total, b)
			}
		}
	}
}
