package main

import "graph/graphalgo"

func AllPath() {
	a := graphalgo.AllPath{}
	a.AllPathInit(4)
	a.AddEdges(0, 2)
	a.AddEdges(2, 0)
	a.AddEdges(0, 1)
	a.AddEdges(1, 2)
	a.AddEdges(1, 3)
	a.AddEdges(2, 3)
	a.AddEdges(3, 3)
	a.FindAllPath(2, 3)
}
