package graphalgo

import (
	"fmt"
	"math"
)

type Weight struct {
	source int
	dest   int
	wght   int
}

type BellmanFord struct {
	weights []Weight
	vertex  int
}

func (bl *BellmanFord) BellmanFordInit(vertex int) {
	bl.vertex = vertex
	bl.weights = make([]Weight, 0)
}

func (bl *BellmanFord) AddEdges(u, v, w int) {
	wght := Weight{}
	wght.source = u
	wght.dest = v
	wght.wght = w
	bl.weights = append(bl.weights, wght)
}

func (bl *BellmanFord) BellmanFordExec(source int) {
	dst := make([]float64, bl.vertex)
	for i := 0; i < bl.vertex; i++ {
		dst[i] = math.Inf(1)
	}
	dst[source] = float64(0)
	fmt.Println(dst)
	fmt.Println(bl.weights)
	for i := 0; i < bl.vertex-1; i++ {
		for j := 0; j < len(bl.weights); j++ {
			source := bl.weights[j].source
			dest := bl.weights[j].dest
			wght := bl.weights[j].wght
			if dst[source] < math.Inf(1) && (dst[source]+float64(wght)) < dst[dest] {
				dst[dest] = dst[source] + float64(wght)
			}
		}
	}

	fmt.Println(dst)

	for l := 0; l < len(bl.weights); l++ {
		source := bl.weights[l].source
		dest := bl.weights[l].dest
		wght := bl.weights[l].wght
		if dst[source]+float64(wght) < dst[dest] {
			fmt.Println("There is a cycle in the graph:")
		}
	}

	for n := 0; n < bl.vertex; n++ {
		fmt.Println("index:", n, " cost:", dst[n])
	}
}
