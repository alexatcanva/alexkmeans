package alexkmeans

import (
	"math"
)

func Calculate3DDistance(pointA, pointB Vector) float64 {
	var c1, c2, c3 float64
	c1 = math.Pow((float64(pointA.X) - float64(pointB.X)), 2)
	c2 = math.Pow((float64(pointA.Y) - float64(pointB.Y)), 2)
	c3 = math.Pow((float64(pointA.Z) - float64(pointB.Z)), 2)
	res := math.Sqrt(c1 + c2 + c3)
	return res
}
