package polyhedra

import (
	"testing"
)

func TestIcosahedronCreation(t *testing.T) {
	ico := NewIcosahedron()

	for _, face := range ico.faces{
		for i := range face.Loop{
			p1 := face.Loop[i].Position()
			p2 := face.Loop[(i+1)%len(face.Loop)].Position()
			t.Logf("vec4(%v,%v,%v,1.0),", p1.X, p1.Y, p1.Z)
			t.Logf("vec4(%v,%v,%v,1.0),", p2.X, p2.Y, p2.Z)
		}

	}


	errors := IcosahedralGeodesicIntegrityChecker(IcosahedralGeodesic{Geodesic{ico, 1, 0}}).CheckIntegrity()
	if len(errors) != 0 {
		t.Fatalf("Geodesic is in illegal state: %v ", errors)
	}

	if len(ico.faces) != 20 {
		t.Errorf("Icosahedron has %v faces instead of 20", len(ico.faces))
	}

	if len(ico.faces) != 20 {
		t.Errorf("Icosahedron has %v faces instead of 20", len(ico.faces))
	}

	if len(ico.edges) != 30 {
		t.Errorf("Icosahedron has %v edges instead of 30", len(ico.edges))
	}

	if len(ico.vertices) != 12 {
		t.Errorf("Icosahedron has %v vertices instead of 12", len(ico.vertices))
	}

	for _, vertex := range ico.vertices {
		edge_count := 0
		for _, edge := range ico.edges {
			if edge.Contains(vertex) {
				edge_count++
			}
		}
		if edge_count != 5 {
			t.Errorf("Vertex %v has %v edges instaed of 5", vertex, edge_count)
		}
	}

	for _, edge := range ico.edges {
		face_count := 0
		if face_count != 2 {
			t.Errorf("Edge %v is contained in %v faces instaed of 2", edge, face_count)
		}
	}

}
