package alexkmeans

import (
	"fmt"
	"image"
	"math/rand"
	"sort"
	"time"
)

/*
This file is responsible for the quantization of a Image.image object
*/

func CalculateKMeansGrouping(image image.Image, passes, groups int) []Vector {

	bounds := image.Bounds()
	vectoryArray := getVectorArray(bounds, image)
	clusters := selectInitialPartitions(image, bounds.Max.X, bounds.Max.Y, groups)

	clusterPointers := make([]*Cluster, len(clusters))

	for c := range clusters {
		clusterPointers[c] = &(clusters[c])
	}

	for p := 0; p < passes; p++ {

		fmt.Printf("\nRunning pass %d\n", p)

		for v1 := range vectoryArray {
			for v2 := range vectoryArray[v1] {
				v := vectoryArray[v1][v2]

				v.assign(clusterPointers)
			}
		}

		for c := range clusters {
			clusters[c].update()
			clusters[c].clearVectors()
		}
	}

	res := make([]Vector, len(clusters))
	for c := range clusters {
		res[c] = clusters[c].Mean
	}

	// sort
	sort.Slice(res, func(i, j int) bool {
		return res[i].magnitude() < res[j].magnitude() &&
			res[i].X < res[j].X
	})

	return res
}

func selectInitialPartitions(image image.Image, width, height, groups int) []Cluster {

	centroids := make([]Cluster, groups)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < groups; i++ {
		r, g, b, _ := image.At(rand.Intn(width), rand.Intn(height)).RGBA()
		centroids[i].Mean = Vector{
			X: float64(r >> 8),
			Y: float64(g >> 8),
			Z: float64(b >> 8),
		}
	}

	return centroids
}

func getVectorArray(bounds image.Rectangle, image image.Image) [][]Vector {

	width, height := bounds.Max.X, bounds.Max.Y

	// Bad performance, but iterate over image til we have a 2D list of vectors
	vecArr := make([][]Vector, width)
	for x := 0; x < width; x++ {
		row := make([]Vector, height)
		for y := 0; y < height; y++ {
			r, g, b, _ := image.At(x, y).RGBA()
			row[y] = Vector{
				X: float64(r >> 8),
				Y: float64(g >> 8),
				Z: float64(b >> 8),
			}
		}
		vecArr[x] = row
	}

	return vecArr
}
