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

func NewPolyhedron(vertices []Vertex, edges []Edge, faces []Face) Polyhedron {
	poly := polyhedron{faces: faces, vertices: vertices}
	poly.vertexNeighbors = make(map[Vertex][]Vertex)
	poly.AddEdges(edges)
	return &poly
}

type polyhedron struct {
	faces    []Face
	vertices []Vertex

	vertexNeighbors map[Vertex][]Vertex
	edgeCache       []Edge
}

func (p *polyhedron) init() {
	p.vertexNeighbors = make(map[Vertex][]Vertex)
}

func (p *polyhedron) Vertices() []Vertex {
	return p.vertices
}

func (p *polyhedron) Edges() []Edge {
	if len(p.edgeCache) == 0 {
		edges := make([]Edge, 0)
		for v, vns := range p.vertexNeighbors {
			for _, vn := range vns {
				edges = append(edges, edge{v, vn})
			}
		}
		p.edgeCache = cullDuplicates(edges)
	}
	return p.edgeCache
}

func (p *polyhedron) resetEdgeCache() {
	if len(p.edgeCache) > 0 {
		p.edgeCache = make([]Edge, 0)
	}
}

func (p *polyhedron) Faces() []Face {
	return p.faces
}

func (p *polyhedron) AddVertex(v Vertex) {
	p.vertices = append(p.vertices, v)
}

func (p *polyhedron) addSingleEdge(v1 Vertex, v2 Vertex) {
	vn, ok := p.vertexNeighbors[v1]
	if !ok {
		p.vertexNeighbors[v1] = make([]Vertex, 0)
		vn = p.vertexNeighbors[v1]
	}
	p.vertexNeighbors[v1] = append(vn, v2)
}

func (p *polyhedron) AddEdge(v1 Vertex, v2 Vertex) {
	p.addSingleEdge(v1, v2)
	p.addSingleEdge(v2, v1)

	p.resetEdgeCache()
}

func (p *polyhedron) AddEdges(edges []Edge) {
	for _, e := range edges {
		v := e.Vertices()
		p.AddEdge(v[0], v[1])
	}
}

func (p *polyhedron) SetEdges(edges []Edge) {
	p.vertexNeighbors = make(map[Vertex][]Vertex, len(edges))
	p.AddEdges(edges)
}

func (p *polyhedron) AddFace(vertices []Vertex) {
	edges := make([]edge, len(vertices))
	for i, vertex := range vertices {
		nextI := (i + 1) % len(vertices)
		edges[i] = edge{vertex, vertices[nextI]}
	}
	p.faces = append(p.faces, NewFace(vertices))
}

func (p *polyhedron) VertexDegree(vertex Vertex) int {
	return len(p.vertexNeighbors[vertex])
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
	return p.vertexNeighbors[vertex]
}
