package main

import "graph/graphalgo"

func ShortestPath() {
	sp := graphalgo.ShortestPath{}
	sp.ShortestPathInit(8)
	sp.AddEdges(1, 2)
	sp.AddEdges(1, 0)
	sp.AddEdges(0, 3)
	sp.AddEdges(3, 7)
	sp.AddEdges(3, 4)
	sp.AddEdges(7, 4)
	sp.AddEdges(7, 6)
	sp.AddEdges(6, 4)
	sp.AddEdges(6, 5)
	sp.AddEdges(4, 5)
	sp.FindShortestPath(0, 7)
}
