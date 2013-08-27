package nn

type NeuralNetwork struct {
	layers  []Layer
	weights []Matrix
}

type Layer struct {
	nodes []int
}

func New(hiddenSizes []int) *NeuralNetwork {
	layers = make([]Layer, len(hiddenSizes))
	for i, size := range hiddenSizes {
		layers[i].nodes = make([]int, size)
	}

	return &NeuralNetwork{layers, nil}
}
