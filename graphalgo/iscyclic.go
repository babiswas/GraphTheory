package graphalgo

import "container/list"

type DirectedCyclic struct {
	vertex int
	adj    map[int]*list.List
}

func (d *DirectedCyclic) DirectedCyclicInit(vertex int) {
	d.vertex = vertex
	d.adj = make(map[int]*list.List)
	for i := 0; i < d.vertex; i++ {
		d.adj[i] = list.New()
	}
}

func (d *DirectedCyclic) Addedges(i, j int) {
	_, ok := d.adj[i]
	if !ok {
		d.adj[i] = list.New()
	}
	d.adj[i].PushBack(j)
}

func (d *DirectedCyclic) IsCyclicUtil(visited []bool, u int, recstack []bool) bool {
	visited[u] = true
	recstack[u] = true
	lst := d.adj[u]
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		index := elm.Value.(int)
		if !visited[index] {
			if d.IsCyclicUtil(visited, index, recstack) {
				return true
			}
		} else if recstack[index] {
			return true
		}
	}
	recstack[u] = false
	return false
}

func (d *DirectedCyclic) IsCyclic() bool {
	visited := make([]bool, d.vertex)
	recstack := make([]bool, d.vertex)
	for i := 0; i < d.vertex; i++ {
		visited[i] = false
		recstack[i] = false
	}
	for j := 0; j < d.vertex; j++ {
		if !visited[j] {
			if d.IsCyclicUtil(visited, j, recstack) {
				return true
			}
		}
	}
	return false
}
