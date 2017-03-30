package polyhedra

import "testing"

func assertFaceCount(p Polyhedron, fn int, t *testing.T) {
	faces := len(p.Faces())
	if faces != fn {
		t.Errorf("Wrong number of faces: %v instead of %v", faces, fn)
	}
}

func assertEdgeCount(p Polyhedron, en int, t *testing.T) {
	edges := len(p.Edges())
	if edges != en {
		t.Errorf("Wrong number of edges: %v instead of %v", edges, en)
	}
}

func assertVertexCount(p Polyhedron, vn int, t *testing.T) {
	vertices := len(p.Vertices())
	if vertices != vn {
		t.Errorf("Wrong number of vertices: %v instead of %v", vertices, vn)
	}
}

func assertVertexDegrees(p Polyhedron, t *testing.T) {
	vertices := p.Vertices()
	for _, v := range vertices{
		vd := p.VertexDegree(v)
		if vd != 3 {
			t.Errorf("Vertex degree for vertex %v should be 3 but is %v", v, vd)
		}
	}
}

func assertVertexAdjacentFaceCount(p Polyhedron, t *testing.T) {
	vertices := p.Vertices()
	for _, v := range vertices{
		vd := len(p.VertexAdjacentFaces(v))
		if vd != 3 {
			t.Errorf("Vertex %v should have 3 adjacent faces but has %v", v, vd)
		}
	}
}

func TestNewIcosahedralGoldbergPolyhedronCreation(t *testing.T) {

	m, n := 1, 0
	igp, _ := NewIcosahedralGoldbergPolyhedron(m, n)
	T := m*n + m*m + n*n
	expectedFaces := 10*T + 2
	expectedVertices := 20 * T
	expectedEdges := 30 * T

	assertFaceCount(igp, expectedFaces, t)
	assertVertexCount(igp, expectedVertices, t)
	assertEdgeCount(igp, expectedEdges, t)
	assertVertexDegrees(igp,t)
}
