package graphalgo

import (
	"container/list"
	"fmt"
)

type TransitiveClosure struct {
	vertex int
	adj    map[int]*list.List
}

func (t *TransitiveClosure) TransitiveClosureInit(vertex int) {
	t.vertex = vertex
	t.adj = make(map[int]*list.List)
	for i := 0; i < t.vertex; i++ {
		t.adj[i] = list.New()
	}
}

func (t *TransitiveClosure) AddEdges(i, j int) {
	_, ok := t.adj[i]
	if !ok {
		t.adj[i] = list.New()
	}
	t.adj[i].PushBack(j)
}

func (t *TransitiveClosure) TClosureUtil(M [][]int, u, v int) {
	M[u][v] = 1
	lst := t.adj[v]
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		index := elm.Value.(int)
		if M[u][index] != 1 {
			t.TClosureUtil(M, u, index)
		}
	}
}

func (t *TransitiveClosure) TClosure() {
	M := make([][]int, t.vertex)
	for i := 0; i < t.vertex; i++ {
		M[i] = make([]int, t.vertex)
	}
	for i := 0; i < t.vertex; i++ {
		t.TClosureUtil(M, i, i)
	}
	fmt.Println(M)
}
