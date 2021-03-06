// Package r3 implements some utility functionality related to three-dimensional geometry.
package r3

import (
	"math"
	"sort"
)

// Point represent a point in 3D space through cartesian coordinates.
type Point struct {
	X, Y, Z float64
}

// Vector represent a vector in 3D space .
type Vector struct {
	X, Y, Z float64
}

// Plane represent a plane in 3D space .
type Plane struct {
	orig   Point
	normal Vector
}

// Dot computes the dot product between this and the given vector.
func (v Vector) Dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// Cross computes the cross product between this and the given vector.
func (v Vector) Cross(v2 Vector) Vector {
	return Vector{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

// Length returns the length of the vector.
func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalised returns a copy of this vector that is scaled to length 1.
func (v Vector) Normalised() Vector {
	return v.Scale(1 / v.Length())
}

// Scale returns a copy of this vector that has its length multiplied by the given value.
func (v Vector) Scale(s float64) Vector {
	return Vector{v.X * s, v.Y * s, v.Z * s}
}

// VectorTo returns the vector from this point to the given point.
func (p Point) VectorTo(p2 Point) Vector {
	return Vector{p2.X - p.X, p2.Y - p.Y, p2.Z - p.Z}
}

// Add returns the point that results by displacing this point by the given vector.
func (p Point) Add(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

// ProjectPointOnPlane returns the point in the given plane  that is closes to the given point.
func ProjectPointOnPlane(point Point, plane Plane) Point {
	dist := plane.orig.VectorTo(point).Dot(plane.normal)
	vecToPlane := plane.normal.Scale(-dist)
	return point.Add(vecToPlane)
}

// PlaneFromPoints returns the plane defined by the given three points.
func PlaneFromPoints(p1, p2, p3 Point) Plane {
	dir1 := p1.VectorTo(p2)
	dir2 := p1.VectorTo(p3)
	pNormal := dir1.Cross(dir2).Normalised()
	return Plane{p1, pNormal}
}

// Spherical returns the spherical coordinates of the given point.
func (p Point) Spherical() SphericalCoordinate {
	x, y, z := p.X, p.Y, p.Z
	r := math.Sqrt(x*x + y*y + z*z)
	theta := math.Acos(z / r)
	phi := math.Atan2(y, x)
	return SphericalCoordinate{r, theta, phi}
}

// SphericalCoordinate represents a point in 3D space using spherical coordinates.
type SphericalCoordinate struct {
	R, Theta, Phi float64
}

// UnitSphereCoordinates represents a point in 3D space on the unit sphere.
type UnitSphereCoordinates struct {
	theta, phi float64
}

// Centroid3D computes the centroid of the given points.
func Centroid3D(points []Point) Point {
	x, y, z := 0.0, 0.0, 0.0
	for _, p := range points {
		x += p.X
		y += p.Y
		z += p.Z
	}
	x /= float64(len(points))
	y /= float64(len(points))
	z /= float64(len(points))

	return Point{x, y, z}
}

// WeightedCentroid computes the centroid of the given points, but the contribution of each point is scaled by its weight.
// The weights represent relative weights and will be scaled by the overall sum of weights.
func WeightedCentroid(points []Point, weights []float64) Point {
	x, y, z := 0.0, 0.0, 0.0
	wSum := 0.0
	for i, p := range points {
		w := weights[i]
		wSum += w
		x += p.X * w
		y += p.Y * w
		z += p.Z * w
	}
	x /= wSum
	y /= wSum
	z /= wSum

	return Point{x, y, z}
}

// Distance computes the distance between the two given points.
func Distance(p1 Point, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// CounterClockwise3D sorts the given points clockwise around the given normal.
func CounterClockwise3D(v []Point, normal Vector) sort.Interface {
	cc := counterClockwise3D{}
	cc.v = v
	cc.center = Centroid3D(v)
	cc.normal = normal.Normalised()
	return cc
}

type counterClockwise3D struct {
	v      []Point
	normal Vector
	center Point
}

func (v counterClockwise3D) Len() int      { return len(v.v) }
func (v counterClockwise3D) Swap(i, j int) { v.v[i], v.v[j] = v.v[j], v.v[i] }
func (v counterClockwise3D) Less(i, j int) bool {
	v1 := v.center.VectorTo(v.v[i])
	v2 := v.center.VectorTo(v.v[j])
	vc := v1.Cross(v2)
	n := v.normal.Dot(vc)
	return n < 0
}

// IsCCW returns whether point a is clockwise to point b relative to the given normal.
func IsCCW(a, b Point, center Point, normal Vector) bool {
	v1 := center.VectorTo(a)
	v2 := center.VectorTo(b)
	vc := v1.Cross(v2)
	n := normal.Dot(vc)
	return n > 0
}
