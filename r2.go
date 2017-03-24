package polyhedra

import "errors"

type Polygon struct {
	edges []Edge
}

func SubdivideTriangle(triangle Triangle, n, m int) ([]Triangle, error) {
	if m == n {
		return nil, errors.New("Class II not supported")
	}
	if m != 0 {
		return nil, errors.New("Class III not supported")
	}
	if n == 1 {
		return []Triangle{triangle}, nil
	}

	vertexRows := make([][]Vertex, n+1)

	rowSize := 1
	for row := 0; row < n+1; row++ {
		vertexRows[row] = make([]Vertex, rowSize)
		for n := range vertexRows[row] {
			vertexRows[row][n] = NewVertex()
		}
		rowSize++
	}

	vertexRows[0][0] = triangle.vertices[0]
	vertexRows[n][0] = triangle.vertices[1]
	vertexRows[n][n] = triangle.vertices[2]

	resultTriangles := make([]Triangle, 0)

	// Walk through the rows of the triangle
	// Connect the vertices above and below

	for row := 0; row < n; row++ {
		for i, vertex := range vertexRows[row] {
			v1 := vertexRows[row+1][i]
			v2 := vertexRows[row+1][i+1]
			t := Triangle{[3]Vertex{vertex, v1, v2}}
			resultTriangles = append(resultTriangles, t)
		}
	}
	for row := 1; row < n; row++ {
		for i := 0; i < len(vertexRows[row])-1; i++ {
			v0 := vertexRows[row][i]
			v1 := vertexRows[row][i+1]
			v2 := vertexRows[row+1][i+1]
			t := Triangle{[3]Vertex{v0, v1, v2}}
			resultTriangles = append(resultTriangles, t)
		}
	}
	return resultTriangles, nil

}

type Triangle struct {
	vertices [3]Vertex
}

type Pentagon struct {
	edges    [5]Edge
	vertices [5]Vertex
}

type Hexagon struct {
	edges    [6]Edge
	vertices [6]Vertex
}
