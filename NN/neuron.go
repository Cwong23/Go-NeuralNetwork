package NN

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

func (n Neuron) predict(inputs []float64) float64 {
	n.Input = inputs
	var sum float64
	for i, x := range inputs {
		sum += x * n.Weights[i]
	}
	n.Output = sum + n.Bias
	return n.Output
}
