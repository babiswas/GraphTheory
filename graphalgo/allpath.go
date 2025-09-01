package graphalgo

import (
	"container/list"
	"fmt"
)

type AllPath struct {
	vertex  int
	adj     map[int]*list.List
	allPath [][]int
}

func (a *AllPath) AllPathInit(vertex int) {
	a.vertex = vertex
	a.adj = make(map[int]*list.List)
	for i := 0; i < a.vertex; i++ {
		a.adj[i] = list.New()
	}
}

func (a *AllPath) AddEdges(i, j int) {
	_, ok := a.adj[i]
	if !ok {
		a.adj[i] = list.New()
	}
	a.adj[i].PushBack(j)
}

func (a *AllPath) AllPathUtil(visited []bool, u int, temp []int, v int) {
	visited[u] = true
	temp = append(temp, u)
	if u == v {
		newtemp := make([]int, len(temp))
		copy(newtemp, temp)
		a.allPath = append(a.allPath, newtemp)
	} else {
		lst := a.adj[u]
		for elm := lst.Front(); elm != nil; elm = elm.Next() {
			index := elm.Value.(int)
			if !visited[index] {
				a.AllPathUtil(visited, index, temp, v)
			}
		}

	}

	temp = temp[:len(temp)-1]
	visited[u] = false
	_ = temp

}

func (a *AllPath) FindAllPath(u, v int) {
	temp := make([]int, 0)
	a.allPath = make([][]int, 0)
	visited := make([]bool, a.vertex)
	for i := 0; i < a.vertex; i++ {
		visited[i] = false
	}
	a.AllPathUtil(visited, u, temp, v)
	fmt.Println(a.allPath)
}
