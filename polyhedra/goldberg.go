package polyhedra

// GoldbergPolyhedron represents a Polyhedron made of hexagons and pentagons.
// For more information see https://en.wikipedia.org/wiki/Goldberg_polyhedron
type GoldbergPolyhedron struct {
	Polyhedron
	m, n int
}

// GeodesicToGoldberg returns the goldberg Polyhedron that corresponds to the given geodesic Polyhedron.
// This is achieved by replacing all faces with vertices and adding edges between vertices that corresponded to neighbouring faces.
func GeodesicToGoldberg(g *Geodesic) (*GoldbergPolyhedron, error) {

	// For each Edge create a new vertex
	vertexMap := make(map[string]Vertex)
	newVertices := make([]Vertex, 0)
	for i := range g.faces {
		f := g.faces[i]
		v := NewVertex()
		vertexMap[f.String()] = v
		newVertices = append(newVertices, v)
		v.setPosition(f.Center())
	}

	// Turn adjacent vertices into faces
	newFaces := make([]Face, 0)
	for _, v := range g.vertices {
		af := g.VertexAdjacentFaces(v)
		loop := make([]Vertex, len(af))
		for i, f := range af {
			loop[i] = vertexMap[f.String()]
		}
		newFace := NewFace(SortedClockwise(loop))
		newFaces = append(newFaces, newFace)
	}

	newEdges := make([]Edge, 0)
	for _, e := range g.Edges() {
		fs := g.EdgeAdjacentFaces(e)
		v1 := vertexMap[fs[0].String()]
		v2 := vertexMap[fs[1].String()]
		newEdges = append(newEdges, NewEdge(v1, v2))
	}

	poly := GoldbergPolyhedron{}
	var err error
	polyBase, err := NewPolyhedron(newVertices, newEdges, newFaces)
	if err != nil {
		return nil, err
	}
	poly.Polyhedron = *polyBase
	poly.m = g.m
	poly.n = g.n

	return &poly, nil
}

// NewIcosahedralGoldbergPolyhedron creates a new GoldbergPolyhedron that has an icosahedron as a base and is subdivided
// according to the breakdown (n,m).
func NewIcosahedralGoldbergPolyhedron(m int, n int) (*GoldbergPolyhedron, error) {
	baseGeodesic := NewIcosahedralGeodesic()
	baseGeodesic.Subdivide(m, n)
	result, err := GeodesicToGoldberg(baseGeodesic)
	return result, err
}
