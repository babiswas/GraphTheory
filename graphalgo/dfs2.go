package graphalgo

import (
	"container/list"
	"fmt"
)

type DFSStack struct {
	vertex int
	adj    map[int]*list.List
}

func (d *DFSStack) DFSStackInit(vertex int) {
	d.vertex = vertex
	d.adj = make(map[int]*list.List)
	for i := 0; i < d.vertex; i++ {
		d.adj[i] = list.New()
	}
}

func (d *DFSStack) Addedges(i, j int) {
	_, ok := d.adj[i]
	if !ok {
		d.adj[i] = list.New()
	}
	d.adj[i].PushBack(j)
}

func (d *DFSStack) DFS(u int) {
	stack := list.New()
	visited := make([]bool, d.vertex)
	for i := 0; i < d.vertex; i++ {
		visited[i] = false
	}
	stack.PushFront(u)
	for stack.Len() > 0 {
		elm := stack.Front()
		index := elm.Value.(int)
		if !visited[index] {
			fmt.Println(index)
			visited[index] = true
		}
		stack.Remove(elm)
		lst := d.adj[index]
		for el := lst.Front(); el != nil; el = el.Next() {
			index = el.Value.(int)
			if !visited[index] {
				stack.PushFront(index)
			}
		}
	}
}
