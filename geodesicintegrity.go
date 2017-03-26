package polyhedra

import (
	"math"
	"errors"
	"fmt"
	"log"
)

type IcosahedriGeodesicIntegrityChecker IcosahedralGeodesic

func (gic IcosahedriGeodesicIntegrityChecker) checkFaces() error {
	faceNum := len(gic.faces)
	if faceNum%20 != 0 {
		return errors.New("Number of faces is not a multiple of 20.")
	}
	return nil
}

func (gic IcosahedriGeodesicIntegrityChecker) checkEdges() error {
	edgeNum := len(gic.edges)
	if edgeNum%30 != 0 {
		return errors.New("Number of faces is not a multiple of 30.")
	}
	for _, edge := range gic.edges {
		if edge.v2 == edge.v1 {
			return errors.New("Edges contain illegal self-loops.")
		}
		zero := Point3D{0.0, 0.0, 0.0}
		if edge.Center() == zero {
			return errors.New(fmt.Sprintf("Contains edge %v centered at zero with vertices %v to %v", edge, edge.v1.String(), edge.v2.String()))
		}

	}
	return nil
}

func (gic IcosahedriGeodesicIntegrityChecker) checkVertexNum() error {
	vertexNum := len(gic.vertices)
	if (vertexNum-2)%10 != 0 {
		return errors.New("Number of vertices does not fulfill V=(T*10+2)")
	}
	return nil
}

func (gic IcosahedriGeodesicIntegrityChecker) checkVertexDegrees() error {
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

func (gic IcosahedriGeodesicIntegrityChecker) checkDistinctVertexNeighbors() error {
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

func (gic IcosahedriGeodesicIntegrityChecker) checkVertexDistances() error {
	bp1 := gic.edges[0].v1.Position()
	bp2 := gic.edges[0].v2.Position()
	baseLineDistance := bp1.VectorTo(bp2).Length()
	epsilon := baseLineDistance * 0.1
	for _, edge := range gic.edges {
		p1 := edge.v1.Position()
		p2 := edge.v2.Position()
		dist := p1.VectorTo(p2).Length()
		delta := math.Abs(dist - baseLineDistance)
		if delta > epsilon {
			return errors.New(fmt.Sprintf("Vertices %v and %v vary in distance too much: %v with a baseline of %v", p1, p2, delta, baseLineDistance))
		}
	}
	return nil
}

func (gic IcosahedriGeodesicIntegrityChecker) CheckIntegrity() []error {

	var checks = []func() error{
		gic.checkFaces,
		gic.checkVertexDegrees,
		gic.checkEdges,
		gic.checkVertexNum,
		gic.checkVertexDistances,
		gic.checkDistinctVertexNeighbors,
	}
	errors := make([]error, 0, len(checks))
	for _, check := range checks {
		err := check()
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
