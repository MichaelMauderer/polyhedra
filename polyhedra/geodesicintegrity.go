package polyhedra

import (
	"math"
	"errors"
	"fmt"
	"log"
	"github.com/MichaelMauderer/geometry/r3"
)

type IcosahedralGeodesicIntegrityChecker IcosahedralGeodesic

func (gic IcosahedralGeodesicIntegrityChecker) checkFaces() error {
	faceNum := len(gic.faces)
	if faceNum%20 != 0 {
		return errors.New("Number of faces is not a multiple of 20.")
	}
	return nil
}

func (gic IcosahedralGeodesicIntegrityChecker) checkEdges() error {
	edgeNum := len(gic.edges)
	if edgeNum%30 != 0 {
		return errors.New("Number of faces is not a multiple of 30.")
	}
	for _, edge := range gic.edges {
		if edge.v2 == edge.v1 {
			return errors.New("Edges contain illegal self-loops.")
		}
		zero := r3.Point3D{0.0, 0.0, 0.0}
		if edge.Center() == zero {
			return errors.New(fmt.Sprintf("Contains edge %v centered at zero with vertices %v to %v", edge, edge.v1.String(), edge.v2.String()))
		}

	}
	return nil
}

func (gic IcosahedralGeodesicIntegrityChecker) checkVertexNum() error {
	vertexNum := len(gic.vertices)
	if (vertexNum-2)%10 != 0 {
		return errors.New("Number of vertices does not fulfill V=(T*10+2)")
	}
	return nil
}

func (gic IcosahedralGeodesicIntegrityChecker) checkVertexDegrees() error {
	foundWrongOne := false
	for _, vertex := range gic.vertices {
		if vertex == 0 {
			return errors.New(fmt.Sprintf("Contains illegal zero vertex."))
		}
		vD := gic.VertexDegree(vertex)
		if (vD != 5) && (vD != 6) {
			log.Printf("Vertex %v in %v has degree %v", vertex, gic, vD)
			foundWrongOne = true
		}
	}
	if foundWrongOne {
		return errors.New(fmt.Sprintf("Found invalid number of edges at vertex. Should be 5 or 6"))
	}
	return nil
}

func (gic IcosahedralGeodesicIntegrityChecker) checkDistinctVertexNeighbors() error {
	for _, vertex := range gic.vertices {
		neighbors := gic.AdjacentVertices(vertex)
		counts := make(map[Vertex]int)
		for _, oV := range neighbors {
			if oV == vertex {
				return errors.New(fmt.Sprintf("Vertex %v is its own neigbor", vertex))
			}
			counts[oV] += 1
		}
		for v, c := range counts {
			if c > 1 {
				return errors.New(fmt.Sprintf("Vertex %v is %v more than once (%v) as neighbor", vertex, v, c))
			}
		}
	}
	return nil
}

func (gic IcosahedralGeodesicIntegrityChecker) checkVertexDistances() error {
	baseLineDistance := gic.edges[0].Length()
	epsilon := 0.2
	for _, edge := range gic.edges {
		dist := edge.Length()
		delta := math.Abs(dist - baseLineDistance)
		if delta > epsilon {
			return errors.New(fmt.Sprintf("Edge %v deviates in length too much: %v with a baseline of %v", edge, delta, baseLineDistance))
		}
	}
	return nil
}

func (gic IcosahedralGeodesicIntegrityChecker) checkCenter() error {
	vertices := gic.vertices
	positions := make([]r3.Point3D, len(vertices))
	for i := range vertices{
		positions[i] = vertices[i].Position()
	}
	center := r3.Centroid3D(positions)
	epsilon := 0.000001
	if r3.Distance(center, r3.Point3D{0,0,0}) > epsilon{
		return errors.New(fmt.Sprintf("Center has mvoed from origin to %v", center))

	}
	return nil
}


func (gic IcosahedralGeodesicIntegrityChecker) CheckIntegrity() []error {

	var checks = []func() error{
		gic.checkFaces,
		gic.checkVertexDegrees,
		gic.checkEdges,
		gic.checkVertexNum,
		gic.checkVertexDistances,
		gic.checkDistinctVertexNeighbors,
		gic.checkCenter,
	}
	errs := make([]error, 0, len(checks))
	for _, check := range checks {
		err := check()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
