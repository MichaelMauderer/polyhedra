package polyhedra

import "github.com/MichaelMauderer/geometry/r3"

type Face interface {
	Loop() []Vertex
	Edges() []Edge
	Center() r3.Point3D
}

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

func (f *face) initEdges() {
	f.edges = make([]Edge, len(f.loop))
	for i := range f.loop {
		v1 := f.loop[i]
		v2 := f.loop[(i+1)%len(f.loop)]
		f.edges[i] = Edge{v1, v2}
	}
}

func (f *face) initCenter() {
	f.center = vertexCentroid(f.loop)
}

func (f *face) Loop() []Vertex {
	return f.loop
}

func (f *face) Edges() []Edge {
	return f.edges
}

func (f *face) Center() r3.Point3D {
	return f.center
}
