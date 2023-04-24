Create a new file named `README.md` in the root folder of your `graphmatrix5xn` package, and add the following content:

```markdown
# GraphMatrix5xN

GraphMatrix5xN is an efficient algorithm for finding the shortest path between two nodes in a graph using a 5xN matrix structure. This algorithm was designed to solve a wide range of problems including those involving predicates and artificial intelligence.

## Features

- Efficient algorithm for finding the shortest path in a graph
- Utilizes a 5xN matrix structure for improved memory usage compared to Dijkstra's algorithm
- Can handle graphs with negative edge weights
- Ideal for problems involving predicates and artificial intelligence

## Installation

```bash
go get github.com/infosave2007/graphmatrixgo5xn
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/infosave2007/graphmatrixgo5xn/src"
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
```
