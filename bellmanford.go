package main

import "graph/graphalgo"

func BellmanFord() {
	graph := graphalgo.BellmanFord{}
	graph.BellmanFordInit(5)
	graph.AddEdges(0, 1, -1)
	graph.AddEdges(0, 2, 4)
	graph.AddEdges(1, 2, 3)
	graph.AddEdges(1, 3, 2)
	graph.AddEdges(1, 4, 2)
	graph.AddEdges(3, 2, 5)
	graph.AddEdges(3, 1, 1)
	graph.AddEdges(4, 3, -3)
	graph.BellmanFordExec(0)
}
