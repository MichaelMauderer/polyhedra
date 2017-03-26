package polyhedra

import (
	"testing"
	"math/rand"
	"sort"
)

func shuffled(a []Point3D) []Point3D {
	shuffled := make([]Point3D, len(a))
	perm := rand.Perm(len(a))
	for i, v := range perm {
		shuffled[v] = a[i]
	}
	return shuffled
}

func TestCounterClockwiseSorting(t *testing.T) {
	sorted := []Point3D{
		{0, 0, 1},
		{-1, 0, 2},
		{-1, 1, 3},
		{0, 1, 4},
	}

	normal := Vector3D{0, 0, 1}

	unsorted := shuffled(sorted)

	sort.Sort(CounterClockwise3D(unsorted, normal))

	for i := range sorted {
		if unsorted[i] != sorted[i] {
			t.Errorf("Sorting produced %v instead of %v", unsorted, sorted)
		}
	}
}
