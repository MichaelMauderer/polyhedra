package polyhedra

import (
	"errors"
	"math"
	"sort"
)

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

func CounterClockwise2D(v []Point2D) sort.Interface {
	cc := counterClockwise{}
	cc.v = v
	cc.center = Centroid2D(v)
	return cc
}

type counterClockwise struct {
	v      []Point2D
	center Point2D

}

func (v counterClockwise) Len() int      { return len(v.v) }
func (v counterClockwise) Swap(i, j int) { v.v[i], v.v[j] = v.v[j], v.v[i] }
func (v counterClockwise) Less(i, j int) bool {
	return ccLess(v.v[i], v.v[j], v.center)
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

func SubdivideTriangle(triangle Triangle, n, m int) ([]Triangle, error) {
	if m == n {
		return nil, errors.New("Class II not supported")
	}
	if m != 0 {
		return nil, errors.New("Class III not supported")
	}
	if n == 1 {
		return []Triangle{triangle}, nil
	}

	vertexRows := make([][]Vertex, n+1)

	rowSize := 1
	for row := 0; row < n+1; row++ {
		vertexRows[row] = make([]Vertex, rowSize)
		for n := range vertexRows[row] {
			vertexRows[row][n] = NewVertex()
		}
		rowSize++
	}

	vertexRows[0][0] = triangle.vertices[0]
	vertexRows[n][0] = triangle.vertices[1]
	vertexRows[n][n] = triangle.vertices[2]

	resultTriangles := make([]Triangle, 0)

	// Walk through the rows of the triangle
	// Connect the vertices above and below
	for row := 0; row < n; row++ {
		for i, vertex := range vertexRows[row] {
			v1 := vertexRows[row+1][i]
			v2 := vertexRows[row+1][i+1]
			t := Triangle{[3]Vertex{vertex, v1, v2}}
			resultTriangles = append(resultTriangles, t)
		}
	}
	for row := 1; row < n; row++ {
		for i := 0; i < len(vertexRows[row])-1; i++ {
			v0 := vertexRows[row][i]
			v1 := vertexRows[row][i+1]
			v2 := vertexRows[row+1][i+1]
			t := Triangle{[3]Vertex{v0, v1, v2}}
			resultTriangles = append(resultTriangles, t)
		}
	}
	return resultTriangles, nil

}

type Triangle struct {
	vertices [3]Vertex
}

type Pentagon struct {
	edges    [5]Edge
	vertices [5]Vertex
}

type RegularHexagon struct {
}

func (rh RegularHexagon) EdgeLength() float64 {
	return math.Sqrt(3) * 2.0 / 3.0
}

func NewUnitHexagon() Polygon {
	return RegularHexagon{}
}
