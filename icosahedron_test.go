package polyhedra

import "testing"

func TestIcosahedronCreation(t *testing.T) {
	ico := NewIcosahedron()

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
		for _, face := range ico.faces {
			for _, face_edge := range face.edges {
				if face_edge.Equal(edge) {
					face_count++
				}
			}

		}
		if face_count != 2 {
			t.Errorf("Edge %v is contained in %v faces instaed of 2", edge, face_count)
		}
	}

	for _, face := range ico.faces {
		edgeCount := len(face.edges)
		if edgeCount != 3 {
			t.Errorf("Face %v has %v edges instead of 3", face, edgeCount)
		}

	}
}
