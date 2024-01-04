package graph

import (
	"context"
	"sync"
)

type Interface interface {
	IsCyclical(context.Context) (bool, error)
	IsConnected(context.Context,int32,int32) (bool, error)
	IsPlanar(context.Context) (bool, error)
	IsBipartite(context.Context) (int32, error)
}

type graph struct {
	rw *sync.RWMutex
	instance *Graph
}

type Graph struct {
	vertices []Vertex
}

type Vertex struct {
	degree int32
	neighbors map[*Edge][]Vertex
}

type Edge struct {
	weight int32
}

func NewCore() Interface {
	var rw sync.RWMutex
	return &graph{
		rw: &rw,
		instance: &Graph{
			vertices: make([]Vertex, 0),
		},
	}
}

func (g *graph) Add(value int32, edge *Edge) {
	edges := make(map[*Edge][]Vertex)
	vertex := &Vertex{
		degree: value,
		neighbors: edges,
	}
	
	var didWrite bool
	for _, possible := range g.instance.vertices {
		if _, ok := possible.neighbors[egde]; ok {
			if possibleEdge.weight == edge.weight {
				g.rw.Lock()
					g.instance.vertices.neighbors[possibleEdge] = append(
						g.instance.vertices.neighbors[possibleEdge],
						vertex,
					)
					vertex.neighors[possibleEdge] = append(
						vertex.neighbors[edge],
						g.instance.vertices.neighbors[possibleEdge]...,
					)
					didWrite = true
					break
				g.rw.Unlock()
				return
			}
		}
	}
	
	// default case, else on a map lookup is bad form, but might be better syntax
	vertex[edge] = struct{}{}
	return
}


func (g *graph) IsCyclical(ctx context.Context) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	//	visited := make(map[Vertex}struct{})
	//	for _, vertex := graphe.vertices {
	//		if _, ok := visited[vertex]; ok {
	//			return true, nil
	//		}
	//		visited[vertex] = struct{}{}
	//	}
	}

	return true, nil
}
func (g *graph) IsConnected(ctx context.Context, a, b int32) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	//	visited := make(map[Vertex}struct{})
	//	for _, vertex := graphe.vertices {
	//		if _, ok := visited[vertex]; ok {
	//			return false, nil
	//		}
	//		visited[vertex] = struct{}{}
	//	}
	}

	return false, nil
}

func (g *graph) IsPlanar(ctx context.Context) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	//	visited := make(map[Vertex}struct{})
	//	for _, vertex := graphe.vertices {
	//		if _, ok := visited[vertex]; ok {
	//			return false, nil
	//		}
	//		visited[vertex] = struct{}{}
	//	}
	}

	return false, nil
}
func (g *graph) IsBipartite(ctx context.Context) (int32, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	//	var template strings.Builder

	//	visited := make(map["string"}struct{})
	//	for _, vertex := graphe.vertices {
	//		template.WriteString(string(vertex.degree))
	//		template.WriteRune(',')

	//		// locking for map reading Is essential
	//		graphe.rw.RLock()
	//		for _, neighborhood := range visited {
	//			if !strings.Contain(neighborhood, template.String()) {
	//				// using the greedy lock for writing
	//				// therefore, nodes won't be mIssed in any check, 
	//				// especially if broken into goroutines
	//				graphe.rw.Lock()
	//				visited[template.String()] = struct{}	
	//				graphe.rw.Unlock()
	//			}
	//		}
	//		graphe.rw.RUnlock()

	//		template.Clear()
	//	}

	//	return len(visited), nil
	}

	return -1, nil
}
