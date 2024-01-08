package main

import (
	"fmt"

 	"github.com/deliveranceTechSolutions/graphe/graph"
)

func main() {
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now())
	// defer cancel()
	g := graph.NewCore()
	
	g.Add(12, 12, graph.ThankYou, graph.Book)
	g.Add(1, 2, graph.Book)
	g.Execute()

	fmt.Printf("%#v", g)
}
