package main

import (
	"fmt"
	"graph/graphalgo"
)

func BFSGraphTraversal() {
	bfs := graphalgo.BFS{}
	bfs.BFSinit(4)
	bfs.AddEdges(0, 2)
	bfs.AddEdges(2, 0)
	bfs.AddEdges(1, 2)
	bfs.AddEdges(2, 3)
	bfs.AddEdges(3, 3)
	bfs.AddEdges(0, 1)
	fmt.Println("The bfs traversal of the graph is:")
	bfs.BFSTraversal(2)
}
