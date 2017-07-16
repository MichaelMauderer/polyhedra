package polyhedra

import (
	"fmt"

	"github.com/MichaelMauderer/polyhedra/r3"
)

// NewEdge creates a new edge from two vertices.
func NewEdge(v1, v2 Vertex) Edge {
	if v1 > v2 {
	    return Edge{v2, v1}
	}
	return Edge{v1, v2}
}

// Edge represents an edge between two vertices.
type Edge struct {
	v1, v2 Vertex
}

// Length returns the length of the Edge, defined as the distance between the two end vertices.
func (e Edge) Length() float64 {
	return r3.Distance(e.v1.Position(), e.v2.Position())
}

// Center returns the midpoint of the Edge.
func (e Edge) Center() r3.Point {
	return r3.Centroid3D([]r3.Point{e.v1.Position(), e.v2.Position()})
}

// Contains check whether either end point equals the given Vertex v.
func (e Edge) Contains(v Vertex) bool {
	return e.v1 == v || e.v2 == v
}

// Equal checks whether to edges represent the same Edge..
func (e Edge) Equal(other Edge) bool {
	ov1, ov2 := other.Vertices()[0], other.Vertices()[1]
	if e.v1 == ov1 && e.v2 == ov2 {
		return true
	}
	return e.v1 == ov2 && e.v2 == ov1
}

// Reversed returns a a version of the Edge that has its vertices reversed.
func (e Edge) Reversed() Edge {
	redge := Edge{e.v2, e.v1}
	return redge
}

// Vertices return the two vertices that make up the Edge.
func (e Edge) Vertices() [2]Vertex {
	return [2]Vertex{e.v1, e.v2}
}

// String return a string representation of the Edge.
func (e Edge) String() string {
	return fmt.Sprintf("Edge(%v, %v)", e.v1, e.v2)
}

// CullDuplicates is a helper method that takes a list of edges an removes all duplicate edges.
// BUG: The ordering of edges in the output slice is undefined.
func cullDuplicates(edges []Edge) (uniqueEdges []Edge) {
	edgeSet := make(map[Edge]bool, len(edges))

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
