package graphalgo

import "container/list"

type MotherVertex struct {
	vertex int
	adj    map[int]*list.List
}

func (m *MotherVertex) MotherVertexInit(vertex int) {
	m.vertex = vertex
	m.adj = make(map[int]*list.List)
	for i := 0; i < m.vertex; i++ {
		m.adj[i] = list.New()
	}
}

func (m *MotherVertex) Addedges(i, j int) {
	_, ok := m.adj[i]
	if !ok {
		m.adj[i] = list.New()
	}
	m.adj[i].PushBack(j)
}

func (m *MotherVertex) MVertexUtil(visited []bool, u int) {
	visited[u] = true
	lst := m.adj[u]
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		index := elm.Value.(int)
		if !visited[index] {
			m.MVertexUtil(visited, index)
		}
	}
}

func (m *MotherVertex) FindMotherVertex() int {
	visited := make([]bool, m.vertex)
	mvertex := -1
	for i := 0; i < m.vertex; i++ {
		visited[i] = false
	}
	for i := 0; i < m.vertex; i++ {
		if !visited[i] {
			m.MVertexUtil(visited, i)
			mvertex = i
		}
	}
	for i := 0; i < m.vertex; i++ {
		visited[i] = false
	}
	m.MVertexUtil(visited, mvertex)
	for i := 0; i < m.vertex; i++ {
		if !visited[i] {
			return -1
		}
	}
	return mvertex
}
