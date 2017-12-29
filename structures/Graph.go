package structures

import (
	"errors"
	"fmt"
)

type Vertex struct {
	Value int
	VisitingState int
	Neighbors []*Vertex
}

type Graph struct {
	Vertexes []*Vertex
}

func NewGraph() *Graph {
	return &Graph{make([]*Vertex, 0)}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.Vertexes = append(g.Vertexes, v)
}

func (g *Graph) RemoveVertex(vertexIndex int) error {

	vertex := g.Vertexes[vertexIndex]

	if vertexIndex >= len(g.Vertexes) || vertexIndex < 0 {
		return errors.New("There is no such vertex in Graph")
	}


	g.Vertexes = append(g.Vertexes[0 : vertexIndex], g.Vertexes[vertexIndex : len(g.Vertexes)]...)

	for i := 0; i < len(g.Vertexes); i++ {
		for j := 0; j < len(g.Vertexes[i].Neighbors); j++ {
			if g.Vertexes[i].Neighbors[j] == vertex {
				g.Vertexes =
					append(g.Vertexes[i].Neighbors[0 : j],
					g.Vertexes[i].Neighbors[j + 1 : len(g.Vertexes[i].Neighbors)]...)
			}
		}
	}

	return nil
}

func (g *Graph) BreadthFirstSearch(startVertex int) {
	currentVertex := g.Vertexes[startVertex]
	queue := NewQueue()
	queue.Enqueue(currentVertex)
	currentVertex.VisitingState = 1

	for !queue.IsEmpty {
		currentVertex, _ = queue.Dequeue()
		fmt.Println(currentVertex.Value)

		for i := 0; i < len(currentVertex.Neighbors); i++ {
			if currentVertex.Neighbors[i].VisitingState > 0 {
				continue
			}

			queue.Enqueue(currentVertex.Neighbors[i])
			currentVertex.Neighbors[i].VisitingState = 1
		}
	}
}

func (g *Graph) DepthFirstSearch(startVertex int) {
	currentVertex := g.Vertexes[startVertex]
	stack := NewStack()
	stack.Push(currentVertex)
	currentVertex.VisitingState = 1

	for !stack.IsEmpty {
		currentVertex, _ = stack.Pop()
		fmt.Println(currentVertex.Value)

		for i := 0; i < len(currentVertex.Neighbors); i++ {
			if currentVertex.Neighbors[i].VisitingState > 0 {
				continue
			}

			stack.Push(currentVertex.Neighbors[i])
			currentVertex.Neighbors[i].VisitingState = 1
		}
	}

}