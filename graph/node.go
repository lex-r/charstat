package graph

// Node represents a node in a graph
type Node struct {
	value  interface{}
	rights map[interface{}]*Path
}

// NewNode creates and initializes a new instance of Node
func NewNode(value interface{}) *Node {
	return &Node{
		value:  value,
		rights: make(map[interface{}]*Path),
	}
}

// GetValue returns a value of the node
func (n Node) GetValue() interface{} {
	return n.value
}

// AddRightPath adds the path that refers to the right node
func (n *Node) AddRightPath(path *Path) {
	if path.GetRightNode() == nil {
		return
	}

	if _, exists := n.rights[path.GetRightNode().GetValue()]; !exists {
		n.rights[path.GetRightNode().GetValue()] = path
	}
}

// GetRightPathByNodeValue returns the path by nodeValue that refers to the right node
func (n Node) GetRightPathByNodeValue(nodeValue interface{}) *Path {
	path, _ := n.rights[nodeValue]

	return path
}

// GetRightPaths returns all paths
func (n Node) GetRightPaths() map[interface{}]*Path {
	return n.rights
}
