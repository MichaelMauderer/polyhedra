package polyhedra

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
