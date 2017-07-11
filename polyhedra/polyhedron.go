package polyhedra

// Polyhedron represents the functionality provided by a polyhedron.
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

// NewPolyhedron creates a polyhedron from the given vertices, edges and faces.
func NewPolyhedron(vertices []Vertex, edges []Edge, faces []Face) Polyhedron {
	poly := polyhedron{vertices: vertices}
	poly.SetFaces(faces)
	poly.vertexNeighbors = make(map[Vertex][]Vertex)
	poly.AddEdges(edges)
	return &poly
}

// polyhedron represents a polyhedron consisting of vertices, edges and faces.
type polyhedron struct {
	faces    []Face
	vertices []Vertex

	vertexNeighbors map[Vertex][]Vertex
	edgeCache       []Edge

	edgeToFace map[Edge][]Face
}

// init initialises the polyhedrons access caches.
func (p *polyhedron) init() {
	p.vertexNeighbors = make(map[Vertex][]Vertex)
	p.edgeToFace = make(map[Edge][]Face)
}

// Vertices returns the polyhedrons vertices.
func (p *polyhedron) Vertices() []Vertex {
	return p.vertices
}

// Edges returns the polyhedrons edges.
func (p *polyhedron) Edges() []Edge {
	if len(p.edgeCache) == 0 {
		edges := make([]Edge, 0)
		for _, v := range p.vertices {
			vns := p.vertexNeighbors[v]
			for _, vn := range vns {
				edges = append(edges, Edge{v, vn})
			}
		}
		p.edgeCache = cullDuplicates(edges)
	}
	return p.edgeCache
}

// resetEdge caches invalidates the cache that contains the polyhedrons edges.
func (p *polyhedron) resetEdgeCache() {
	if len(p.edgeCache) > 0 {
		p.edgeCache = make([]Edge, 0)
	}
}

// Faces returns the polyhedrons faces.
func (p *polyhedron) Faces() []Face {
	return p.faces
}

// AddVertex adds the given vertex to the polyhedron.
func (p *polyhedron) AddVertex(v Vertex) {
	p.vertices = append(p.vertices, v)
}

// addSingleEdge adds a Edge between the two given vertices.
func (p *polyhedron) addSingleEdge(v1 Vertex, v2 Vertex) {
	vn, ok := p.vertexNeighbors[v1]
	if !ok {
		p.vertexNeighbors[v1] = make([]Vertex, 0)
		vn = p.vertexNeighbors[v1]
	}
	p.vertexNeighbors[v1] = append(vn, v2)
}

// AddEdge adds an Edge between the two given vertices.
func (p *polyhedron) AddEdge(v1 Vertex, v2 Vertex) {
	p.addSingleEdge(v1, v2)
	p.addSingleEdge(v2, v1)

	p.resetEdgeCache()
}

// AddEdges adds the given edges to the polyhedron.
func (p *polyhedron) AddEdges(edges []Edge) {
	for _, e := range edges {
		v := e.Vertices()
		p.AddEdge(v[0], v[1])
	}
}

// SetEdges clears all current edges and adds the given edges instead.
func (p *polyhedron) SetEdges(edges []Edge) {
	p.vertexNeighbors = make(map[Vertex][]Vertex, len(edges))
	p.AddEdges(edges)
}

// addFace adds the given Face to the polyhedron.
func (p *polyhedron) addFace(f Face) {
	p.faces = append(p.faces, f)
	for _, e := range f.Edges() {
		p.edgeToFace[e] = append(p.edgeToFace[e], f)
		redge := e.Reversed()
		p.edgeToFace[redge] = append(p.edgeToFace[redge], f)
	}
}

// AddFace adds a Face defined by the given vertices to the polyhedron.
func (p *polyhedron) AddFace(vertices []Vertex) {
	edges := make([]Edge, len(vertices))
	for i, vertex := range vertices {
		nextI := (i + 1) % len(vertices)
		edges[i] = Edge{vertex, vertices[nextI]}

	}
	f := NewFace(vertices)
	p.addFace(f)
}

// AddFaces adds all the given faces to the polyhedron.
func (p *polyhedron) AddFaces(faces []Face) {
	for _, face := range faces {
		p.addFace(face)
	}
}

// SetFaces clears all current faces and adds the given faces instead.
func (p *polyhedron) SetFaces(faces []Face) {
	p.faces = make([]Face, 0, len(faces))
	p.edgeToFace = make(map[Edge][]Face, len(faces))
	for _, face := range faces {
		p.addFace(face)
	}
}

// VertexDegree returns the number of neighbours of the given vertex.
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

// EdgeAdjacentFaces returns the faces that are adjacent to the given Edge.
func (p *polyhedron) EdgeAdjacentFaces(e Edge) [2]Face {
	faces := p.edgeToFace[e]
	return [2]Face{faces[0], faces[1]}
}

// FaceEdgeAdjacentFaces returns the faces that share an Edge with the given facce.
func (p *polyhedron) FaceEdgeAdjacentFaces(f Face) []Face {
	resultFaces := make([]Face, 0)
	for _, e := range f.Edges() {
		for _, ef := range p.EdgeAdjacentFaces(e) {
			if !f.Equals(ef) {
				resultFaces = append(resultFaces, ef)
			}
		}

	}
	return resultFaces
}

// FaceVertexAdjacentFaces returns the faces that share a vertex with the given Face.
func (p *polyhedron) FaceVertexAdjacentFaces(f Face) []Face {
	resultFaces := make([]Face, 0)
	for _, face := range p.faces {
		for _, v := range face.Loop() {
			for _, vf := range p.VertexAdjacentFaces(v) {
				if !f.Equals(vf) {
					resultFaces = append(resultFaces, f)
				}
			}
		}
	}
	return resultFaces
}

// AdjacentVertices returns all vertices that are part of an Edge with the given vertex.
func (p *polyhedron) AdjacentVertices(vertex Vertex) []Vertex {
	return p.vertexNeighbors[vertex]
}
