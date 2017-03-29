package polyhedra

import (
	"github.com/MichaelMauderer/polyhedra"
	"errors"
)


func GedoesicToGoldber(g *Geodesic) polyhedra.Interface{

	// For each edge create a new vertex
	verticeMap := make(map[Edge]Vertex)
	for _, e := range g.edges{
		v := NewVertex()
		verticeMap[e] = v
		v.setPosition(e.Center())
	}

	faceMap := make(map[Vertex]Face)
	for _, v := range g.vertices{
		nv := g.AdjacentVertices(v)
		// TODO Sort clockwise?
		faceMap[v] = Face{nv}
	}

	edgeSet := make(map[Edge]bool)
	for _, f := range(faceMap){

		for _, e := range f.Edges(){
			hasE := edgeSet[e]
			hasRE := edgeSet[e.Reversed()]
			if !(hasE && hasRE){
				edgeSet[e] = true
			}
		}
	}

	poly := Polyhedron{}

	poly.faces = make([]Face, 0, len(faceMap))
	for _, f := range faceMap{
		poly.faces = append(poly.faces, f)
	}

	poly.edges = make([]Edge, 0, len(edgeSet))
	for e, _:= range edgeSet{
		poly.edges = append(poly.edges, e)
	}

	poly.vertices = make([]Vertex, 0, len(verticeMap))
	for _, v:= range verticeMap{
		poly.vertices = append(poly.vertices, v)
	}

	return &poly
}

func NewGoldberPolyhedra(m int, n int) (polyhedra.Interface, error) {
	if n != m {
		return nil, errors.New("Class III not supported.")
	}
	if m != 0 {
		return nil, errors.New("Class II not supported.")
	}

	baseGeodesic := NewIcosahedralGeodesic()
	baseGeodesic.Subdivide(m ,n)

	return nil, nil
}
