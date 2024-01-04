package main

import (
	"context"
	"fmt"
	"time"


 	"github.com/deliveranceTechSolutions/graphe/graph"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Second*30)
	defer cancel()

	g := graph.NewCore()
}
