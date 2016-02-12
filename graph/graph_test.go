package graph

import "testing"

func TestAddNode(t *testing.T) {
	graph := NewGraph()

	n := NewNode("n")
	graph.AddNode(n)
	actualNode, exists := graph.GetNode("n")

	if !exists {
		t.Errorf("Node not found")
	}

	if actualNode.GetValue() != "n" {
		t.Errorf("Received node value not equals to added node value")
	}
}
