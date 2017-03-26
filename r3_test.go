package polyhedra

import (
	"testing"
	"math/rand"
	"sort"
	"math"
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
	p := Point3D{0, 0, 1}
	s := p.Spherical()
	if s.R != 1 {
		t.Errorf("Expected radius %v but got %v", 1, s.R)
	}
	if s.Theta != 0 {
		t.Errorf("Expected theta %v but got %v", 0, s.Theta)
	}

	p = Point3D{1, 1, 1}
	s = p.Spherical()
	expected := SphericalCoordinate{
		math.Sqrt(3),
		math.Asin(math.Sqrt(2.0/3.0)),
		math.Pi / 4,
	}
	assertSphericalClose(expected, s, t)

	p = Point3D{2, 2, 2}
	s = p.Spherical()
	expected = SphericalCoordinate{
		math.Sqrt(12),
		math.Asin(math.Sqrt(2.0/3.0)),
		math.Pi / 4,
	}
	assertSphericalClose(expected, s, t)


	p = Point3D{1, 0, 0}
	s = p.Spherical()
	expected = SphericalCoordinate{
		1,
		math.Pi/2,
		0,
	}
	assertSphericalClose(expected, s, t)

	p = Point3D{1, 0, -1}
	s = p.Spherical()
	expected = SphericalCoordinate{
		math.Sqrt2,
		3*math.Pi/4,
		0,
	}
	assertSphericalClose(expected, s, t)
}
