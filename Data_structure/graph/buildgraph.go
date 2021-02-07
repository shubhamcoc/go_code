package main

import (
	"fmt"
)

// Graph structure represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key int
	neighbour []*Vertex
}

// addVertex adds a Vertex to the Graph
func (g *Graph) addVertex(val int) {
	if contains(g.vertices, val) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", val)
		fmt.Println(err.Error())
	}else{
		g.vertices = append(g.vertices, &Vertex{key: val})
	}
}

// addEdge will connect 2 vertex
func (g *Graph) addEdge(from, to int){
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid Edge (%v--->%v)", from, to)
		fmt.Println(err.Error())
	}else if contains(fromVertex.neighbour, to){
		err := fmt.Errorf("Already Exists (%v--->%v)", from, to)
		fmt.Println(err.Error())
	}else{
		// add edge
		fromVertex.neighbour = append(fromVertex.neighbour, toVertex)
	}
}

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains will match if any vertex already exist in the graph
func contains(g []*Vertex, k int) bool {
	for _, ele := range g{
		if k == ele.key{
			return true
		}
	}
	return false
}

// print will print all the vertices of graph
func (g *Graph) print(){
	for _, ele := range g.vertices{
		fmt.Printf("\nVertex %v :", ele.key)
		for _, v := range ele.neighbour{
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}


func main() {
	test := &Graph{}
	
	// add vertex to the graph
	for i := 0; i < 5; i++ {
		test.addVertex(i)
	}
	
	test.addEdge(0, 4)
	test.addEdge(6, 4)
	
	test.print()
}
