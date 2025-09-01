package main

import (
	"flag"
	"fmt"
)

func main() {
	ops := flag.String("op", "bfs", "graph traversal option.")
	flag.Parse()
	graph_operation := *ops
	switch graph_operation {
	case "bfs":
		fmt.Println("The bfs traversal of the graph is:")
		BFSGraphTraversal()
	case "dfs":
		fmt.Println("The dfs traversal of the graph is:")
		DFSTraversal()
	case "dfsstack":
		fmt.Println("The dfs traversal using stack is:")
		DFSStackTraversal()
	case "mvertex":
		fmt.Println("Displaying mother vertex.")
		MotherVetex()
	case "topsort":
		fmt.Println("Executing topological sorting.")
		Topsort()
	case "kalgo":
		fmt.Println("Executing kahns algo.")
		KahnsAlgo()
	case "iscle":
		fmt.Println("Verify if the directed graph is cyclic.")
		Iscyclic()
	case "iscleu":
		fmt.Println("Is the undirected graph cyclic.")
		IscyclicU()
	case "tclsre":
		fmt.Println("The transitive closure of the directed graph is:")
		Tclosure()
	case "apth":
		fmt.Println("All path from source to destination.")
		AllPath()
	case "shortp":
		fmt.Println("The shortest path from source to destination:")
		ShortestPath()
	case "bllman":
		fmt.Println("Single source shortest path:")
		BellmanFord()

	}
}
