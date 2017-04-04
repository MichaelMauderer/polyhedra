package polyhedra

type Polyhedron interface {
	Vertices() []Vertex
	Edges() []Edge
	Faces() []Face

	AddVertex(Vertex)
	AddEdge(Vertex, Vertex)
	AddFace([]Vertex)

	VertexDegree(vertex Vertex) int
	AdjacentVertices(vertex Vertex) []Vertex

	VertexAdjacentFaces(v Vertex) []Face
	EdgeAdjacentFaces(e Edge) [2]Face
	FaceEdgeAdjacentFaces(f Face) []Face
	FaceVertexAdjacentFaces(f Face) []Face
}

//func NewPolyhedron([]Vertex, []Edge, []Face) Polyhedron

type polyhedron struct {
	faces    []Face
	edges    []Edge
	vertices []Vertex
}

func (p *polyhedron) Vertices() []Vertex {
	return p.vertices
}

func (p *polyhedron) Edges() []Edge {
	return p.edges
}

func (p *polyhedron) Faces() []Face {
	return p.faces
}

func (p *polyhedron) AddVertex(v Vertex) {
	p.vertices = append(p.vertices, v)
}

func (p *polyhedron) AddEdge(v1 Vertex, v2 Vertex) {
	p.edges = append(p.edges, Edge{v1, v2})
}

func (p *polyhedron) AddFace(vertices []Vertex) {
	edges := make([]Edge, len(vertices))
	for i, vertex := range vertices {
		nextI := (i + 1) % len(vertices)
		edges[i] = Edge{vertex, vertices[nextI]}
	}
	p.faces = append(p.faces, NewFace(vertices))
}

func (p *polyhedron) VertexDegree(vertex Vertex) int {
	degree := 0
	for _, edge := range p.edges {
		if vertex == edge.v1 || vertex == edge.v2 {
			degree += 1
		}
	}
	return degree
}

func (p *polyhedron) VertexAdjacentFaces(v Vertex) []Face {
	resultFaces := make([]Face, 0)
	for i, face := range p.faces {
		for _, vf := range face.Loop() {
			if v == vf {
				resultFaces = append(resultFaces, p.faces[i])
			}
		}
	}
	return resultFaces
}

func (p *polyhedron) EdgeAdjacentFaces(e Edge) [2]Face {
	var resultFaces [2]Face
	iR := 0
	for i, face := range p.Faces() {
		for _, ve := range face.Edges() {
			if e.Equal(ve) {
				resultFaces[iR] = p.faces[i]
				iR += 1
			}
		}
	}
	return resultFaces
}

func (p *polyhedron) FaceEdgeAdjacentFaces(f Face) []Face {
	resultFaces := make([]Face, 0)
	for _, e := range f.Edges() {
		for _, ef := range p.EdgeAdjacentFaces(e) {
			if f != ef {
				resultFaces = append(resultFaces, ef)
			}
		}

	}
	return resultFaces
}

func (p *polyhedron) FaceVertexAdjacentFaces(f Face) []Face {
	resultFaces := make([]Face, 0)
	for _, face := range p.faces {
		for _, v := range face.Loop() {
			for _, vf := range p.VertexAdjacentFaces(v) {
				if f != vf {
					resultFaces = append(resultFaces, f)
				}
			}
		}
	}
	return resultFaces
}

func (p *polyhedron) AdjacentVertices(vertex Vertex) []Vertex {
	result := make([]Vertex, 0)
	for _, edge := range p.edges {
		if vertex == edge.v1 {
			result = append(result, edge.v2)
		}
		if vertex == edge.v2 {
			result = append(result, edge.v1)
		}
	}
	return result
}
