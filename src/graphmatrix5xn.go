package graphmatrix5xn

import (
	"errors"
	"fmt"
	"math"
)

// Node represents a node in the graph.
type Node struct {
	ID string
}

// Edge represents a directed edge between two nodes in the graph.
type Edge struct {
	From   *Node
	To     *Node
	Weight int
}

// Graph represents a graph using a 5xN matrix structure.
type Graph struct {
	Nodes map[string]*Node
	Edges []*Edge
}

// NewGraph creates a new Graph instance.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: []*Edge{},
	}
}

// AddNode adds a new node with the given ID to the graph.
func (g *Graph) AddNode(id string) {
	g.Nodes[id] = &Node{ID: id}
}

// AddEdge adds a directed edge from the node with the 'from' ID to the node with the 'to' ID, and assigns the given weight to it.
func (g *Graph) AddEdge(from string, to string, weight int) {
	fromNode, ok1 := g.Nodes[from]
	toNode, ok2 := g.Nodes[to]
	if !ok1 || !ok2 {
		return
	}
	edge := &Edge{From: fromNode, To: toNode, Weight: weight}
	g.Edges = append(g.Edges, edge)
}

// FindShortestPath finds the shortest path between two nodes with the given 'start' and 'end' IDs.
func (g *Graph) FindShortestPath(startNodeID, endNodeID string) ([]string, int, error) {
	startNode, ok := g.Nodes[startNodeID]
	if !ok {
		return nil, 0, fmt.Errorf("start node not found: %s", startNodeID)
	}

	endNode, ok := g.Nodes[endNodeID]
	if !ok {
		return nil, 0, fmt.Errorf("end node not found: %s", endNodeID)
	}

	// Create the 5xN matrix
	matrix := make([][]int, 5)
	for i := range matrix {
		matrix[i] = make([]int, len(g.Edges))
	}

	nodeIndex := make(map[string]int)
	indexNode := make(map[int]string)
	index := 0
	for _, edge := range g.Edges {
		if _, ok := nodeIndex[edge.From.ID]; !ok {
			nodeIndex[edge.From.ID] = index
			indexNode[index] = edge.From.ID
			index++
		}
		if _, ok := nodeIndex[edge.To.ID]; !ok {
			nodeIndex[edge.To.ID] = index
			indexNode[index] = edge.To.ID
			index++
		}
	}

	// Initialize the matrix
	for i, edge := range g.Edges {
		matrix[0][i] = nodeIndex[edge.From.ID]
		matrix[1][i] = nodeIndex[edge.To.ID]
		matrix[2][i] = edge.Weight
		matrix[3][i] = math.MaxInt32
		matrix[4][i] = 0
	}

	matrix[3][nodeIndex[startNode.ID]] = 0

	// Apply the algorithm
	for {
		changed := false
		for i, edge := range g.Edges {
			if matrix[4][i] == 0 && matrix[3][matrix[0][i]] != math.MaxInt32 && matrix[3][matrix[0][i]]+edge.Weight < matrix[3][matrix[1][i]] {
				matrix[3][matrix[1][i]] = matrix[3][matrix[0][i]] + edge.Weight
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	// Find the shortest path and its distance
	minDistance := math.MaxInt32
	var minDistanceIndex int
	for i := 0; i < len(g.Edges); i++ {
		if matrix[1][i] == nodeIndex[endNode.ID] && matrix[3][i] < minDistance {
			minDistance = matrix[3][i]
			minDistanceIndex = i
		}
	}

	if minDistance == math.MaxInt32 {
		return nil, 0, errors.New("No path found")
	}

	// Build the path by going backward from the end node
	path := []string{endNode.ID}
	currentNodeID := matrix[0][minDistanceIndex]
	for currentNodeID != nodeIndex[startNode.ID] {
		for i := 0; i < len(g.Edges); i++ {
			if matrix[1][i] == currentNodeID && matrix[4][i] == 1 {
				path = append([]string{indexNode[matrix[0][i]]}, path...)
				currentNodeID = matrix[0][i]
				break
			}
		}
	}

	return path, minDistance, nil
}
