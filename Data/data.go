package data

import "fmt"

type DataSet struct {
	A float64
	B float64
	C float64
}

func NewDataSet(a float64, b float64, c float64) *DataSet {
	return &DataSet{
		A: a,
		B: b,
		C: c,
	}
}

func (d DataSet) calculateY(x float64) float64 {
	return (d.A * x * x) + (d.B * x) + d.C
}

func (d DataSet) printEquation() {
	fmt.Printf("Equation: %fx^2 + %fx + %f", d.A, d.B, d.C)
}
