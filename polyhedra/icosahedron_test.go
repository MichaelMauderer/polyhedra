package polyhedra

import (
	"testing"
)

func TestIcosahedronCreation(t *testing.T) {
	ico := newIcosahedron()

	errors := IcosahedralGeodesicIntegrityChecker(IcosahedralGeodesic(Geodesic{ico, 1, 0})).CheckIntegrity()
	if len(errors) != 0 {
		t.Fatalf("Geodesic is in illegal state: %v ", errors)
	}

	if len(ico.faces) != 20 {
		t.Errorf("Icosahedron has %v faces instead of 20", len(ico.Faces()))
	}

	if len(ico.faces) != 20 {
		t.Errorf("Icosahedron has %v faces instead of 20", len(ico.Faces()))
	}

	if len(ico.Edges()) != 30 {
		t.Errorf("Icosahedron has %v edges instead of 30", len(ico.Edges()))
	}

	if len(ico.vertices) != 12 {
		t.Errorf("Icosahedron has %v vertices instead of 12", len(ico.vertices))
	}

	for _, vertex := range ico.vertices {
		edge_count := 0
		for _, edge := range ico.Edges() {
			if edge.Contains(vertex) {
				edge_count++
			}
		}
		if edge_count != 5 {
			t.Errorf("Vertex %v has %v edges instaed of 5", vertex, edge_count)
		}
	}
}

func TestIcosahedronEdgeOrder(t *testing.T) {
	ico1 := NewIcosahedron()
	ico2 := NewIcosahedron()

	e1S, e2S := ico1.Edges(), ico2.Edges()

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
