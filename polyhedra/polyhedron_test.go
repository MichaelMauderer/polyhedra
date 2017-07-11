package polyhedra

import (
	"math/rand"
	"testing"

	"github.com/MichaelMauderer/geometry/r3"
)

func TestEdgeReversal(t *testing.T) {
	e1 := Edge{NewVertex(), NewVertex()}
	e2 := Edge{e1.v2, e1.v1}

	if e1.Reversed() != e2 {
		t.Error("Edge.Reversed does not proiduce the reversed Edge.")
	}
	if e1.Reversed() == e1 {
		t.Error("edge.Reversed equals itself.")
	}
	if e1.Reversed().Reversed() != e1 {
		t.Error("Twice reversed edge not euqal itself.")
	}
}

func shuffledVertices(a []Vertex) []Vertex {
	shuffled := make([]Vertex, len(a))
	perm := rand.Perm(len(a))
	for i, v := range perm {
		shuffled[v] = a[i]
	}
	return shuffled
}

func TestSortedClockwise(t *testing.T) {
	t.Skip("Skipping clockwise sorting test.")

	sortedPositions := []r3.Point{
		{2, 2, 1},
		{2, 1, 2},
		{1, 1, 3},
		{1, 2, 4},
	}
	vertices := []Vertex{
		NewVertex(),
		NewVertex(),
		NewVertex(),
		NewVertex(),
	}
	for i := range vertices {
		vertices[i].setPosition(sortedPositions[i])
	}

	for i := 0; i < 99; i++ {
		shuffledPart := shuffledVertices(vertices[1:])

		shuffled := make([]Vertex, 1)
		shuffled[0] = vertices[0]
		for _, item := range shuffledPart {
			shuffled = append(shuffled, item)
		}

		resorted := SortedClockwise(shuffled)
		for i := range vertices {
			if vertices[i] != resorted[i] {
				t.Errorf("Expected %v but got %v", vertices, resorted)
			}
		}
	}

}
