package main

import (
	"fmt"
	"github.com/lex-r/charstat/graph"
	"os"
	"strings"
)

func main() {

	text := os.Args[1:]

	g := graph.NewGraph()

	for _, word := range text {
		if len(word) < 2 {
			continue
		}

		word = strings.ToLower(word)

		var leftChar int32

		for index, runeValue := range word {
			if index == 0 {
				leftChar = runeValue
				continue
			}
			leftLetter := leftChar
			rightLetter := runeValue

			leftNode := createNode(g, leftLetter)
			rightNode := createNode(g, rightLetter)

			rightPath := leftNode.GetRightPathByNodeValue(rightNode.GetValue())
			if rightPath == nil {
				rightPath = graph.NewPath(1, leftNode, rightNode)
				leftNode.AddRightPath(rightPath)
			} else {
				rightPath.IncreaseWeight(1)
			}

			leftChar = runeValue
		}
	}

	fmt.Print(
		"digraph finite_state_machine {\n" +
			"\trankdir=LR;\n" +
			"\tnode [shape = circle];\n",
	)

	for _, n := range g.GetNodes() {
		for _, rightPath := range n.GetRightPaths() {
			fmt.Printf("\t\"%c\" -> \"%c\" [ label = \"%d\" ];\n", n.GetValue(), rightPath.GetRightNode().GetValue(), rightPath.GetWeight())
		}
	}

	fmt.Print("}")
}

func createNode(g *graph.Graph, value interface{}) *graph.Node {
	node, exists := g.GetNode(value)
	if !exists {
		node = graph.NewNode(value)
		g.AddNode(node)
	}

	return node
}
