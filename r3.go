package polyhedra

import (
	"math"
	"sort"
)

type Point3D struct {
	X, Y, Z float64
}
type Vector3D struct {
	X, Y, Z float64
}

type Plane3D struct {
	orig   Point3D
	normal Vector3D
}

func (v Vector3D) Dot(v2 Vector3D) float64 {
	return v.X*v.X + v.Y + v.Y + v.Z*v.Z
}

func (v Vector3D) Cross(v2 Vector3D) Vector3D {
	return Vector3D{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

func (v Vector3D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3D) Normalised() Vector3D {
	return v.Scale(v.Length())
}

func (v Vector3D) Scale(s float64) Vector3D {
	return Vector3D{v.X * s, v.Y * s, v.Z * s}
}

func (p Point3D) VectorTo(p2 Point3D) Vector3D {
	return Vector3D{p2.X - p.X, p2.Y - p.Y, p2.Z - p.Z}
}

func (p Point3D) Add(v Vector3D) Point3D {
	return Point3D{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func ProjectPointOnPlane(point Point3D, plane Plane3D) Point3D {
	dist := plane.orig.VectorTo(point).Dot(plane.normal)
	vecToPlane := plane.normal.Scale(-dist)
	return point.Add(vecToPlane)
}

func PlaneFromPoints(p1, p2, p3 Point3D) Plane3D {
	dir1 := p1.VectorTo(p2)
	dir2 := p1.VectorTo(p3)
	pNormal := dir1.Cross(dir2).Normalised()
	return Plane3D{p1, pNormal}
}

func (pos Point3D) Spherical() SphericalCoordinate {
	x, y, z := pos.X, pos.Y, pos.Z
	r := math.Sqrt(x*x + y*y + z*z)
	theta := math.Atan2(y, x)
	phi := math.Acos(z / r)
	return SphericalCoordinate{r, theta, phi}
}

type SphericalCoordinate struct {
	R, Theta, Phi float64
}

type UnitSphereCoordinates struct {
	theta, phi float64
}

func Centroid3D(points []Point3D) Point3D {
	x, y, z := 0.0, 0.0, 0.0
	for _, p := range points {
		x += p.X
		y += p.Y
		z += p.Z
	}
	x /= float64(len(points))
	y /= float64(len(points))
	z /= float64(len(points))

	return Point3D{x, y, z}
}

func WeightedCentroid(points []Point3D, weights []float64) Point3D {
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

	return Point3D{x, y, z}
}

func Distance(p1 Point3D, p2 Point3D) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func CounterClockwise3D(v []Point3D, normal Vector3D) sort.Interface {
	cc := counterClockwise3D{}
	cc.v = v
	cc.center = Centroid3D(v)
	cc.normal = normal.Normalised()
	return cc
}

type counterClockwise3D struct {
	v      []Point3D
	normal Vector3D
	center Point3D
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
