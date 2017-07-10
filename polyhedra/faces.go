package polyhedra

import "github.com/MichaelMauderer/geometry/r3"

// Face represents a surface face in a polyhedron.
type Face interface {
	Loop() []Vertex
	Edges() []Edge
	Center() r3.Point3D
}

// NewFace creates a face from the given list of vertices.
func NewFace(loop []Vertex) Face {
	f := face{}
	f.loop = loop
	f.initEdges()
	f.initCenter()
	return &f
}

type face struct {
	loop   []Vertex
	edges  []Edge
	center r3.Point3D
}

// initEdges computes the edges between all consecutive vertices in the given list, as well as the last and first one.
// The result is stored in the face.edges attribute.
func (f *face) initEdges() {
	f.edges = make([]Edge, len(f.loop))
	for i := range f.loop {
		v1 := f.loop[i]
		v2 := f.loop[(i+1)%len(f.loop)]
		f.edges[i] = edge{v1, v2}
	}
}

// initCenter precomputes the geometric center of all vertices and stores it in face.center.
func (f *face) initCenter() {
	f.center = vertexCentroid(f.loop)
}

// Loop returns the list of vertices that make up the face.
func (f *face) Loop() []Vertex {
	return f.loop
}

// Edges returns the list of edges that are part of the face.
func (f *face) Edges() []Edge {
	return f.edges
}

// Center returns the geometric mean of all vertices that make up this face.
func (f *face) Center() r3.Point3D {
	return f.center
}
