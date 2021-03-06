package polyhedra

import (
	"testing"
)

func TestGGSubdivision(t *testing.T) {

	checkSubdivision := func(m, n int) {
		gg := NewIcosahedralGeodesic()

		err := gg.Subdivide(m, n)
		if err != nil {
			t.Fatalf("Legal subdivision failed: %v", err)
		}
		T := n*m + n*n + m*m
		faceNum := len(gg.faces)
		if faceNum != T*20 {
			t.Errorf("Number of faces is %v instead of %v.", faceNum, 20*T)
		}
		edgeNum := len(gg.Edges())
		if edgeNum != T*30 {
			t.Errorf("Number of edges is %v instead of %v.", edgeNum, 30*T)
		}
		vertexNum := len(gg.vertices)
		if vertexNum != T*10+2 {
			t.Errorf("Number of vertices is %v instead of %v.", vertexNum, 10*T+3)
		}
		errs := IcosahedralGeodesicIntegrityChecker(*gg).CheckIntegrity()
		if len(errs) != 0 {
			t.Errorf("Subdivision (n=%v,m=%v) created illegal structure: %v", n, m, errs)
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
	err := IcosahedralGeodesicIntegrityChecker(*gg).CheckIntegrity()
	if err != nil && len(err) != 0 {
		t.Fatal(err)
	}

	n := 0
	m := 1
	for i := 2; i < 5; i++ {
		m = m * 2
		t.Logf("Testing (m=%v,n=%v) from n=%v", m, n, n/2)
		err := gg.Subdivide(2, n)
		if err != nil {
			t.Fatalf("Legal subdivision failed: %v", err)
		}
		errs := IcosahedralGeodesicIntegrityChecker(*gg).CheckIntegrity()
		if len(errs) != 0 {
			t.Fatalf("Subdivision (%v,%v) created illegal GG: %v", n, m, errs)
		}
	}
}

func TestGGEdgeOrder(t *testing.T) {
	gg1 := NewIcosahedralGeodesic()
	gg2 := NewIcosahedralGeodesic()

	e1S, e2S := gg1.Edges(), gg2.Edges()

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
