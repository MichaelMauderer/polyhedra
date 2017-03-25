package polyhedra

import "math"

type CartesianCoordinate struct {
	x, y, z float64
}

type UnitSphereCoordinates struct {
	rhi, phi float64
}

func Centroid(points ...CartesianCoordinate) CartesianCoordinate {
	x, y, z := 0.0, 0.0, 0.0
	for _, p := range points {
		x += p.x
		y += p.y
		z += p.z
	}
	x /= float64(len(points))
	y /= float64(len(points))
	z /= float64(len(points))

	return CartesianCoordinate{x, y, z}
}

func WeightedCentroid(points []CartesianCoordinate, weights []float64) CartesianCoordinate {
	x, y, z := 0.0, 0.0, 0.0
	wSum := 0.0
	for i, p := range points {
		w := weights[i]
		wSum += w
		x += p.x * w
		y += p.y * w
		z += p.z * w
	}
	x /= float64(len(points)) * wSum
	y /= float64(len(points)) * wSum
	z /= float64(len(points)) * wSum

	return CartesianCoordinate{x, y, z}
}

func Distance(p1 CartesianCoordinate, p2 CartesianCoordinate) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
