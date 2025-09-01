package graphalgo

import (
	"container/list"
	"fmt"
)

type Kahns struct {
	vertex int
	adj    map[int]*list.List
}

func (k *Kahns) KahnsInit(vertex int) {
	k.vertex = vertex
	k.adj = make(map[int]*list.List)
	for i := 0; i < k.vertex; i++ {
		k.adj[i] = list.New()
	}
}

func (k *Kahns) Addedges(i, j int) {
	_, ok := k.adj[i]
	if !ok {
		k.adj[i] = list.New()
	}
	k.adj[i].PushBack(j)
}

func (k *Kahns) KahnsUtil(visited []bool, u int, stack *list.List) {
	visited[u] = true
	lst := k.adj[u]
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		index := elm.Value.(int)
		if !visited[index] {
			k.KahnsUtil(visited, index, stack)
		}
	}
	stack.PushFront(u)
}

func (k *Kahns) KahnsAlgo() {
	stack := list.New()
	visited := make([]bool, k.vertex)
	for i := 0; i < k.vertex; i++ {
		visited[i] = false
	}
	for i := 0; i < k.vertex; i++ {
		if !visited[i] {
			k.KahnsUtil(visited, i, stack)
		}
	}
	for elm := stack.Front(); elm != nil; elm = elm.Next() {
		fmt.Println(elm.Value.(int))
	}
}
