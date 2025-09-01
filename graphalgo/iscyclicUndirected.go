package graphalgo

import "container/list"

type Undirected struct {
	vertex int
	adj    map[int]*list.List
}

func (u *Undirected) UndirectedInit(vertex int) {
	u.vertex = vertex
	u.adj = make(map[int]*list.List)
	for i := 0; i < u.vertex; i++ {
		u.adj[i] = list.New()
	}
}

func (u *Undirected) AddEdges(i, j int) {
	_, ok := u.adj[i]
	if !ok {
		u.adj[i] = list.New()
	}
	u.adj[i].PushBack(j)
	_, ok = u.adj[j]
	if !ok {
		u.adj[j] = list.New()
	}
	u.adj[j].PushBack(i)
}

func (u *Undirected) Iscyclic() bool {
	visited := make([]bool, u.vertex)
	for i := 0; i < u.vertex; i++ {
		visited[i] = false
	}
	for k := 0; k < u.vertex; k++ {
		if !visited[k] {
			if u.IscyclicU(visited, k, -1) {
				return true
			}
		}
	}
	return false
}

func (u *Undirected) IscyclicU(visited []bool, v int, parent int) bool {
	visited[v] = true
	lst := u.adj[v]
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		index := elm.Value.(int)
		if !visited[index] {
			if u.IscyclicU(visited, index, v) {
				return true
			}
		} else if index != parent {
			return true
		}
	}
	return false
}
