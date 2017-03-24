package polyhedra

type Interface interface{}

type Polyhedron struct {
	faces    []Face
	edges    []Edge
	vertices []Vertex
}

func (p *Polyhedron) AddEdge(v1 Vertex, v2 Vertex) {
	p.edges = append(p.edges, Edge{v1, v2})
}

func (p *Polyhedron) AddFace(vertices []Vertex) {
	edges := make([]Edge, len(vertices))
	for i, vertex := range vertices {
		nextI := (i + 1) % len(vertices)
		edges[i] = Edge{vertex, vertices[nextI]}
	}
	p.faces = append(p.faces, Face{edges})
}

type Face struct {
	edges []Edge
}

type Edge struct {
	v1, v2 Vertex
}

func (e Edge) Contains(v Vertex) bool {
	return e.v1 == v || e.v2 == v
}

func (e Edge) Equal(other Edge) bool {
	if e.v1 == other.v1 && e.v2 == other.v2 {
		return true
	}
	return e.v1 == other.v2 && e.v2 == other.v1
}
func (e Edge) Reversed() Edge {
	return Edge{e.v2, e.v1}
}

type Vertex uint

var vertexId Vertex = 0
var vertexPositions = make(map[Vertex]CartesianCoordinates)

func NewVertex() Vertex {
	vertexId++
	return vertexId
}

func (v Vertex) setPosition(coords CartesianCoordinates) {
	vertexPositions[v] = coords
}

func (v Vertex) Position() CartesianCoordinates {
	return vertexPositions[v]
}
