package r2

type Polygon interface {
}

type Point2D struct {
	X, Y float64
}

type Vector2D struct {
	X, Y float64
}

func Centroid2D(points []Point2D) Point2D {
	x, y := 0.0, 0.0
	for _, p := range points {
		x += p.X
		y += p.Y
	}
	x /= float64(len(points))
	y /= float64(len(points))

	return Point2D{x, y}
}

func ccLess(a, b, center Point2D) bool {
	dax := a.X-center.X
	dbx := b.X-center.X
	// http://stackoverflow.com/a/6989383/1175813
	if  dax >= 0.0 &&  dbx < 0.0 {
		return true
	}
	if dax < 0.0 && dbx >= 0.0 {
		return false
	}
	if dax == 0.0 && dbx == 0.0 {
		if a.Y-center.Y >= 0 || b.Y-center.Y >= 0 {
			return a.Y > b.Y
		}
		return b.Y > a.Y
	}

	// compute the cross product of vectors (center -> a) X (center -> b)
	det := dax*(b.Y-center.Y) - dbx*(a.Y-center.Y)
	if det < 0 {
		return true
	}
	if det > 0 {
		return false
	}

	// points a and b are on the same line from the center
	// check which point is closer to the center
	d1 := (a.X-center.X)*(a.X-center.X) + (a.Y-center.Y)*(a.Y-center.Y)
	d2 := (b.X-center.X)*(b.X-center.X) + (b.Y-center.Y)*(b.Y-center.Y)
	return d1 > d2
}

