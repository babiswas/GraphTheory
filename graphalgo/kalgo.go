package graphalgo

import (
	"container/list"
	"fmt"
)

type Tsort struct {
	vertex int
	adj    map[int]*list.List
}

func (t *Tsort) Topsortinit(vertex int) {
	t.vertex = vertex
	t.adj = make(map[int]*list.List)
	for i := 0; i < t.vertex; i++ {
		t.adj[i] = list.New()
	}
}

func (t *Tsort) AddEdges(i, j int) {
	_, ok := t.adj[i]
	if !ok {
		t.adj[i] = list.New()
	}
	t.adj[i].PushBack(j)
}

func (t *Tsort) KAlgo() {
	count := 0
	indegree := make([]int, t.vertex)
	for j := 0; j < t.vertex; j++ {
		lst := t.adj[j]
		for elm := lst.Front(); elm != nil; elm = elm.Next() {
			index := elm.Value.(int)
			indegree[index] += 1
		}
	}
	queue := list.New()
	for i := 0; i < t.vertex; i++ {
		if indegree[i] == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() > 0 {
		elm := queue.Front()
		num := elm.Value.(int)
		fmt.Println(num)
		lst := t.adj[num]
		queue.Remove(elm)
		for el := lst.Front(); el != nil; el = el.Next() {
			indegree[el.Value.(int)] -= 1
			if indegree[el.Value.(int)] == 0 {
				queue.PushBack(el.Value.(int))
			}
		}
		count += 1
	}
	if count == t.vertex {
		fmt.Println("Topological sorting is possible.")
	} else {
		fmt.Println("There is a cycle in the graph.")
	}

}
