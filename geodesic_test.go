package polyhedra

import (
	"testing"
)

func TestGGSubdivision(t *testing.T) {

	checkSubdivision := func(m, n int) {
		gg := NewIcosahedralGeodesic()
		err := gg.subdivide(m, n)
		if err != nil {
			t.Fatalf("Legal subdivision failed: %v", err)
		}
		T := n*m + n*n + m*m
		faceNum := len(gg.faces)
		if faceNum != T*20 {
			t.Errorf("Number of faces is %v instead of %v.", faceNum, 20*T)
		}
		edgeNum := len(gg.edges)
		if edgeNum != T*30 {
			t.Errorf("Number of edges is %v instead of %v.", edgeNum, 30*T)
		}
		vertexNum := len(gg.vertices)
		if vertexNum != T*10+2 {
			t.Errorf("Number of vertices is %v instead of %v.", vertexNum, 10*T+3)
		}
		err = gg.CheckIntegrity()
		if err != nil {
			t.Errorf("Subdivision (n=%v,m=%v) created illegal structure: %v", n, m, err)
		}
	}
	n := 0
	for m := 2; m < 3; m++ {
		t.Logf("Testing (m=%v,n=%v)", m, n)
		checkSubdivision(m, n)
	}
}

func TestGGRepeatedSubdivision(t *testing.T) {

	gg := NewIcosahedralGeodesic()
	err := gg.CheckIntegrity()
	if err != nil {
		t.Fatal(err)
	}

	n := 0
	m := 1
	for i := 2; i < 5; i++ {
		m = m * 2
		t.Logf("Testing (m=%v,n=%v) from n=%v", m, n, n/2)
		err := gg.subdivide(2, n)
		if err != nil {
			t.Fatalf("Legal subdivision failed: %v", err)
		}
		err = gg.CheckIntegrity()
		if err != nil {
			t.Fatalf("Subdivision (%v,%v) created illegal GG: %v", n, m, err)
		}
	}
}
