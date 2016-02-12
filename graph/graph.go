package graph

// Graph represents a weighted graph
type Graph struct {
	nodes map[interface{}]*Node
}

// NewGraph creates and initializes a new instance of Graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[interface{}]*Node),
	}
}

// GetNode returns a node by nodeValue if exists
func (g Graph) GetNode(nodeValue interface{}) (*Node, bool) {
	node, exists := g.nodes[nodeValue]
	return node, exists
}

// GetNodes returns a nodes from the graph
func (g Graph) GetNodes() map[interface{}]*Node {
	return g.nodes
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(node *Node) bool {
	if _, exists := g.nodes[node.GetValue()]; exists {
		return false
	}

	g.nodes[node.GetValue()] = node
	return true
}
