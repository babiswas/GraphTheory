package main

import (
	"fmt"
	"graph/graphalgo"
)

func DFSTraversal() {
	fmt.Println("DFS traversal of directed graph.")
	dfsObj := graphalgo.DFS{}
	dfsObj.DFSinit(4)
	dfsObj.AddEdges(0, 2)
	dfsObj.AddEdges(2, 0)
	dfsObj.AddEdges(0, 1)
	dfsObj.AddEdges(1, 2)
	dfsObj.AddEdges(2, 3)
	dfsObj.AddEdges(3, 3)
	fmt.Println("The dfs traversal is.")
	dfsObj.DFS()
}

func DFSStackTraversal() {
	dfs_stack := graphalgo.DFSStack{}
	dfs_stack.DFSStackInit(4)
	dfs_stack.Addedges(0, 2)
	dfs_stack.Addedges(2, 0)
	dfs_stack.Addedges(0, 1)
	dfs_stack.Addedges(1, 2)
	dfs_stack.Addedges(2, 3)
	dfs_stack.Addedges(3, 3)
	fmt.Println("The dfs traversal using stack:")
	dfs_stack.DFS(2)
}
