package adt_test

import (
	"algorythms-golang/pkg/adt/graph"
	"fmt"
	"testing"
)

func TestGraphSearch(t *testing.T) {
	g := graph.NewGraph(graph.WithAdjecencyList(map[int][]int{
		0: {4, 2, 1},
		1: {},
		2: {3},
		4: {5},
		5: {},
	}))

	fmt.Println(g.BreadthFirstSearch(0, 3))
	fmt.Println(g.DepthFirstSearch(0, 3))
	fmt.Println(g.DepthFirstSearch(0, 6))
}
