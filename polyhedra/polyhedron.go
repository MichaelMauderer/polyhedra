// Package polyhedra implements basic functionality to create and modify geometric polyhedra.
package polyhedra

// NewPolyhedron creates a Polyhedron from the given vertices, edges and faces.
func NewPolyhedron(vertices []Vertex, edges []Edge, faces []Face) (*Polyhedron, error) {
	poly := Polyhedron{vertices: vertices}
	poly.setFaces(faces)
	poly.vertexNeighbors = make(map[Vertex][]Vertex)
	poly.addEdges(edges)
	return &poly, nil
}

// Polyhedron represents a Polyhedron consisting of vertices, edges and faces.
type Polyhedron struct {
	faces    []Face
	vertices []Vertex

	vertexNeighbors map[Vertex][]Vertex
	edgeCache       []Edge

	edgeToFace map[Edge][]Face
}

// init initialises the polyhedrons access caches.
func (p *Polyhedron) init() {
	p.vertexNeighbors = make(map[Vertex][]Vertex)
	p.edgeToFace = make(map[Edge][]Face)
}

// Vertices returns the polyhedrons vertices.
func (p *Polyhedron) Vertices() []Vertex {
	return p.vertices
}

// Edges returns the polyhedrons edges.
func (p *Polyhedron) Edges() []Edge {
	if len(p.edgeCache) == 0 {
		edges := make([]Edge, 0)
		for _, v := range p.vertices {
			vns := p.vertexNeighbors[v]
			for _, vn := range vns {
				edges = append(edges, NewEdge(v, vn))
			}
		}
		p.edgeCache = cullDuplicates(edges)
	}
	return p.edgeCache
}

// resetEdge caches invalidates the cache that contains the polyhedrons edges.
func (p *Polyhedron) resetEdgeCache() {
	if len(p.edgeCache) > 0 {
		p.edgeCache = make([]Edge, 0)
	}
}

// Faces returns the polyhedrons faces.
func (p *Polyhedron) Faces() []Face {
	return p.faces
}

// addVertex adds the given vertex to the Polyhedron.
func (p *Polyhedron) addVertex(v Vertex) {
	p.vertices = append(p.vertices, v)
}

// addSingleEdge adds a Edge between the two given vertices.
func (p *Polyhedron) addSingleEdge(v1 Vertex, v2 Vertex) {
	vn, ok := p.vertexNeighbors[v1]
	if !ok {
		p.vertexNeighbors[v1] = make([]Vertex, 0)
		vn = p.vertexNeighbors[v1]
	}
	p.vertexNeighbors[v1] = append(vn, v2)
}

// addEdge adds an Edge between the two given vertices.
func (p *Polyhedron) addEdge(v1 Vertex, v2 Vertex) error {
	p.addSingleEdge(v1, v2)
	p.addSingleEdge(v2, v1)

	p.resetEdgeCache()
	return nil
}

// addEdges adds the given edges to the Polyhedron.
func (p *Polyhedron) addEdges(edges []Edge) {
	for _, e := range edges {
		v := e.Vertices()
		err := p.addEdge(v[0], v[1])
		if err != nil {
			panic("Added illegal edge.")
		}
	}
}

// setEdges clears all current edges and adds the given edges instead.
func (p *Polyhedron) setEdges(edges []Edge) {
	p.vertexNeighbors = make(map[Vertex][]Vertex, len(edges))
	p.addEdges(edges)
}

// addFace adds the given Face to the Polyhedron.
func (p *Polyhedron) addFace(f Face) {
	p.faces = append(p.faces, f)
	for _, e := range f.Edges() {
		p.edgeToFace[e] = append(p.edgeToFace[e], f)
		redge := e.Reversed()
		p.edgeToFace[redge] = append(p.edgeToFace[redge], f)
	}
}

// addFaceFromLoop adds a Face defined by the given vertices to the Polyhedron.
func (p *Polyhedron) addFaceFromLoop(vertices []Vertex) {
	edges := make([]Edge, len(vertices))
	for i, vertex := range vertices {
		nextI := (i + 1) % len(vertices)
		edges[i] = NewEdge(vertex, vertices[nextI])

	}
	f := NewFace(vertices)
	p.addFace(f)
}

// addFaces adds all the given faces to the Polyhedron.
func (p *Polyhedron) addFaces(faces []Face) {
	for _, face := range faces {
		p.addFace(face)
	}
}

// setFaces clears all current faces and adds the given faces instead.
func (p *Polyhedron) setFaces(faces []Face) {
	p.faces = make([]Face, 0, len(faces))
	p.edgeToFace = make(map[Edge][]Face, len(faces))
	for _, face := range faces {
		p.addFace(face)
	}
}

// VertexDegree returns the number of neighbours of the given vertex.
func (p *Polyhedron) VertexDegree(vertex Vertex) int {
	return len(p.vertexNeighbors[vertex])
}

// VertexAdjacentFaces returns all faces that contain the given vertex.
func (p *Polyhedron) VertexAdjacentFaces(v Vertex) []Face {
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
func (p *Polyhedron) EdgeAdjacentFaces(e Edge) [2]Face {
	faces := p.edgeToFace[e]
	return [2]Face{faces[0], faces[1]}
}

// FaceEdgeAdjacentFaces returns the faces that share an Edge with the given facce.
func (p *Polyhedron) FaceEdgeAdjacentFaces(f Face) []Face {
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
func (p *Polyhedron) FaceVertexAdjacentFaces(f Face) []Face {
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
func (p *Polyhedron) AdjacentVertices(vertex Vertex) []Vertex {
	return p.vertexNeighbors[vertex]
}
