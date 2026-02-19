package main

import (
	"Go-NerualNetwork/NN"
	"Go-NerualNetwork/data"
)

func main() {
	model := NN.NewNeuralNetwork()
	myDataset := data.NewDataSet(1.0, 2.0, 3.0)
}
