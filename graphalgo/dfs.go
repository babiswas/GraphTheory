package graphalgo

import (
	"container/list"
	"fmt"
)

type DFS struct {
	vertex int
	adj    map[int]*list.List
}

func (d *DFS) DFSinit(vertex int) {
	d.vertex = vertex
	d.adj = make(map[int]*list.List)
	for i := 0; i < d.vertex; i++ {
		d.adj[i] = list.New()
	}
}

func (d *DFS) AddEdges(i, j int) {
	_, ok := d.adj[i]
	if !ok {
		d.adj[i] = list.New()
	}
	d.adj[i].PushBack(j)
}

func (d *DFS) DFSUtil(visited []bool, u int) {
	visited[u] = true
	fmt.Println(u)
	lst := d.adj[u]
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		index := elm.Value.(int)
		if !visited[index] {
			d.DFSUtil(visited, index)
		}
	}
}

func (d *DFS) DFS() {
	visited := make([]bool, d.vertex)
	for i := 0; i < d.vertex; i++ {
		visited[i] = false
	}
	for j := 0; j < d.vertex; j++ {
		if !visited[j] {
			d.DFSUtil(visited, j)
		}
	}
}
