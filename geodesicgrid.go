package polyhedra

import (
	"errors"
	"fmt"
)

type GeodesicGrid struct {
	Polyhedron
}

type IcoGG struct {
	GeodesicGrid
}

func (gg*IcoGG) checkIntegrity() error {
	faceNum := len(gg.faces)
	if faceNum%20 != 0 {
		return errors.New("Number of faces is not a multiple of 20.")
	}
	edgeNum := len(gg.edges)
	if edgeNum%30 != 0 {
		return errors.New("Number of faces is not a multiple of 30.")
	}
	vertexNum := len(gg.vertices)
	if (vertexNum-2)%10 != 0 {
		return errors.New("Number of vertices does not fulfill V=(T*10+2)")
	}
	for _, vertex := range gg.vertices {
		vD := gg.VertexDegree(vertex)
		if (vD != 5) || (vD != 6) {
			return errors.New(fmt.Sprintf("Invalid number of edges at vertex %v: %v. Should be 5 or 6", vertex, vD))
		}
	}
	return nil
}

func (gg *GeodesicGrid) subdivide(n, m int) error {

	if m == n {
		return errors.New("Class II not supported")
	}
	if m != 0 {
		return errors.New("Class III not supported")
	}
	if n == 1 {
		return nil
	}

	t := m*m + m*n + n*n

	newFaces := make([]Face, 0, 20*t)
	newEdges := make([]Edge, 0, 30*t)

	newVertices := make(map[Edge]([]Vertex))
	for edge_i := range gg.edges {
		nV := make([]Vertex, n-1)
		for j := range nV {
			nV[j] = NewVertex()
			gg.vertices = append(gg.vertices, nV[j])
		}
		newVertices[gg.edges[edge_i]] = nV
	}

	for _, face := range gg.faces {

		e0 := face.edges[0]
		v0 := e0.v1
		e1 := face.edges[1]
		v1 := e1.v1
		e2 := face.edges[2]
		v2 := e2.v1

		// Create subdivision vertices
		vertexRows := make([][]Vertex, n+1)
		rowSize := 1
		for row := 0; row < n+1; row++ {
			vertexRows[row] = make([]Vertex, rowSize)
			for n := range vertexRows[row] {
				vertexRows[row][n] = NewVertex()
			}
			rowSize++
		}

		// Replace existing vertices fo correct linkage to neighbours
		vertexRows[0][0] = v0
		vertexRows[n][0] = v1
		vertexRows[n][n] = v2

		getReplacements := func(e Edge) []Vertex {
			rep := newVertices[e]
			if rep == nil {
				rep_reversed := newVertices[e.Reversed()]
				rep = make([]Vertex, len(rep_reversed))
				for i, j := 0, len(rep_reversed)-1; i < j; i, j = i+1, j-1 {
					rep[i], rep[j] = rep_reversed[j], rep_reversed[i]
				}
			}
			return rep
		}
		rep0 := getReplacements(e0)
		rep1 := getReplacements(e1)
		rep2 := getReplacements(e2)

		for i, iR := 1, 0; i < (n - 1); i, iR = i+1, i {
			// v0 -> v1
			vertexRows[i][0] = rep0[iR]
			// v1 -> v2
			vertexRows[n][i] = rep1[iR]
			// v2 -> v0
			vertexRows[i][i] = rep2[iR]
		}

		// Walk through the rows of the vertices
		// Connect the vertices above and below
		connectNewFace := func(nV0, nV1, nV2 Vertex) {
			ne0 := Edge{nV0, nV1}
			ne1 := Edge{nV1, nV2}
			ne2 := Edge{nV2, nV0}

			newEdges = append(newEdges, ne0)
			newEdges = append(newEdges, ne1)
			newEdges = append(newEdges, ne2)

			nF := Face{[]Edge{ne0, ne1, ne2}}

			newFaces = append(newFaces, nF)
		}

		for row := 0; row < n; row++ {
			for i, vertex := range vertexRows[row] {
				nv1 := vertexRows[row+1][i]
				nv2 := vertexRows[row+1][i+1]
				connectNewFace(vertex, nv1, nv2)
			}
		}
		for row := 1; row < n; row++ {
			for i := 0; i < len(vertexRows[row])-1; i++ {
				nv0 := vertexRows[row][i]
				nv1 := vertexRows[row][i+1]
				nv2 := vertexRows[row+1][i+1]
				connectNewFace(nv0, nv1, nv2)
			}
		}
	}
	gg.edges = newEdges
	gg.faces = newFaces
	return nil
}
