package main

import "fmt"

type Graph struct {
	adjacentList  map[int][]int
	numberOfNodes int
}

func (g *Graph) addVertex(node int) {
	if _, ok := g.adjacentList[node]; ok {
		return
	}
	if g.numberOfNodes == 0 {
		g.adjacentList = make(map[int][]int)

	}
	g.adjacentList[node] = []int{}
	g.numberOfNodes++
}

// addEdge - adds connections between nodes
func (g *Graph) addEdge(firstnode, secondnode int) {
	g.adjacentList[firstnode] = append(g.adjacentList[firstnode], secondnode)
	g.adjacentList[secondnode] = append(g.adjacentList[secondnode], firstnode)
}

func main() {
	graph := Graph{}
	graph.addVertex(0)
	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(3)
	graph.addVertex(4)
	graph.addVertex(5)
	graph.addVertex(6)
	graph.addEdge(3, 1)
	graph.addEdge(3, 4)
	graph.addEdge(4, 2)
	graph.addEdge(4, 5)
	graph.addEdge(1, 2)
	graph.addEdge(1, 0)
	graph.addEdge(0, 2)
	graph.addEdge(6, 5)
	fmt.Println(graph)
}
