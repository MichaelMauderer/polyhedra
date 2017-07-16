package polyhedra

// Interface represents the functionality provided by a Polyhedron.
type Interface interface {
	Vertices() []Vertex
	Edges() []Edge
	Faces() []Face

	VertexDegree(vertex Vertex) int
	AdjacentVertices(vertex Vertex) []Vertex

	VertexAdjacentFaces(v Vertex) []Face
	EdgeAdjacentFaces(e Edge) [2]Face
	FaceEdgeAdjacentFaces(f Face) []Face
	FaceVertexAdjacentFaces(f Face) []Face
}
