package main

import (
	"fmt"
	//"io/ioutil"
	//"log"
	"encoding/json"
	"math"
	"os"
)

func formData() ([][]float64, []float64) {
	a := [][]float64{
		[]float64{81, -45, 45},
		[]float64{-45, 50, -15},
		[]float64{45, -15, 38},
	}
	b := []float64{531, -460, 193}
	return a, b
}

//MatrixSpaceSeparatedReader dv
type MatrixSpaceSeparatedReader struct {
	A [][]float64
	b []float64
}

//MatrixReader dv
type MatrixReader interface {
	Read() MatrixSpaceSeparatedReader
}

//JSONSLinSys struct
type JSONSLinSys struct {
	mat  string
	vect string
}

func (v JSONSLinSys) Read() MatrixSpaceSeparatedReader {
	var A [][]float64
	json.Unmarshal([]byte(v.mat), &A)
	var b []float64
	json.Unmarshal([]byte(v.vect), &b)
	return MatrixSpaceSeparatedReader{A, b}
}

//Read function
func (v MatrixSpaceSeparatedReader) Read() MatrixSpaceSeparatedReader {

	f, _ := os.Open("test.txt")
	defer f.Close()

	var n int
	fmt.Fscan(f, &n)

	A := make([][]float64, n)
	for i := range A {
		A[i] = make([]float64, n)
		for j := range A[i] {
			fmt.Fscan(f, &A[i][j])
		}
	}

	b := make([]float64, n)
	for i := range b {
		fmt.Fscan(f, &b[i])
	}
	return MatrixSpaceSeparatedReader{A, b}
}

func getMatrixU(a [][]float64) [][]float64 {
	u := make([][]float64, len(a))
	for i := 0; i < len(a); i++ {
		if len(u[i]) < 1 {
			u[i] = make([]float64, len(a[i]))
		}
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if i == j {
				sum := 0.0
				for k := 0; k < (i); k++ {
					sum += u[k][i] * u[k][i]
				}
				u[i][j] = math.Sqrt(a[i][i] - sum)
			}
			if (i == 0) && (j > i) {
				u[i][j] = a[i][j] / u[0][0]
			}
			if (j > i) && (i != 0) {
				sum := 0.0
				for k := 0; k < (i); k++ {
					sum += u[k][i] * u[k][j]
				}
				u[i][j] = (a[i][j] - sum) / u[i][i]
			}
			if i > j {
				u[i][j] = 0
			}
		}
	}
	return u
}

func transpose(x [][]float64) [][]float64 {
	out := make([][]float64, len(x[0]))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func findY(u [][]float64, b []float64) []float64 {
	y := make([]float64, len(u))
	for i := 0; i < len(b); i++ {
		sum := 0.0
		for k := 0; k < i; k++ {
			sum += u[i][k] * y[k]
		}
		y[i] = (b[i] - sum) / u[i][i]
	}
	return y
}

func findX(u [][]float64, y []float64) []float64 {
	x := make([]float64, len(u))
	for i := (len(y) - 1); i >= 0; i-- {
		sum := 0.0
		for k := i + 1; k < len(y); k++ {
			sum += u[i][k] * x[k]
		}
		x[i] = (y[i] - sum) / u[i][i]
	}
	return x
}

func solver(a [][]float64, b []float64) []float64 {
	u := getMatrixU(a)
	fmt.Println("Matrix U")
	fmt.Println(u)
	ut := transpose(u)
	y := findY(ut, b)
	x := findX(u, y)
	return x
}

func main() {
	fmt.Println("Hello world")
	//a, b := formData()

	//var inter MatrixReader = MatrixSpaceSeparatedReader{}
	var inter MatrixReader = JSONSLinSys{"[[2,3],[3,5]]", "[8,13]"}
	r := inter.Read()
	a := r.A
	b := r.b
	//a, b = formData()
	fmt.Println(a)
	fmt.Println(b)
	res := solver(a, b)
	fmt.Println(res)
}
