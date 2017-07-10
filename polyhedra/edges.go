package polyhedra

import (
	"fmt"

	"github.com/MichaelMauderer/geometry/r3"
)

// Edge represents an edge within a polyhedron.
type Edge interface {
	Length() float64
	Center() r3.Point3D
	Contains(Vertex) bool
	Equal(Edge) bool
	Reversed() Edge

	Vertices() [2]Vertex
}

func NewEdge(v1, v2 Vertex) edge {
	if v1 > v2 {
		return edge{v2, v1}
	} else {
		return edge{v1, v2}
	}
}

type edge struct {
	v1, v2 Vertex
}

// Length returns the length of the edge, defined as the distance between the two end vertices.
func (e edge) Length() float64 {
	return r3.Distance(e.v1.Position(), e.v2.Position())
}

// Center returns the midpoint of the edge.
func (e edge) Center() r3.Point3D {
	return r3.Centroid3D([]r3.Point3D{e.v1.Position(), e.v2.Position()})
}

// Contains check whether either end point equals the given Vertex v.
func (e edge) Contains(v Vertex) bool {
	return e.v1 == v || e.v2 == v
}

// Equal checks whether to edges represent the same edge..
func (e edge) Equal(other Edge) bool {
	ov1, ov2 := other.Vertices()[0], other.Vertices()[1]
	if e.v1 == ov1 && e.v2 == ov2 {
		return true
	}
	return e.v1 == ov2 && e.v2 == ov1
}

// Reversed returns a a version of the edge that has its vertices reversed.
func (e edge) Reversed() Edge {
	redge := edge{e.v2, e.v1}
	return redge
}

// Vertices return the two vertices that make up the edge.
func (e edge) Vertices() [2]Vertex {
	return [2]Vertex{e.v1, e.v2}
}

// String return a string representation of the edge.
func (e edge) String() string {
	return fmt.Sprintf("Edge(%v, %v)", e.v1, e.v2)
}

// CullDuplicates is a helper method that takes a list of edges an removes all duplicate edges.
// BUG: The ordering of edges in the output slice is undefined.
func cullDuplicates(edges []Edge) (uniqueEdges []Edge) {
	edgeSet := make(map[edge]bool, len(edges))

	for _, newEdge := range edges {
		vs := newEdge.Vertices()
		e := NewEdge(vs[0], vs[1])
		if !edgeSet[e] {
			edgeSet[e] = true
			uniqueEdges = append(uniqueEdges, e)
		}
	}
	return
}
