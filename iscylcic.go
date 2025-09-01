package main

import (
	"fmt"
	"graph/graphalgo"
)

func Iscyclic() {
	d := graphalgo.DirectedCyclic{}
	d.DirectedCyclicInit(4)
	d.Addedges(0, 2)
	d.Addedges(2, 0)
	d.Addedges(0, 1)
	d.Addedges(1, 2)
	d.Addedges(2, 3)
	d.Addedges(3, 3)
	fmt.Println("Detect if the directed graph has cycle.")
	fmt.Println(d.IsCyclic())
}
