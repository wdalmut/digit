package command

import (

	// register the PNG format with the image package

	"image/png"
	"io/ioutil"
	"math"
	"os"

	"gopkg.in/wdalmut/cli.v1"
)

type ConvertCommand struct{}

func (h ConvertCommand) Help() string {
	return `NAME:
   convert - convert a PNG picture to grayscale
USAGE:
   command convert [command options] [arguments...]
OPTIONS:
    -i, --in, -in "" The input image
	-o, --out, -out "" The output path
`
}

func (h ConvertCommand) Run(args []string) int {
	flag := cli.Parse(args)
	input := flag.String("", "-i", "--in", "-in")
	out := flag.String("", "-o", "--out", "-out")

	if input == "" {
		return 1
	}

	if out == "" {
		return 1
	}

	in, err := os.Open(input)
	if err != nil {
		return 1
	}

	png, err := png.Decode(in)
	if err != nil {
		return 1
	}

	bounds := png.Bounds()

	data := make([]byte, bounds.Max.X*bounds.Max.Y)

	for i := 0; i < bounds.Max.X; i++ {
		for j := 0; j < bounds.Max.Y; j++ {
			color := png.At(j, i)
			r, g, b, _ := color.RGBA()

			gray := int(math.Abs(255 - math.Floor(float64((r+g+b))/(256*3))))

			data[(i*bounds.Max.Y)+j] = byte(gray)
		}
	}

	ioutil.WriteFile(out, data, 0755)

	return 0
}

func (r ConvertCommand) Synopsis() string {
	return "Convert a PNG image to grayscale"
}

func ConvertCommandFactory() cli.Command {
	return ConvertCommand{}
}
