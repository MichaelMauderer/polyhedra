package polyhedra

import "github.com/MichaelMauderer/geometry/r3"

// NewFace creates a Face from the given list of vertices.
func NewFace(loop []Vertex) Face {
	f := Face{}
	f.loop = loop
	f.initEdges()
	f.initCenter()
	return f
}

type Face struct {
	loop   []Vertex
	edges  []Edge
	center r3.Point
}

// initEdges computes the edges between all consecutive vertices in the given list, as well as the last and first one.
// The result is stored in the Face.edges attribute.
func (f *Face) initEdges() {
	f.edges = make([]Edge, len(f.loop))
	for i := range f.loop {
		v1 := f.loop[i]
		v2 := f.loop[(i+1)%len(f.loop)]
		f.edges[i] = Edge{v1, v2}
	}
}

// initCenter precomputes the geometric center of all vertices and stores it in Face.center.
func (f *Face) initCenter() {
	f.center = vertexCentroid(f.loop)
}

// Loop returns the list of vertices that make up the Face.
func (f *Face) Loop() []Vertex {
	return f.loop
}

// Edges returns the list of edges that are part of the Face.
func (f *Face) Edges() []Edge {
	return f.edges
}

// Center returns the geometric mean of all vertices that make up this Face.
func (f *Face) Center() r3.Point {
	return f.center
}

func (f *Face) Equals(fo Face) bool {
	if len(f.loop) != len(fo.loop) {
		return false
	}
	for i := range f.loop {
		if f.loop[i] != fo.loop[i] {
			return false
		}
	}
	return true
}
