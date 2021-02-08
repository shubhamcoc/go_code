package main

import (
	"fmt"
)

var unique map[int]int = make(map[int]int)

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

// BFS implementation
func (g *Graph) bfs(root int) []int {
	q := &Queue{}
	q.enqueue(root)
	var res []int
	for i := root; i >= 0; {
		for _, ele := range g.vertices{
			if ele.key == i {
				for _, v := range ele.neighbour{
					q.enqueue(v.key)
				}
			}else{
				continue
			}
		}
		if i == root {
			i = q.dequeue()
			res = append(res, i)
		}
		i = q.dequeue()
		if i > -1 {
			res = append(res, i)
		}
		
	}
	return res
}

// Queue structure
type Queue struct {
	queue []int
}

// enqueue
func (q *Queue) enqueue(key int) {
	_, ok := unique[key]
	if !ok {
		unique[key] = 0
		q.queue = append(q.queue, key)
	}
}

// dequeue
func (q *Queue) dequeue() int {
	if len(q.queue) == 0 {
		err := fmt.Errorf("Queue is empty")
		fmt.Println(err.Error())
		return -1
	}
		
	ele := q.queue[0]
	q.queue = q.queue[1:]
	return ele
}

func main() {
	test := &Graph{}
	
	// add vertex to the graph
	for i := 0; i < 7; i++ {
		test.addVertex(i)
	}
	
	test.addEdge(0, 1)
	test.addEdge(0, 3)
	test.addEdge(6, 1)
	test.addEdge(6, 4)
	test.addEdge(1, 0)
	test.addEdge(1, 2)
	test.addEdge(1, 3)
	test.addEdge(1, 5)
	test.addEdge(1, 6)
	test.addEdge(2, 1)
	test.addEdge(2, 3)
	test.addEdge(2, 4)
	test.addEdge(2, 5)
	test.addEdge(3, 0)
	test.addEdge(3, 1)
	test.addEdge(3, 2)
	test.addEdge(3, 4)
	test.addEdge(4, 2)
	test.addEdge(4, 3)
	test.addEdge(4, 6)
	test.addEdge(5, 1)
	test.addEdge(5, 2)
	
	test.print()
	result := test.bfs(0)
	fmt.Println("BFS for the given graph test is:", result)
	
}
