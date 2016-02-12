package graph

import "testing"

func TestGetValue(t *testing.T) {
	node := NewNode("n")

	if node.GetValue() != "n" {
		t.Errorf("Returned value is wrong")
	}
}

func TestGetRightPathByNodeValue(t *testing.T) {
	node := NewNode("n")
	rightNode := NewNode("l")
	path := NewPath(1, nil, rightNode)

	node.AddRightPath(path)

	rightPath := node.GetRightPathByNodeValue(rightNode.GetValue())

	if rightPath != path {
		t.Errorf("GetRightPathByNodeValue returned a wrong path")
	}
}

func TestGetRightPaths(t *testing.T) {
	node := NewNode("n")
	rightNode := NewNode("r")
	path := NewPath(1, nil, rightNode)

	node.AddRightPath(path)

	rightPaths := node.GetRightPaths()

	actualRightPath, exists := rightPaths[rightNode.GetValue()]
	if !exists {
		t.Errorf("GetRightPaths returned paths without added path")
	}

	if actualRightPath != path {
		t.Errorf("GetRightPaths returned a wrong path")
	}
}
