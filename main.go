package main

import (
	"fmt"
	"sync"
	"time"

 	"github.com/deliveranceTechSolutions/graphe/graph"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now())
	// defer cancel()
	g := graph.NewCore()

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	testConst := int32(100000)
	testGroups := int(testConst)
	wg.Add(1)
	go func() {
		wg.Add(testGroups)
		for i := int32(0); i < int32(testConst); i++ {
			go func(i int32, g graph.IGraph, wg *sync.WaitGroup) {
				fmt.Println(i)
				g.Add(i, i, wg, graph.ThankYou, graph.Book)
			}(i, g, &wg)
		}
	}()
	wg.Done()

	wg.Add(1)
	go func() {
		wg.Add(testGroups)
		for i := int32(0); i < int32(testConst); i++ {
			go g.Execute(&wg)
		}
	}()
	wg.Done()
	return	
}
