package polyhedra

import ()

type GoldbergPolyhedron struct {
	polyhedron
	m, n int
}

func GedoesicToGoldberg(g *Geodesic) *GoldbergPolyhedron {

	// For each edge create a new vertex
	vertexMap := make(map[*Face]Vertex)
	for i, _ := range g.faces {
		f := &g.faces[i]
		v := NewVertex()
		vertexMap[f] = v
		v.setPosition(f.Center())
	}

	// Turn adjacent vertices into faces into edges
	newFaces := make([]Face, 0)
	for _, v := range g.vertices {
		af := g.VertexAdjacentFaces(v)
		loop := make([]Vertex, len(af))
		for i, f := range af {
			loop[i] = vertexMap[f]
		}
		newFace := Face{SortedClockwise(loop)}
		newFaces = append(newFaces, newFace)
	}

	newEdges := make([]Edge, 0)
	for _, e := range g.edges {
		fs := g.EdgeAdjacentFaces(e)
		v1 := vertexMap[fs[0]]
		v2 := vertexMap[fs[1]]
		newEdges = append(newEdges, Edge{v1, v2})
	}

	poly := GoldbergPolyhedron{}

	poly.faces = newFaces

	poly.edges = newEdges

	poly.vertices = make([]Vertex, 0, len(vertexMap))
	for _, v := range vertexMap {
		poly.vertices = append(poly.vertices, v)
	}

	poly.m = g.m
	poly.n = g.n

	return &poly
}

func NewIcosahedralGoldbergPolyhedron(m int, n int) (*GoldbergPolyhedron, error) {
	baseGeodesic := NewIcosahedralGeodesic()
	baseGeodesic.Subdivide(m, n)
	result := GedoesicToGoldberg(baseGeodesic)
	return result, nil
}
