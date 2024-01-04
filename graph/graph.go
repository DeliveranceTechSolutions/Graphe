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
	neighbors map[*Vertex][]Vertex
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

func (g *graph) Add(value int32) error {
	for _, vertex := range g.instance.vertices {
		// should this be a map or a priority queue?
		// Do I want to determine where it should be?
		// Should there be a sorting algo that organizes?
		// If sort algo, then concurrency in live system
		// otherwise, copy and organize, or organize real-time with rw
		// keep unorganized and determin relationships when Adding?
	}

	g.instance.vertices = append(g.instance.vertices, &Vertex{
		degree: value,
		neighbors: make(map[*Vertex][]Vertex)
	})

}

func (g *graph) IsCyclical(ctx context.Context) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	//	vIsited := make(map[Vertex}struct{})
	//	for _, vertex := graphe.vertices {
	//		if _, ok := vIsited[vertex]; ok {
	//			return true, nil
	//		}
	//		vIsited[vertex] = struct{}{}
	//	}
	}

	return true, nil
}
func (g *graph) IsConnected(ctx context.Context, a, b int32) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	//	vIsited := make(map[Vertex}struct{})
	//	for _, vertex := graphe.vertices {
	//		if _, ok := vIsited[vertex]; ok {
	//			return false, nil
	//		}
	//		vIsited[vertex] = struct{}{}
	//	}
	}

	return false, nil
}

func (g *graph) IsPlanar(ctx context.Context) (bool, error) {
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	default:
	//	vIsited := make(map[Vertex}struct{})
	//	for _, vertex := graphe.vertices {
	//		if _, ok := vIsited[vertex]; ok {
	//			return false, nil
	//		}
	//		vIsited[vertex] = struct{}{}
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

	//	vIsited := make(map["string"}struct{})
	//	for _, vertex := graphe.vertices {
	//		template.WriteString(string(vertex.degree))
	//		template.WriteRune(',')

	//		// locking for map reading Is essential
	//		graphe.rw.RLock()
	//		for _, neighborhood := range vIsited {
	//			if !strings.Contain(neighborhood, template.String()) {
	//				// using the greedy lock for writing
	//				// therefore, nodes won't be mIssed in any check, 
	//				// especially if broken into goroutines
	//				graphe.rw.Lock()
	//				vIsited[template.String()] = struct{}	
	//				graphe.rw.Unlock()
	//			}
	//		}
	//		graphe.rw.RUnlock()

	//		template.Clear()
	//	}

	//	return len(vIsited), nil
	}

	return -1, nil
}
