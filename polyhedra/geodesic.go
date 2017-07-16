package polyhedra

import (
	"errors"
	"github.com/MichaelMauderer/polyhedra/r3"
)

// Geodesic represents a geodesic Polyhedron.
// More information on geodesic polyhedra can be found at https://en.wikipedia.org/wiki/Geodesic_polyhedron
type Geodesic struct {
	Polyhedron
	m, n int
}

// IcosahedralGeodesic represents a geodesic Polyhedron with an icosahedron as a base.
type IcosahedralGeodesic Geodesic

// NewIcosahedralGeodesic creates a geodesic Polyhedron from an icosahedron through subdivision.
func NewIcosahedralGeodesic() *Geodesic {
	ico := newIcosahedron()
	geo := Geodesic{ico, 1, 0}
	return &geo
}

func createVerticesForEdges(gg *Geodesic, m int, vertexToEdgeMap map[Edge]([]Vertex)) {
	for _, e := range gg.Edges() {
		nV := make([]Vertex, m-1)
		for j := range nV {
			nV[j] = NewVertex()
			ev := e.Vertices()
			c := r3.WeightedCentroid(
				[]r3.Point{
					vertexPositions[ev[0]],
					vertexPositions[ev[1]],
				},
				[]float64{
					float64(j + 1),
					float64(len(nV)),
				},
			)
			vertexPositions[nV[j]] = c
			gg.vertices = append(gg.vertices, nV[j])
		}
		vertexToEdgeMap[e] = nV
	}
}

func subdividedFace(face Face, gg *Geodesic, m int, newEdgeSet map[Edge]bool, vertexToEdgeMap map[Edge]([]Vertex)) (newEdges []Edge, newFaces []Face) {

	newFaces = make([]Face, 0)
	newEdges = make([]Edge, 0)

	v0 := face.Loop()[0]
	v1 := face.Loop()[1]
	v2 := face.Loop()[2]

	e0 := NewEdge(v0, v1)
	e1 := NewEdge(v1, v2)
	e2 := NewEdge(v2, v0)

	// Create subdivision vertices
	vertexRows := make([][]Vertex, m+1)
	rowSize := 1
	for row := 0; row < m+1; row++ {
		vertexRows[row] = make([]Vertex, rowSize)
		rowSize++
	}

	// Create new interior Vertices
	for row := 1; row < len(vertexRows)-1; row++ {
		for j := 1; j < (len(vertexRows[row]) - 1); j++ {
			vertexRows[row][j] = NewVertex()
			gg.vertices = append(gg.vertices, vertexRows[row][j])
		}
	}

	// Replace existing vertices fo correct linkage to neighbours
	vertexRows[0][0] = v0
	vertexRows[m][0] = v1
	vertexRows[m][m] = v2

	getReplacements := func(e Edge) []Vertex {
		rep := vertexToEdgeMap[e]
		if rep == nil {
			rep_reversed := vertexToEdgeMap[e.Reversed()]
			rep = make([]Vertex, len(rep_reversed))
			copy(rep, rep_reversed)
			for i, j := 0, len(rep)-1; i < j; i, j = i+1, j-1 {
				rep[i], rep[j] = rep[j], rep[i]
			}
		}
		return rep
	}
	rep0 := getReplacements(e0)
	rep1 := getReplacements(e1)
	rep2 := getReplacements(e2)

	for i, iR := 1, 0; i <= (m - 1); i, iR = i+1, iR+1 {
		// v0 -> v1
		vertexRows[i][0] = rep0[iR]
		// v1 -> v2
		vertexRows[m][i] = rep1[iR]
		// v2 -> v0
		vertexRows[i][len(vertexRows[i])-1] = rep2[iR]
	}

	// Walk through the rows of the vertices
	// Connect the vertices above and below
	connectNewFace := func(nV0, nV1, nV2 Vertex) {
		ne0 := NewEdge(nV0, nV1)
		ne1 := NewEdge(nV1, nV2)
		ne2 := NewEdge(nV2, nV0)

		if !newEdgeSet[ne0] {
			newEdgeSet[ne0] = true
			newEdges = append(newEdges, ne0)
		}
		if !newEdgeSet[ne1] {
			newEdgeSet[ne1] = true
			newEdges = append(newEdges, ne1)
		}
		if !newEdgeSet[ne2] {
			newEdgeSet[ne2] = true
			newEdges = append(newEdges, ne2)
		}

		nF := NewFace([]Vertex{nV0, nV1, nV2})

		newFaces = append(newFaces, nF)
	}

	for row := 0; row < m; row++ {
		for i, vertex := range vertexRows[row] {
			nv1 := vertexRows[row+1][i]
			nv2 := vertexRows[row+1][i+1]
			connectNewFace(vertex, nv1, nv2)
		}
	}
	for row := 1; row < m; row++ {
		for i := 0; i < len(vertexRows[row])-1; i++ {
			nv0 := vertexRows[row][i]
			nv1 := vertexRows[row][i+1]
			nv2 := vertexRows[row+1][i+1]
			// This creates duplicate edges (only needs to create faces)
			connectNewFace(nv0, nv1, nv2)
		}
	}
	return
}

// Subdivide applies the surface subdivision modifier to the geodesic using the given breakdown structure (n,m).
// Currently only the breakdown structure (2,0) is supported.
// For more information see https://en.wikibooks.org/wiki/Geodesic_Grids/Breakdown_structures
func (gg *Geodesic) Subdivide(m, n int) error {

	// TODO: implement Class I and II breakdowns.
	if m == n {
		return errors.New("Class II not supported")
	}
	if n != 0 {
		return errors.New("Class III not supported")
	}
	if m == 1 {
		return nil
	}
	if m != 2 {
		return errors.New("Only (m=2,n=0) subdivision supported.")
	}

	t := m*m + m*n + n*n
	newFaces := make([]Face, 0, 20*t)
	newEdges := make([]Edge, 0)
	newEdgeSet := make(map[Edge]bool)

	vertexToEdgeMap := make(map[Edge]([]Vertex))
	createVerticesForEdges(gg, m, vertexToEdgeMap)

	for _, face := range gg.faces {
		nE, nF := subdividedFace(face, gg, m, newEdgeSet, vertexToEdgeMap)
		newFaces = append(newFaces, nF...)
		newEdges = append(newEdges, nE...)
	}

	gg.setEdges(newEdges)
	gg.setFaces(newFaces)
	gg.m *= m
	gg.n *= n

	return nil
}
