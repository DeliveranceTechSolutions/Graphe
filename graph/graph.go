package graph

type Graph struct {
	vertices map[*Vertex]Neighbors // maps are always references
}

type Vertex struct {
	degree int32
	neighbors Neighbors
}

type Neighbors struct {
	nextDoor map[*Vertex]bool
}

type Edge struct {
	weight int32
}

func NewCore() *Graph {
	return &Graph{}
}

func (g *Graph) isNeighbor() {

}

func (g Graph) IsCycle() {

}

func (g Graph) IsPlanar() {

}

func (g Graph) IsBiparte() {
	for vertex, neighbors := range g.vertices {
		if _, ok := neighbors.nextDoor[vertex]; !ok {
			
		}
	}
}