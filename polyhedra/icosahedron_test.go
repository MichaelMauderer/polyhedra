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
