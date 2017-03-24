package polyhedra

import (
	"testing"
	"log"
)

func TestGGSubdivision(t *testing.T) {
	poly := NewIcosahedron()

	checkSubdivision := func(n, m int) {
		var gg IcoGG = IcoGG{GeodesicGrid{poly}};
		err := gg.subdivide(n, m)
		if err != nil {
			t.Fatalf("Legal subdivision failed: %v", err)
		}
		err = gg.checkIntegrity()
		if err != nil {
			t.Fatalf("Subdivision (%v,%v) created illegal GG: %v", n, m, err)
		}
	}
	m := 0
	for n := 2; n < 99; n++ {
		log.Printf("Testing (m=%v,n=%v)", m, n)
		checkSubdivision(n, m)
	}
}

func TestGGRepeatedSubdivision(t *testing.T) {


	poly := NewIcosahedron()
	gg := IcoGG{GeodesicGrid{poly}}
	m := 0
	n := 1
	for i := 2; i < 10; i++ {
		n = n * 2
		log.Printf("Testing (m=%v,n=%v) from n=%v", m, n, n/2)
		err := gg.subdivide(2, m)
		if err != nil {
			t.Fatalf("Legal subdivision failed: %v", err)
		}
		err = gg.checkIntegrity()
		if err != nil {
			t.Fatalf("Subdivision (%v,%v) created illegal GG: %v", n, m, err)
		}
	}
}
