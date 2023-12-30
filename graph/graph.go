package graph

type Graph struct {
	vertices [][]Vertex
}

type Vertex struct {
	degree int32
	neighbors map[*Vertex][]Vertex
}

type Edge struct {
	weight int32
}

func NewCore() *Graph {

}