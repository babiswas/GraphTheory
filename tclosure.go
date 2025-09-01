package main

import (
	"fmt"
	"graph/graphalgo"
)

func Tclosure() {
	t := graphalgo.TransitiveClosure{}
	t.TransitiveClosureInit(4)
	t.AddEdges(0, 2)
	t.AddEdges(2, 0)
	t.AddEdges(0, 1)
	t.AddEdges(1, 2)
	t.AddEdges(2, 3)
	t.AddEdges(3, 3)
	fmt.Println("The transitive closure of the graph is:")
	t.TClosure()
}
