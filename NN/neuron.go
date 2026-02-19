package NN

import (
	"math"
	"math/rand/v2"
)

/*

Given a set of random points on a predetermined quadratic function,
this should predict that line

*/

type Neuron struct {
	Weights []float64
	Bias    float64
	Output  float64
	Input   []float64
}

func NewNeuron(numInputs int) *Neuron {
	weights := make([]float64, numInputs)
	for i := range weights {
		weights[i] = rand.Float64()*2 - 1
	}

	return &Neuron{
		Weights: weights,
		Bias:    0.01,
		Output:  0,
		Input:   make([]float64, numInputs),
	}
}

type NeuralNetwork struct {
	HiddenLayer []Neuron
	OutputLayer Neuron
}

func NewNeuralNetwork() *NeuralNetwork {
	hiddenNeurons := make([]Neuron, 3)
	for i := range hiddenNeurons {
		hiddenNeurons[i] = *NewNeuron(1)
	}
	return &NeuralNetwork{
		HiddenLayer: hiddenNeurons,
		OutputLayer: *NewNeuron(3),
	}
}

func (n *Neuron) predict(inputs []float64) float64 {
	n.Input = inputs
	var sum float64
	for i, x := range inputs {
		sum += x * n.Weights[i]
	}
	n.Output = sum + n.Bias
	return n.Output
}

func (n *Neuron) UpdateWeights(gradient float64, learningRate float64) {
	for i := range n.Weights {
		n.Weights[i] += learningRate * gradient * n.Input[i]
	}
	n.Bias += learningRate * gradient
}

func (nn *NeuralNetwork) predict(x float64) float64 {
	inputVal := []float64{x}
	hiddenResults := make([]float64, len(nn.HiddenLayer))

	for i := range nn.HiddenLayer {
		raw := nn.HiddenLayer[i].predict(inputVal)
		hiddenResults[i] = math.Max(0, raw)
	}

	finalResult := nn.OutputLayer.predict(hiddenResults)

	return finalResult
}

func (nn *NeuralNetwork) backPropagation(y float64, pred float64, learningRate float64) {
	gradient := y - pred
	for i := range nn.HiddenLayer {
		weightToOutput := nn.OutputLayer.Weights[i]
		reluDerivative := 0.0
		if nn.HiddenLayer[i].Output > 0 {
			reluDerivative = 1.0
		}

		hiddenGradient := gradient * weightToOutput * reluDerivative
		nn.HiddenLayer[i].UpdateWeights(hiddenGradient, learningRate)
	}
	nn.OutputLayer.UpdateWeights(gradient, learningRate)
}
