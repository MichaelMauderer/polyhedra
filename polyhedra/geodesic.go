package polyhedra

import (
	"errors"
	"github.com/MichaelMauderer/geometry/r3"
)

type Geodesic struct {
	polyhedron
	m, n int
}

type IcosahedralGeodesic Geodesic

func NewIcosahedralGeodesic() *Geodesic {
	ico := newIcosahedron()
	geo := Geodesic{ico, 1, 0}
	return &geo
}

func cullDuplicates(edges []Edge) []Edge {
	result := make([]Edge, 0, len(edges))

	for _, newEdge := range edges {
		alreadyIn := false
		for _, existingEdge := range result {
			if newEdge.Equal(existingEdge) {
				alreadyIn = true
				break
			}
		}
		if !alreadyIn {
			result = append(result, newEdge)
		}
	}
	return result
}

func normEdge(v1, v2 Vertex) edge {
	if v1 > v2 {
		return edge{v2, v1}
	} else {
		return edge{v1, v2}
	}
}

func edgeSetToSlice(edges map[Edge]bool) []Edge {
	result := make([]Edge, 0, len(edges))

	for newEdge, _ := range edges {
		result = append(result, newEdge)
	}
	return result
}

func (gg *Geodesic) Subdivide(m, n int) error {

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
		return errors.New("Only (m=2,n=0) subdivission supported.")
	}

	t := m*m + m*n + n*n
	newFaces := make([]Face, 0, 20*t)
	newEdges := make(map[Edge]bool)

	newVertices := make(map[edge]([]Vertex))
	for _, e := range gg.Edges() {
		nV := make([]Vertex, m-1)
		for j := range nV {
			nV[j] = NewVertex()
			ev := e.Vertices()
			c := r3.WeightedCentroid(
				[]r3.Point3D{
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
		newVertices[e.(edge)] = nV
	}

	for _, face := range gg.faces {

		v0 := face.Loop()[0]
		v1 := face.Loop()[1]
		v2 := face.Loop()[2]

		e0 := edge{v0, v1}
		e1 := edge{v1, v2}
		e2 := edge{v2, v0}

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

		getReplacements := func(e edge) []Vertex {
			rep := newVertices[e]
			if rep == nil {
				rep_reversed := newVertices[e.Reversed().(edge)]
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
			ne0 := normEdge(nV0, nV1)
			ne1 := normEdge(nV1, nV2)
			ne2 := normEdge(nV2, nV0)

			newEdges[ne0] = true
			newEdges[ne1] = true
			newEdges[ne2] = true

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
	}

	// TODO: Avoid creation of duplicates in the first place.
	gg.SetEdges(edgeSetToSlice(newEdges))
	gg.SetFaces(newFaces)
	gg.m *= m
	gg.n *= n

	return nil
}
