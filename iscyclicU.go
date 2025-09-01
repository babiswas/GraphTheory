package main

import (
	"fmt"
	"graph/graphalgo"
)

func IscyclicU() {
	u := graphalgo.Undirected{}
	u.UndirectedInit(4)
	u.AddEdges(0, 2)
	u.AddEdges(2, 0)
	u.AddEdges(0, 1)
	u.AddEdges(1, 2)
	u.AddEdges(2, 3)
	u.AddEdges(3, 3)
	fmt.Println("Undirected graph is cyclic or not.")
	fmt.Println(u.Iscyclic())
}
