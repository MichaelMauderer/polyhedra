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
	for _, v := range vertices {
		vd := p.VertexDegree(v)
		if vd != 3 {
			t.Errorf("Vertex degree for vertex %v should be 3 but is %v", v, vd)
		}
	}
}

func assertVertexAdjacentFaceCount(p Polyhedron, t *testing.T) {
	vertices := p.Vertices()
	for _, v := range vertices {
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
	assertVertexDegrees(igp, t)
}

func TestVertexOrder(t *testing.T) {
	igp1, _ := NewIcosahedralGoldbergPolyhedron(1, 0)
	igp2, _ := NewIcosahedralGoldbergPolyhedron(1, 0)

	v1, v2 := igp1.Vertices(), igp2.Vertices()

	epsilon := 0.01
	for i := range v1 {
		d := v1[i].Position().VectorTo(v2[i].Position()).Length()
		if d > epsilon {
			t.Errorf("Expected vertex order to be the same but vertex %v is %v and %v.", i, v1[i], v2[i])

		}
	}
}

func TestEdgeOrder(t *testing.T) {
	igp1, _ := NewIcosahedralGoldbergPolyhedron(1, 0)
	igp2, _ := NewIcosahedralGoldbergPolyhedron(1, 0)

	e1S, e2S := igp1.Edges(), igp2.Edges()

	epsilon := 0.01
	for i := range e1S {
		e1 := e1S[i]
		e2 := e2S[i]

		d := e1.Center().VectorTo(e2.Center()).Length()
		if d > epsilon {
			t.Log(d)
			t.Errorf("Expected egde order to be the same but vertex %v is %v and %v.", i, e1, e2)

		}
	}
}

func TestFaceOrder(t *testing.T) {
	igp1, _ := NewIcosahedralGoldbergPolyhedron(1, 0)
	igp2, _ := NewIcosahedralGoldbergPolyhedron(1, 0)

	f1S, f2S := igp1.Faces(), igp2.Faces()

	epsilon := 0.01
	for i := range f1S {
		f1 := f1S[i]
		f2 := f2S[i]

		d := f1.Center().VectorTo(f2.Center()).Length()
		if d > epsilon {
			t.Log(d)
			t.Errorf("Expected face order to be the same but vertex %v is %v and %v.", i, f1, f2)

		}
	}
}
