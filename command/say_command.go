package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/wdalmut/digit/nn"
	"gopkg.in/wdalmut/cli.v1"
)

type SayCommand struct{}

func (h SayCommand) Help() string {
	return `NAME:
   say - predict an handwritten digit
USAGE:
   command say [command options] [arguments...]
OPTIONS:
    -i, --in, -in "" The input image
	-l, --load, -load "" Load an existing group of weights and biases
`
}

func (h SayCommand) Run(args []string) int {
	flag := cli.Parse(args)
	input := flag.String("", "-i", "--in", "-in")
	weights := flag.String("", "-w", "--weights", "-weights")
	biases := flag.String("", "-b", "--biases", "-biases")

	net := nn.NewNetwork([]int{768, 30, 10})

	if weights != "" && biases != "" {
		log.Printf("Loading existing weights...")
		var w [][][]float64

		data, err := ioutil.ReadFile(weights)
		if err != nil {
			log.Printf("Unable to read weights file %s", weights)
			return 1
		}
		json.Unmarshal(data, &w)

		var b [][]float64

		data, err = ioutil.ReadFile(biases)
		if err != nil {
			log.Printf("Unable to read biases file %s", weights)
			return 1
		}
		json.Unmarshal(data, &b)

		net.Load(w, b)
		log.Printf("Loading completed...")
	}

	image, err := ioutil.ReadFile(input)

	if err != nil {
		log.Printf("%s -> %s\n", err.Error(), input)
		return 1
	}

	log.Printf("Image normalization...\n")
	normalized := normalize(image)
	log.Printf("Image normalization completed...\n")

	activation := net.Activate(normalized)

	fmt.Printf("It is a: %v\n", activation)

	return 0
}

func normalize(image []byte) []float64 {
	var normalized []float64
	for _, v := range image {
		normalized = append(normalized, float64(v)/255)
	}

	return normalized
}

func (r SayCommand) Synopsis() string {
	return "Predict an handwritten digit"
}

func SayCommandFactory() cli.Command {
	return SayCommand{}
}
