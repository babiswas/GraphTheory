package graphalgo

import (
	"container/list"
	"fmt"
)

type ShortestPath struct {
	adjList map[int]*list.List
	vertex  int
}

func (sp *ShortestPath) ShortestPathInit(vertex int) {
	sp.vertex = vertex
	sp.adjList = make(map[int]*list.List)
	for i := 0; i < vertex; i++ {
		sp.adjList[i] = list.New()
	}
}

func (sp *ShortestPath) AddEdges(i, j int) {
	_, ok := sp.adjList[i]
	if !ok {
		sp.adjList[i] = list.New()
	}
	sp.adjList[i].PushBack(j)
	_, ok = sp.adjList[j]
	if !ok {
		sp.adjList[j] = list.New()
	}
	sp.adjList[j].PushBack(i)
}

func (sp *ShortestPath) FindShortestPath(source, target int) {
	prev := sp.BFS(source, target)
	fmt.Println(prev)
	path := make([]int, 0)
	path = append(path, target)
	index := prev[target]
	for index != -1 {
		path = append(path, index)
		index = prev[index]
	}
	start, end := 0, len(path)-1
	for start < end {
		path[start], path[end] = path[end], path[start]
		start += 1
		end -= 1
	}
	if path[0] == source && path[len(path)-1] == target {
		fmt.Println("Shortest path from source to destination is:")
		fmt.Println(path)
	} else {
		fmt.Println("No valid path from:", source, " to target", target)
	}
}

func (sp *ShortestPath) BFS(source, target int) []int {
	prev := make([]int, sp.vertex)
	for i := 0; i < sp.vertex; i++ {
		prev[i] = -1
	}
	visited := make([]bool, sp.vertex)
	for i := 0; i < sp.vertex; i++ {
		visited[i] = false
	}
	queue := list.New()
	queue.PushBack(source)
	visited[source] = true
	for queue.Len() != 0 {
		elm := queue.Front()
		index := elm.Value.(int)
		fmt.Println(index)
		queue.Remove(elm)
		lst := sp.adjList[index]
		for el := lst.Front(); el != nil; el = el.Next() {
			if !visited[el.Value.(int)] {
				queue.PushBack(el.Value.(int))
				visited[el.Value.(int)] = true
				prev[el.Value.(int)] = index
			}
		}
	}
	return prev
}
