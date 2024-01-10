package main

import (
	"context"
	"fmt"
	"time"

 	"github.com/deliveranceTechSolutions/graphe/graph"
)

func driver() bool {
	// var load sync.WaitGroup
	// var execute sync.WaitGroup

	loadCh := make(chan graph.Vertex, 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 30)
	defer cancel()

	g := graph.NewCore()
	testConst := int32(3000)
//	testGroups := int(testConst)
	start := time.Now()
	defer func() {
		fmt.Println("before end")
		fmt.Println(time.Since(start))
	}()

	for i := int32(0); i < int32(testConst); i++ {
		weight := int32(graph.ThankYou.Value() + graph.Book.Value())
		go g.Append(ctx, i, weight, loadCh, graph.ThankYou, graph.Book)
	}
	
	for i := int32(0); i < int32(testConst); i++ {
		fmt.Println(i)
		go g.Execute(ctx, loadCh)
	}

	return true
}

func main() {
	if driver() {
		fmt.Println("finished and returned")
	}
}
