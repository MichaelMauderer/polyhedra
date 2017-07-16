package polyhedra

import (
	"fmt"
	"github.com/MichaelMauderer/polyhedra/r3"
)

// Vertex represents a point within a Polyhedron where edges meet.
type Vertex uint

// vertexId is the global id counter that is used to generate unique ids fo vertices.
var vertexId Vertex = 0

// vertexPositions contains the position for each vertex.
var vertexPositions = make(map[Vertex]r3.Point)

// NewVertex creates an new vertex.
func NewVertex() Vertex {
	vertexId++
	return vertexId
}

// setPosition sets the position of the vertex.
func (v Vertex) setPosition(coords r3.Point) {
	vertexPositions[v] = coords
}

// Position returns the position of the vertex.
func (v Vertex) Position() r3.Point {
	return vertexPositions[v]
}

// String returns the string representation of the vertex.
func (v Vertex) String() string {
	return fmt.Sprintf("Vertex(id=%v, pos=%v)", uint(v), v.Position())
}

// SortedClockwise sorts the vertices clockwise around their geometric center.
func SortedClockwise(vertices []Vertex) []Vertex {
	//Insertion sort based on clockwiseness
	c := vertexCentroid(vertices)
	// The normal of the plane of sorting is defined by the vector from zero to the geometric center.
	n := r3.Point{X: 0, Y: 0, Z: 0}.VectorTo(c).Normalised()
	sorted := make([]Vertex, 1)
	// The initial vertex is chosen as the first vertex in the slice.
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

// vertexCentroid computes the centroid of the given vertices.
func vertexCentroid(vertices []Vertex) r3.Point {
	positions := make([]r3.Point, len(vertices))
	for i, v := range vertices {
		positions[i] = v.Position()
	}
	return r3.Centroid3D(positions)
}

func minVertexIndex(v []Vertex) int{
	var min int
	var m Vertex
	for i, e := range v {
		if e < m {
			m = e
			min = i
		}
	}
	return min
}
