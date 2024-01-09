package main

import (
	"fmt"
	"sync"
	"time"

 	"github.com/deliveranceTechSolutions/graphe/graph"
)

func driver() bool {
	var load sync.WaitGroup
	var execute sync.WaitGroup
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now())
	// defer cancel()
	g := graph.NewCore()
	testConst := int32(1000)
	testGroups := int(testConst)
	start := time.Now()
	defer func() {
		fmt.Println("before end")
		fmt.Println(time.Since(start))
	}()

	load.Add(testGroups + 1)
	go func() {
	defer load.Done()
		for i := int32(0); i < int32(testConst); i++ {
			go func(i int32, g graph.IGraph) {
			defer load.Done()
				g.Add(i, i, graph.ThankYou, graph.Book)
			}(i, g)
		}
		
	}()
	load.Wait()

	execute.Add(testGroups + 1)
	go func() {
	defer execute.Done()
		for i := int32(0); i < int32(testConst); i++ {
			go g.Execute(&execute)
		}
	}()
	execute.Wait()	

	return true
}

func main() {
	if driver() {
		fmt.Println("finished and returned")
	}
}
