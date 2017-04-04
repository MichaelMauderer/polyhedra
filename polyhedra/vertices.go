package polyhedra

import (
	"fmt"
	"github.com/MichaelMauderer/geometry/r3"
)

type Vertex uint

var vertexId Vertex = 0
var vertexPositions = make(map[Vertex]r3.Point3D)

func NewVertex() Vertex {
	vertexId++
	return vertexId
}

func (v Vertex) setPosition(coords r3.Point3D) {
	vertexPositions[v] = coords
}

func (v Vertex) Position() r3.Point3D {
	return vertexPositions[v]
}

func (v Vertex) String() string {
	return fmt.Sprintf("Vertex(id=%v, pos=%v)", uint(v), v.Position())
}

func SortedClockwise(vertices []Vertex) []Vertex {
	//Insertion sort based on clockwiseness
	c := vertexCentroid(vertices)
	n := r3.Point3D{X: 0, Y: 0, Z: 0}.VectorTo(c).Normalised()
	sorted := make([]Vertex, 1)
	sorted[0] = vertices[0]
	for _, v := range vertices[1:] {
		i := 0
		for ; ; i++ {
			if i == len(sorted) {
				break
			}
			vo := sorted[i]
			if !r3.IsCCW(v.Position(), vo.Position(), c, n) {
				break
			}
		}
		//insert at i
		sorted = append(sorted, 0)
		copy(sorted[i+1:], sorted[i:])
		sorted[i] = v
	}
	return sorted
}

func vertexCentroid(vertices []Vertex) r3.Point3D {
	positions := make([]r3.Point3D, len(vertices))
	for i, v := range vertices {
		positions[i] = v.Position()
	}
	return r3.Centroid3D(positions)
}
