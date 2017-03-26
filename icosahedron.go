package polyhedra

import (
	"math"
)

func NewIcosahedron() Polyhedron {

	ico := Polyhedron{}

	c1 := math.Cos(2.0 * math.Pi / 5.0)
	c2 := math.Cos(math.Pi / 5.0)
	s1 := math.Sin(2.0 * math.Pi / 5.0)
	s2 := math.Sin(4.0 * math.Pi / 5.0)
	h := math.Sqrt(3) / 2
	vertexPos := []Point3D{
		// Top Vertex
		{0, 0, 1},
		// Bottom Vertex
		{0, 0, -1},
		// Top Pentagon
		{0, -1, h},
		{s1, -c1, h},
		{s2, c2, h},
		{-s2, c2, h},
		{-s1, -c1, h},
		// Bottom Pentagon
		{s2, -c2, -h},
		{s1, c1, -h},
		{0, 1, -h},
		{-s1, c1, -h},
		{-s2, -c2, -h},
	}

	ico.vertices = make([]Vertex, 12)
	for i := range ico.vertices {
		ico.vertices[i] = NewVertex()
		ico.vertices[i].setPosition(vertexPos[i])
	}

	topVertex := ico.vertices[0]
	bottomVertex := ico.vertices[1]

	topPentagon := ico.vertices[2:2+5]
	bottomPentagon := ico.vertices[7:7+5]

	connectPoles := func(pentagon []Vertex, poleVertex Vertex) {
		for i, vertex := range pentagon {
			neighborLIndex := ( 5 + i - 1 ) % 5
			neighborL := pentagon[neighborLIndex]

			ico.AddEdge(vertex, neighborL)
			ico.AddEdge(vertex, poleVertex)

			ico.AddFace([]Vertex{vertex, neighborL, poleVertex})
		}
	}
	// Connect bottom and top poles
	connectPoles(topPentagon, topVertex)
	connectPoles(bottomPentagon, bottomVertex)

	// Connect bottom pentagon
	for i, vertex := range bottomPentagon {
		bottomNeighbor := bottomPentagon[(5+i-1)%5]
		topNeighbor := topPentagon[i]

		ico.AddEdge(vertex, topNeighbor)
		ico.AddFace([]Vertex{vertex, topNeighbor, bottomNeighbor})
	}

	//Connect top pentagon
	for i, vertex := range topPentagon {
		topNeighbor := topPentagon[(5+i-1)%5]
		bottomNeighbor := bottomPentagon[(5+i-1)%5]

		ico.AddEdge(vertex, bottomNeighbor)
		ico.AddFace([]Vertex{vertex, topNeighbor, bottomNeighbor})
	}
	return ico
}
