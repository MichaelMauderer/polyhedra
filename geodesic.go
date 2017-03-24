package polyhedra

import (
	"errors"
	"fmt"
	"log"
)

type Geodesic struct {
	Polyhedron
	m, n int
}

type IcosahedralGeodesic struct {
	Geodesic
}

func NewIcosahedralGeodesic() *IcosahedralGeodesic {
	ico := NewIcosahedron()
	geo := Geodesic{ico, 1, 0}
	icoGeo := IcosahedralGeodesic{geo}
	return &icoGeo
}

func cullDuplicates(edges []Edge) []Edge {
	result := make([]Edge, 0, len(edges))

	for _, newEdge := range edges {
		alreadyIn := false
		for _, existingEdge := range result {
			if newEdge == existingEdge || newEdge.Reversed() == existingEdge {
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

func (ig *IcosahedralGeodesic) checkFaces() error {
	faceNum := len(ig.faces)
	if faceNum%20 != 0 {
		return errors.New("Number of faces is not a multiple of 20.")
	}
	return nil
}
func (ig *IcosahedralGeodesic) checkEdges() error {
	edgeNum := len(ig.edges)
	if edgeNum%30 != 0 {
		return errors.New("Number of faces is not a multiple of 30.")
	}
	return nil
}
func (ig *IcosahedralGeodesic) checkVertices() error {
	vertexNum := len(ig.vertices)
	if (vertexNum-2)%10 != 0 {
		return errors.New("Number of vertices does not fulfill V=(T*10+2)")
	}
	foundWrongOne := false
	for _, vertex := range ig.vertices {
		if vertex == 0 {
			return errors.New(fmt.Sprintf("Contains illegal zero vertex."))
		}
		vD := ig.VertexDegree(vertex)
		if (vD != 5) && (vD != 6) {
			log.Printf("Vertex %v in %v has degree %v", vertex, ig, vD)
			foundWrongOne = true
		}
	}
	if foundWrongOne {
		return errors.New(fmt.Sprintf("Found invalid number of edges at vertex. Should be 5 or 6"))
	}
	return nil
}

func (gg*IcosahedralGeodesic) CheckIntegrity() error {
	error := gg.checkFaces()
	if error != nil {
		return error
	}
	error = gg.checkEdges()
	if error != nil {
		return error
	}
	error = gg.checkVertices()
	if error != nil {
		return error
	}
	return nil
}

func (gg *Geodesic) subdivide(m, n int) error {

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
	newEdges := make([]Edge, 0, 30*t)

	newVertices := make(map[Edge]([]Vertex))
	for edge_i := range gg.edges {
		nV := make([]Vertex, m-1)
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
			rep := newVertices[e]
			if rep == nil {
				rep_reversed := newVertices[e.Reversed()]
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
			ne0 := Edge{nV0, nV1}
			ne1 := Edge{nV1, nV2}
			ne2 := Edge{nV2, nV0}

			newEdges = append(newEdges, ne0)
			newEdges = append(newEdges, ne1)
			newEdges = append(newEdges, ne2)

			nF := Face{[]Edge{ne0, ne1, ne2}}

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
	gg.edges = cullDuplicates(newEdges)
	gg.faces = newFaces
	gg.m *= m
	gg.n *= n

	return nil
}
