package main

import (
	"fmt"
	"graph/graphalgo"
)

func MotherVetex() {
	m := graphalgo.MotherVertex{}
	m.MotherVertexInit(4)
	m.Addedges(0, 2)
	m.Addedges(2, 0)
	m.Addedges(1, 2)
	m.Addedges(0, 1)
	m.Addedges(2, 3)
	m.Addedges(3, 3)
	fmt.Println("The mother vertex of the graph is:")
	num := m.FindMotherVertex()
	fmt.Println(num)
}
