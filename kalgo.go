package main

import (
	"fmt"
	"graph/graphalgo"
)

func KahnsAlgo() {
	kh := graphalgo.Tsort{}
	kh.Topsortinit(6)
	kh.AddEdges(5, 0)
	kh.AddEdges(4, 0)
	kh.AddEdges(5, 2)
	kh.AddEdges(2, 3)
	kh.AddEdges(3, 1)
	kh.AddEdges(4, 1)
	fmt.Println("Executing kahns algorithm:")
	kh.KAlgo()
}
