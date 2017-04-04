package polyhedra

import "github.com/MichaelMauderer/geometry/r3"

type Edge struct {
	v1, v2 Vertex
}

func (e Edge) Length() float64 {
	return r3.Distance(e.v1.Position(), e.v2.Position())
}

func (e Edge) Center() r3.Point3D {
	return r3.Centroid3D([]r3.Point3D{e.v1.Position(), e.v2.Position()})
}

func (e Edge) Contains(v Vertex) bool {
	return e.v1 == v || e.v2 == v
}

func (e Edge) Equal(other Edge) bool {
	if e.v1 == other.v1 && e.v2 == other.v2 {
		return true
	}
	return e.v1 == other.v2 && e.v2 == other.v1
}
func (e Edge) Reversed() Edge {
	return Edge{e.v2, e.v1}
}
