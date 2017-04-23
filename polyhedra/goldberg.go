package polyhedra

type GoldbergPolyhedron struct {
	Polyhedron
	m, n int
}

func GeodesicToGoldberg(g *Geodesic) *GoldbergPolyhedron {

	// For each edge create a new vertex
	vertexMap := make(map[Face]Vertex)
	newVertices := make([]Vertex, 0)
	for i, _ := range g.faces {
		f := g.faces[i]
		v := NewVertex()
		vertexMap[f] = v
		newVertices = append(newVertices, v)
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
		newFace := NewFace(SortedClockwise(loop))
		newFaces = append(newFaces, newFace)
	}

	newEdges := make([]Edge, 0)
	for _, e := range g.Edges() {
		fs := g.EdgeAdjacentFaces(e)
		v1 := vertexMap[fs[0]]
		v2 := vertexMap[fs[1]]
		newEdges = append(newEdges, edge{v1, v2})
	}

	poly := GoldbergPolyhedron{}
	poly.Polyhedron = NewPolyhedron(newVertices, newEdges, newFaces, )

	poly.m = g.m
	poly.n = g.n

	return &poly
}

func NewIcosahedralGoldbergPolyhedron(m int, n int) (*GoldbergPolyhedron, error) {
	baseGeodesic := NewIcosahedralGeodesic()
	baseGeodesic.Subdivide(m, n)
	result := GeodesicToGoldberg(baseGeodesic)
	return result, nil
}
