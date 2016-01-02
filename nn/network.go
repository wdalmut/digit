package nn

import (
	"math"
	"math/rand"
)

type Network struct {
	weights [][][]float64
	biases  [][]float64
	sizes   []int
}

func NewNetwork(sizes []int) *Network {

	net := &Network{
		sizes: sizes,
	}

	biases := make([][]float64, len(sizes)-1)

	for k, v := range sizes[1:] {
		for i := 0; i < v; i++ {
			s := rand.NormFloat64()

			biases[k] = append(biases[k], s)
		}
	}
	net.biases = biases

	weights := make([][][]float64, len(sizes)-1)
	for k := 1; k < len(sizes); k++ {
		weights[k-1] = make([][]float64, sizes[k])
		for i := 0; i < sizes[k]; i++ {
			for j := 0; j < sizes[k-1]; j++ {
				s := rand.NormFloat64()

				weights[k-1][i] = append(weights[k-1][i], s)
			}
		}
	}
	net.weights = weights

	return net
}

func (n *Network) Load(weights [][][]float64, biases [][]float64) {
	n.weights = weights
	n.biases = biases
}

func (n Network) FeedForward(intensities []float64) []float64 {
	for layer := 1; layer < len(n.sizes); layer++ {
		newActivation := make([]float64, n.sizes[layer])

		pos := layer - 1
		b := n.biases[pos]
		w := n.weights[pos]

		// w * a + b
		for i, weightList := range w {
			newActivation[i] = 0
			for ii, ww := range weightList {
				newActivation[i] += ww * intensities[ii]
			}

			newActivation[i] += b[i]
		}

		// sigmoid
		for k, v := range newActivation {
			newActivation[k] = 1 / (1 + math.Exp(-v))
		}

		// prepare the next layer
		intensities = newActivation
	}

	// return the last layer activation
	return intensities
}

func (n Network) Activate(intensities []float64) int {
	l := n.FeedForward(intensities)

	num := -1
	max := -2.0
	for k, v := range l {
		if v > max {
			num = k
			max = v
		}
	}

	return num
}
