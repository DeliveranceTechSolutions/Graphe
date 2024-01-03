package graph

import (
	"sync"
	"strings"
)

type Interface interface {
	isCyclical(ctx) (bool, error)
	isConnected(ctx,int32,int32) (bool, error)
	isPlanar(ctx) (bool, error)
	isBipartite(ctx) (int32, error)
}

type graph interface {
	rw *sync.RWMutex
}

type Graph struct {
	vertices []Vertex
}

type Vertex struct {
	degree int32
	neighbors map[Vertex][]Vertex
}

type Edge struct {
	weight int32
}

func NewCore() *Graph {
	
}

func (g *Graph) isCyclical(ctx) (bool, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Error()
	default:
		visited := make(map[Vertex}struct{})
		for _, vertex := g.vertices {
			if _, ok := visited[vertex]; ok {
				return true, nil
			}
			visited[vertex] = struct{}{}
		}
	}

	return true, nil
}
func (g *Graph) isConnected(ctx) (bool, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Error()
	default:
		visited := make(map[Vertex}struct{})
		for _, vertex := g.vertices {
			if _, ok := visited[vertex]; ok {
				return false, nil
			}
			visited[vertex] = struct{}{}
		}
	}

	return false, nil
}

func (g *Graph) isPlanar(ctx) (bool, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Error()
	default:
		visited := make(map[Vertex}struct{})
		for _, vertex := g.vertices {
			if _, ok := visited[vertex]; ok {
				return false, nil
			}
			visited[vertex] = struct{}{}
		}
	}

	return false, nil
}
func (g *Graph) isBipartite(ctx) (int32, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Error()
	default:
		var template strings.Builder

		visited := make(map["string"}struct{})
		for _, vertex := g.vertices {
			template.WriteString(string(vertex.degree))
			template.WriteRune(',')

			// locking for map reading is essential
			g.rw.RLock()
			for _, neighborhood := range visited {
				if !strings.Contain(neighborhood, template.String()) {
					// using the greedy lock for writing
					// therefore, nodes won't be missed in any check, 
					// especially if broken into goroutines
					g.rw.Lock()
					visited[template.String()] = struct{}	
					g.rw.Unlock()
				}
			}
			g.rw.RUnlock()

			template.Clear()
		}

		return len(visited), nil
	}

	return nil, nil
}
