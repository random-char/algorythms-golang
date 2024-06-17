package graph

import (
	"algorythms-golang/pkg/adt/queue"
)

type Graph struct {
	Vertices map[int]*Vertex
}

type Edge struct {
	Weight     int
	DestVertex *Vertex
}

type Vertex struct {
	Value int
	Edges map[int]*Edge
}

func (g *Graph) AddVertex(key, value int) {
	g.Vertices[key] = &Vertex{Value: value, Edges: map[int]*Edge{}}
}

func (g *Graph) AddEdge(sourceKey, destKey, weight int) {
	if _, ok := g.Vertices[sourceKey]; !ok {
		return
	}

	if _, ok := g.Vertices[destKey]; !ok {
		return
	}

	g.Vertices[sourceKey].Edges[destKey] = &Edge{
		DestVertex: g.Vertices[destKey],
		Weight:     weight,
	}
}

func (g *Graph) NeighbourKeys(sourceKey int) []int {
	keys := make([]int, 0)

	vertex, ok := g.Vertices[sourceKey]
	if !ok {
		return keys
	}

	for destKey := range vertex.Edges {
		keys = append(keys, destKey)
	}

	return keys
}

func (g *Graph) BreadthFirstSearch(startKey, searchValue int) (int, bool) {
	visited := map[int]bool{}
	q := queue.NewDequeArray[int]()
	q.PushBack(startKey)
	visited[startKey] = true

	for !q.Empty() {
		checking, _ := q.PopFront()
		neighbourKeys := g.NeighbourKeys(checking)
		for _, neighbourKey := range neighbourKeys {
			if g.Vertices[neighbourKey].Value == searchValue {
				return neighbourKey, true
			}

			if v, ok := visited[neighbourKey]; ok && v {
				continue
			}

			visited[neighbourKey] = true
			q.PushBack(neighbourKey)
		}
	}

	return 0, false
}

func (g *Graph) DepthFirstSearch(startKey, searchValue int) (int, bool) {
	discovered := map[int]bool{}

	return g.depthFirstSearchRecursive(startKey, searchValue, &discovered)
}

func (g *Graph) depthFirstSearchRecursive(currentKey, searchValue int, discovered *map[int]bool) (int, bool) {
	currentVertex, ok := g.Vertices[currentKey]
	if !ok {
		return 0, false
	}

	if currentVertex.Value == searchValue {
		return currentKey, true
	}

	(*discovered)[currentKey] = true
	for _, neighbourKey := range g.NeighbourKeys(currentKey) {
		if _, ok := (*discovered)[neighbourKey]; !ok {
			foundKey, found := g.depthFirstSearchRecursive(neighbourKey, searchValue, discovered)
			if found {
				return foundKey, found
			}
		}
	}

	return 0, false
}

type GraphOption func(g *Graph)

func NewGraph(opts ...GraphOption) *Graph {
	g := &Graph{
		Vertices: make(map[int]*Vertex),
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func WithAdjecencyList(list map[int][]int) GraphOption {
	return func(g *Graph) {
		for vertex, edges := range list {
			if _, ok := g.Vertices[vertex]; !ok {
				g.AddVertex(vertex, vertex)
			}

			//edge ~ destination vertex
			for _, edge := range edges {
				if _, ok := g.Vertices[edge]; !ok {
					g.AddVertex(edge, edge)
				}

				g.AddEdge(vertex, edge, 0)
			}
		}
	}
}
