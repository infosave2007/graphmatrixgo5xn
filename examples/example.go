package main

import (
	"fmt"

	"github.com/infosave2007/graphmatrixgo5xn/src/graphmatrix5xn"
)

func main() {
	// Create a new graph instance
	graph := graphmatrix5xn.NewGraph()

	// Add nodes to the graph
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")

	// Add edges to the graph
	graph.AddEdge("A", "B", 1)
	graph.AddEdge("B", "C", 2)
	graph.AddEdge("A", "D", 4)
	graph.AddEdge("D", "E", 5)
	graph.AddEdge("B", "E", 6)

	// Find the shortest path from A to E
	path, distance, err := graph.FindShortestPath("A", "E")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Printf("Shortest path: %v\n", path)
	fmt.Printf("Distance: %d\n", distance)
}
