package main

import (
	"fmt"
	"sync"
	"time"

 	"github.com/deliveranceTechSolutions/graphe/graph"
)

func main() {
	var wg sync.WaitGroup
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now())
	// defer cancel()
	g := graph.NewCore()
	start := time.Now()
	wg.Add(1)
	go func() {
		wg.Add(10000)
		for i := int32(0); i < 10000; i++ {
			go func(i int32, g graph.IGraph, wg *sync.WaitGroup) {
				fmt.Println(i)
				g.Add(i, i, wg, graph.ThankYou, graph.Book)
			}(i, g, &wg)
		}
	wg.Done()
	}()
	wg.Add(1)
	go func() {
		wg.Add(10000)
		for i := int32(0); i < 10000; i++ {
			go g.Execute(&wg)
		}
	wg.Done()
	}()
	
	wg.Wait()

	fmt.Println(time.Since(start))
	return	
}
