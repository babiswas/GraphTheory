package graphalgo

import (
	"container/list"
	"fmt"
)

type BFS struct {
	adj    map[int]*list.List
	vertex int
}

func (b *BFS) BFSinit(vertex int) {
	b.adj = make(map[int]*list.List)
	b.vertex = vertex
	for i := 0; i < vertex; i++ {
		b.adj[i] = list.New()
	}
}

func (b *BFS) AddEdges(i, j int) {
	_, ok := b.adj[i]
	if !ok {
		b.adj[i] = list.New()
	}
	b.adj[i].PushBack(j)
}

func (b *BFS) BFSTraversal(u int) {
	visited := make([]bool, b.vertex)
	for i := 0; i < b.vertex; i++ {
		visited[i] = false
	}
	queue := list.New()
	queue.PushBack(u)
	visited[u] = true
	for queue.Len() > 0 {
		element := queue.Front()
		index := element.Value.(int)
		fmt.Println(index)
		queue.Remove(element)
		lst := b.adj[index]
		for num := lst.Front(); num != nil; num = num.Next() {
			if !visited[num.Value.(int)] {
				queue.PushBack(num.Value.(int))
				visited[num.Value.(int)] = true
			}
		}
	}
}
