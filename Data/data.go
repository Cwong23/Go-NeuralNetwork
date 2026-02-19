package data

import "fmt"

type DataSet struct {
	a float64
	b float64
	c float64
}

func (d DataSet) calculateY(x float64) float64 {
	return (d.a * x * x) + (d.b * x) + d.c
}

func (d DataSet) printEquation() {
	fmt.Printf("Equation: %fx^2 + %fx + %f", d.a, d.b, d.c)
}
