package polyhedra

import "testing"

func TestEdgeReversal(t *testing.T) {
	e1 := Edge{NewVertex(), NewVertex()}
	e2 := Edge{e1.v2, e1.v1}

	if e1.Reversed() != e2{
		t.Error("Edge.Reversed does not proiduce the reversed edge.")
	}
	if e1.Reversed() == e1{
		t.Error("Edge.Reversed equals itself.")
	}
	if e1.Reversed().Reversed() != e1{
		t.Error("Twice reversed edge not euqal itself.")
	}
}