package r2

import "testing"

func countDistinctVertices(ts []Triangle) int {
	m := make(map[Vertex]Vertex)
	for _, t := range ts {
		for _, v := range t.vertices {
			m[v] = v
		}
	}
	return len(m)
}

func TestTriangleSubdivision(t *testing.T) {
	for i := 1; i < 99; i++ {
		tri := Triangle{[3]Vertex{
			NewVertex(),
			NewVertex(),
			NewVertex(),
		}}
		result, _ := SubdivideTriangle(tri, i, 0)
		if len(result) != (i * i) {
			t.Errorf("Subdivision with n %v resulted in T=%v instead of T=%v.", i, len(result), i*i)
		}

		expectedV := ((i + 1) * (1 + i + 1)) / 2
		actualV := countDistinctVertices(result)
		if expectedV != actualV {
			t.Errorf("Subdivision with n=%v resulted in %v unique vertices instead of %v.", i, actualV, expectedV)
		}
	}
	//n*(n+1)/2

}
