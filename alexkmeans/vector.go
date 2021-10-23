package alexkmeans

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

// The Assign function will allow us to assign the vector to the nearest
func (v *Vector) assign(clusters []*Cluster) {

	var (
		distance float64
		nearest  int
	)

	// Calculate the minimum distance
	for c := range clusters {
		cd := Calculate3DDistance((clusters[c]).Mean, *v)
		if c == 0 {
			// assign very first distance as the min
			distance = cd
			nearest = c
			continue
		}
		if cd < distance {
			distance = cd
			nearest = c
		}
		// fmt.Printf("Vector %s is nearest to %s by distance %f\n", (*v).String(), (clusters[c].Mean).String(), distance)
	}

	clusters[nearest].Assign(v)
}

func (v *Vector) String() string {
	return fmt.Sprintf("(%f,%f,%f)", v.X, v.Y, v.Z)
}

func (v *Vector) magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}
