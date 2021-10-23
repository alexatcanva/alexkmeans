package main

import (
	"alexkmeans/alexkmeans"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strings"
)

// A simple program to sort pixels into the correct colours via K-means clustering
// This was written as a sample project to practice GoLang
// K means clustering wikipedia article: https://en.wikipedia.org/wiki/K-means_clustering
// Information on colour quantization can be found here: https://en.wikipedia.org/wiki/Color_quantization

var (
	input    string
	output   string
	clusters int
	passes   int
)

func init() {
	flag.IntVar(&clusters, "clusters", 3, "the amount of clusters to be used")
	flag.IntVar(&passes, "passes", 3, "the amount of passes to be used")
	flag.StringVar(&input, "input", "", "the input file to be used")
	flag.StringVar(&output, "output", "", "the output file to be used")

	flag.Parse()

	if !strings.HasSuffix(output, ".png") {
		output += ".png"
	}
}

func main() {
	data, err := alexkmeans.LoadImage(input)
	if err != nil {
		panic(err)
	}

	results := alexkmeans.CalculateKMeansGrouping(data, passes, clusters)

	drawResults(results)
}

func drawResults(colours []alexkmeans.Vector) {
	fmt.Printf("\n Saving output colours... \n")

	im := image.NewRGBA(image.Rectangle{Max: image.Point{X: 250, Y: 50}})
	var blocks float64 = (250.00 / float64(len(colours)))
	start := 0
	for c := range colours {
		for x := start; x < start+int(math.Ceil(blocks)); x++ {
			for y := 0; y < 50; y++ {
				im.Set(x, y, color.RGBA{uint8(colours[c].X), uint8(colours[c].Y), uint8(colours[c].Y), 255})
			}
		}
		start += int(math.Ceil(blocks))
	}

	imgfile, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer imgfile.Close()

	err = png.Encode(imgfile, im)
	if err != nil {
		panic(err)
	}
}
