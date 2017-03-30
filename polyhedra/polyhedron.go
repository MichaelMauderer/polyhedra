package polyhedra

import (
	"fmt"
	"github.com/MichaelMauderer/geometry/r3"
)

type Polyhedron interface{
	Vertices() []Vertex
	Edges() []Edge
	Faces() []*Face

	VertexDegree(vertex Vertex) int
	AdjacentVertices(vertex Vertex) []Vertex

	VertexAdjacentFaces(v Vertex) []*Face
	EdgeAdjacentFaces(e Edge) [2]*Face
	FaceEdgeAdjacentFaces(f *Face) []*Face
	FaceVertexAdjacentFaces(f *Face) []*Face
}

type polyhedron struct {
	faces    []Face
	edges    []Edge
	vertices []Vertex
}

func (p *polyhedron) AddEdge(v1 Vertex, v2 Vertex) {
	p.edges = append(p.edges, Edge{v1, v2})
}

func (p *polyhedron) Vertices() []Vertex {
	return p.vertices
}

func (p *polyhedron) Edges() []Edge {
	return p.edges
}

func (p *polyhedron) Faces() []*Face {
	result := make([]*Face, len(p.faces))
	for i:= range (p.faces){
		result[i] = &p.faces[i]
	}
	return result
}

func (p *polyhedron) AddFace(vertices []Vertex) {
	edges := make([]Edge, len(vertices))
	for i, vertex := range vertices {
		nextI := (i + 1) % len(vertices)
		edges[i] = Edge{vertex, vertices[nextI]}
	}
	p.faces = append(p.faces, Face{vertices})
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

func (p *polyhedron) VertexAdjacentFaces(v Vertex) []*Face {
	resultFaces := make([]*Face, 0)
	for i, face := range p.faces {
		for _, vf  := range face.Loop{
			if v == vf{
				resultFaces = append(resultFaces, &p.faces[i])
			}
		}
	}
	return resultFaces
}

func (p *polyhedron) EdgeAdjacentFaces(e Edge) [2]*Face {
	var resultFaces [2]*Face
	iR := 0
	for i, face := range p.Faces() {
		for _, ve  := range face.Edges(){
			if e.Equal(ve){
				resultFaces[iR] =&p.faces[i]
				iR +=1
			}
		}
	}
	return resultFaces
}

func (p *polyhedron) FaceEdgeAdjacentFaces(f *Face) []*Face {
	resultFaces := make([]*Face, 0)
		for _, e  := range f.Edges(){
			for _, ef := range p.EdgeAdjacentFaces(e){
				if f != ef{
					resultFaces = append(resultFaces, ef)
				}
			}


		}
return resultFaces
}

func (p *polyhedron) FaceVertexAdjacentFaces(f *Face) []*Face {
	resultFaces := make([]*Face, 0)
	for _, face := range p.faces {
		for _, v  := range face.Loop{
			for _, vf := range p.VertexAdjacentFaces(v){
				if f != vf{
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

type Face struct {
	Loop []Vertex
}

func (f Face) Edges() []Edge{
	edges := make([]Edge, len(f.Loop))
	for i := range f.Loop{
		v1 := f.Loop[i]
		v2 := f.Loop[(i+1)%len(f.Loop)]
		edges[i] = Edge{v1, v2}
	}
	return edges
}

func vertexCentroid(vertices []Vertex) r3.Point3D{
	positions := make([]r3.Point3D, len(vertices))
	for i, v := range vertices{
		positions[i] = v.Position()
	}
	return r3.Centroid3D(positions)
}

func (f Face) Center() r3.Point3D{
	return vertexCentroid(f.Loop)
}

type Edge struct {
	v1, v2 Vertex
}

func (e Edge) Length() float64 {
	return r3.Distance(e.v1.Position(), e.v2.Position())
}

func (e Edge) Center() r3.Point3D {
	return r3.Centroid3D([]r3.Point3D{e.v1.Position(), e.v2.Position()})
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
var vertexPositions = make(map[Vertex]r3.Point3D)

func NewVertex() Vertex {
	vertexId++
	return vertexId
}

func (v Vertex) setPosition(coords r3.Point3D) {
	vertexPositions[v] = coords
}

func (v Vertex) Position() r3.Point3D {
	return vertexPositions[v]
}

func (v Vertex) String() string {
	return fmt.Sprintf("Vertex(id=%v, pos=%v)", uint(v), v.Position())
}



func SortedClockwise(vertices []Vertex) []Vertex{
 	//Insertionsort based on clockwiseness
	c := vertexCentroid(vertices)
	n := r3.Point3D{0,0,0}.VectorTo(c).Normalised()
	sorted := make([]Vertex, 1)
	sorted[0] = vertices[0]
	for _, v := range vertices[1:]{
		i := 0
		for ; ; i++{
			if i == len(sorted){
				break
			}
			vo := sorted[i]
			if !r3.IsCCW(v.Position(), vo.Position(), c, n){
				break
			}
		}
		//insert at i
		sorted = append(sorted, 0)
		copy(sorted[i+1:], sorted[i:])
		sorted[i] = v
	}
	return sorted
}
