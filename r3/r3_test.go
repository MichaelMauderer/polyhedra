package r3

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

func shuffled(a []Point) []Point {
	shuffled := make([]Point, len(a))
	perm := rand.Perm(len(a))
	for i, v := range perm {
		shuffled[v] = a[i]
	}
	return shuffled
}

func TestCounterClockwiseSorting(t *testing.T) {
	sorted := []Point{
		{0, 0, 1},
		{-1, 0, 2},
		{-1, 1, 3},
		{0, 1, 4},
	}

	normal := Vector{0, 0, 1}

	unsorted := shuffled(sorted)

	sort.Sort(CounterClockwise3D(unsorted, normal))

	for i := range sorted {
		if unsorted[i] != sorted[i] {
			t.Errorf("Sorting produced %v instead of %v", unsorted, sorted)
		}
	}
}

var epsilon = 10E-07

func isClose(f1, f2 float64) bool {
	return math.Abs(f1-f2) < epsilon
}

func assertFloatClose(fExpected, fActual float64, t *testing.T) {
	if !isClose(fExpected, fActual) {
		t.Errorf("Expected %v but got %v", fExpected, fActual)
	}
}

func assertSphericalClose(sExpected, sActual SphericalCoordinate, t *testing.T) {
	assertFloatClose(sExpected.Theta, sActual.Theta, t)
	assertFloatClose(sExpected.Phi, sActual.Phi, t)
	assertFloatClose(sExpected.R, sActual.R, t)
}

func TestSphericalConversion(t *testing.T) {
	p := Point{0, 0, 1}
	s := p.Spherical()
	if s.R != 1 {
		t.Errorf("Expected radius %v but got %v", 1, s.R)
	}
	if s.Theta != 0 {
		t.Errorf("Expected theta %v but got %v", 0, s.Theta)
	}

	p = Point{1, 1, 1}
	s = p.Spherical()
	expected := SphericalCoordinate{
		math.Sqrt(3),
		math.Asin(math.Sqrt(2.0 / 3.0)),
		math.Pi / 4,
	}
	assertSphericalClose(expected, s, t)

	p = Point{2, 2, 2}
	s = p.Spherical()
	expected = SphericalCoordinate{
		math.Sqrt(12),
		math.Asin(math.Sqrt(2.0 / 3.0)),
		math.Pi / 4,
	}
	assertSphericalClose(expected, s, t)

	p = Point{1, 0, 0}
	s = p.Spherical()
	expected = SphericalCoordinate{
		1,
		math.Pi / 2,
		0,
	}
	assertSphericalClose(expected, s, t)

	p = Point{1, 0, -1}
	s = p.Spherical()
	expected = SphericalCoordinate{
		math.Sqrt2,
		3 * math.Pi / 4,
		0,
	}
	assertSphericalClose(expected, s, t)
}

func assertIsCCW(p1, p2, c Point, n Vector, t *testing.T) {
	if !IsCCW(p1, p2, c, n) {
		t.Errorf("Expetced %v to be counter clock wise of %v", p1, p2)
	}
}

func assertIsNotCCW(p1, p2, c Point, n Vector, t *testing.T) {
	if IsCCW(p1, p2, c, n) {
		t.Errorf("Expetced %v not to be counter clock wise of %v", p1, p2)
	}
}

func TestIsCCW(t *testing.T) {
	cwSortedPoints := []Point{
		{0, 0, 1},
		{-1, 0, 2},
		{-1, 1, 3},
		{2, 2, 4},
	}
	normal := Vector{0, 0, 1}
	center := Centroid3D(cwSortedPoints)

	cwPairs := [][2]int{
		{1, 0},
		{2, 1},
		{3, 2},
		{2, 0},
		{1, 3},
	}

	for _, ab := range cwPairs {
		i, j := ab[0], ab[1]
		assertIsCCW(cwSortedPoints[i], cwSortedPoints[j], center, normal, t)
		assertIsNotCCW(cwSortedPoints[j], cwSortedPoints[i], center, normal, t)
	}

}
