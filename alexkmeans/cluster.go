package alexkmeans

import "fmt"

type Cluster struct {
	Vectors []*Vector
	Mean    Vector
}

// This function will update the cluster centroid depending on what it has assigned.
func (c *Cluster) update() {
	// Check all the assigned vectors, find the new mean.
	// Drop all vectors.
	c.Mean = c.calculateCentroid()
}

func (c *Cluster) Assign(v *Vector) {
	c.Vectors = append(c.Vectors, v)
}

func (c *Cluster) clearVectors() {
	// Force gc
	c.Vectors = nil
}

func (c *Cluster) calculateCentroid() Vector {
	points := c.Vectors
	deno := len(points)
	if deno == 0 {
		return c.Mean
	}
	x, y, z := 0.0, 0.0, 0.0
	for p := range points {
		x += float64((points)[p].X)
		y += float64((points)[p].Y)
		z += float64((points)[p].Z)
	}
	resultantVector := Vector{
		X: x / float64(deno),
		Y: y / float64(deno),
		Z: z / float64(deno),
	}
	fmt.Printf("Moving cluster %s to %s\n", &c.Mean, resultantVector.String())

	return resultantVector
}
