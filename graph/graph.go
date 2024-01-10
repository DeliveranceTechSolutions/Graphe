package graph

import (
	"context"
	"sync"
	"time"
)

type IGraph interface {
	IsCyclical(context.Context) (bool, error)
	IsConnected(context.Context,int32,int32) (bool, error)
	IsPlanar(context.Context) (bool, error)
	IsBipartite(context.Context) (int32, error)
	Append(context.Context,int32,int32,chan Vertex,...EmailVertex)
	Execute(context.Context,chan Vertex)
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
	edges []int32
	subgraph map[EmailVertex]struct{}
}

// EVH EmailVertexHandler is an important concept when these events
// begin to be broken off into it's own environment
// this will take in the selection, but also the
// email metadata (i.e. sender name, title, other people mentioned, etc.)
type EmailVertex int8
type EmailVertexHandler struct {
	selection map[EmailVertex]struct{}
}

// Any request coming in that is most nessecary to execute first
// needs to maintain the bottom position
const (
	ThankYou EmailVertex = iota + 1
	Prayer
	Book
)

func (ev EmailVertex) Value() int8 {
	return int8(ev)
}

// Heaviest will determine what process will manage the email creation
// eventually the idea is to create DAG edges with the heaviest weight.
// Therefore, a k8s cluster can handle Book, ThankYou, etc. essentially
// future load-balancing within the system's requests. 
func (evh EmailVertexHandler) Heaviest() (maxOpt EmailVertex) {
	var max int8
	for emailVert := range evh.selection {
		if v := emailVert.Value(); max < v {
			max = v
			maxOpt = emailVert
		}
	}

	return 
}

type any interface{}

// conversely, this switch statement will be opposite order, to process the heavier execution first
func (g *graph) Execute(ctx context.Context, loadCh chan Vertex) {
	var evh EmailVertexHandler
	
	for {
		select {
		case <-ctx.Done():	
			return 
		case selection := <-loadCh:
			evh.selection = selection.subgraph
			switch evh.Heaviest() {
			case Book:
				// fmt.Println("Book")				
			case Prayer:
				// fmt.Println("Prayer")				
			case ThankYou:
				// fmt.Println("ThankYou")				
			}
		default:
			time.Sleep(time.Second * 5)
		}
	}
}

func NewCore() IGraph {
	var rw sync.RWMutex
	return &graph{
		rw: &rw,
		instance: &Graph{
			vertices: make([]Vertex, 0),
		},
	}
}

func (g *graph) Append(
	ctx context.Context, 
	degree int32, 
	edge int32, 
	loadCh chan Vertex, 
	tasks ...EmailVertex,
) {
	edges := make([]int32, 0)
	length := len(g.instance.vertices)
	subgraph := make(map[EmailVertex]struct{})
	vertex := Vertex{
		degree: degree,
		edges: edges,
		subgraph: subgraph,
	}
	
	for _, vertice := range g.instance.vertices {
		if vertice.degree == degree {
			for _, task := range tasks {
				g.rw.RLock()
				if _, ok := vertice.subgraph[task]; !ok {
				g.rw.RUnlock()
				g.rw.Lock()

					vertice.subgraph[task] = struct{}{}

				g.rw.Unlock()
				g.rw.RLock()
				}
				g.rw.RUnlock()
			}
				
			loadCh<-vertice	
			return
		}
	}
	
	// default case, else on a map lookup is bad form, but might be better syntax
	if length == len(g.instance.vertices) {
		g.rw.Lock()
		for _, task := range tasks {
			vertex.subgraph[task] = struct{}{}
		}
		g.instance.vertices = append(g.instance.vertices, vertex)
		g.rw.Unlock()
	}
	
	loadCh<-vertex
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
