package graph

import "testing"

func TestGetWeight(t *testing.T) {
	path := NewPath(0, nil, nil)
	path.IncreaseWeight(2)

	if path.GetWeight() != 2 {
		t.Errorf("Incorrect returned weight")
	}
}

func TestGetLeftNode(t *testing.T) {
	node := NewNode("n")

	path := NewPath(0, node, nil)

	if node != path.GetLeftNode() {
		t.Errorf("The returned left node is wrong")
	}
}

func TestGetRightNode(t *testing.T) {
	node := NewNode("n")

	path := NewPath(0, nil, node)

	if node != path.GetRightNode() {
		t.Errorf("The returned right node is wrong")
	}
}
