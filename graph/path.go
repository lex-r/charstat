package graph

// Path represents path from the left node to right node
type Path struct {
	weight int64
	left   *Node
	right  *Node
}

// NewPath creates and initializes a new instance of Path
func NewPath(weight int64, leftNode, rightNode *Node) *Path {
	return &Path{
		weight: weight,
		left:   leftNode,
		right:  rightNode,
	}
}

// GetWeight returns a weight of path
func (p Path) GetWeight() int64 {
	return p.weight
}

// IncreaseWeight increases weight of path
func (p *Path) IncreaseWeight(weight int64) {
	p.weight += weight
}

// GetLeftNode returns a reference to left node
func (p Path) GetLeftNode() *Node {
	return p.left
}

// GetRightNode return a reference to right node
func (p Path) GetRightNode() *Node {
	return p.right
}
