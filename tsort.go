package main

import (
	"fmt"
	"graph/graphalgo"
)

func Topsort() {
	kh := graphalgo.Kahns{}
	kh.KahnsInit(6)
	kh.Addedges(5, 0)
	kh.Addedges(4, 0)
	kh.Addedges(5, 2)
	kh.Addedges(2, 3)
	kh.Addedges(3, 1)
	kh.Addedges(4, 1)
	fmt.Println("Executing kahns algo.")
	kh.KahnsAlgo()
}
