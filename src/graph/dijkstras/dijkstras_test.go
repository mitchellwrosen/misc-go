package dijkstras

import (
	"mytesting"
	"os"
	"testing"
)

func TestNewGraph(te *testing.T) {
	t := mytesting.T{te}

	file, err := os.Open("tests/simple1.txt")
	if err != nil {
		t.Fatal(err)
	}

	graph, err := NewGraph(file)
	if err != nil {
		t.Fatal(err)
	}

	t.AssertEq(4, len(graph.Nodes))

	t.ExpectEq(0, graph.Nodes[0].Id)
	t.AssertEq(3, len(graph.Nodes[0].Edges))

	t.ExpectEq(1, graph.Nodes[0].Edges[0].To)
	t.ExpectEq(10, graph.Nodes[0].Edges[0].Weight)
	t.ExpectEq(2, graph.Nodes[0].Edges[1].To)
	t.ExpectEq(10, graph.Nodes[0].Edges[1].Weight)
	t.ExpectEq(3, graph.Nodes[0].Edges[2].To)
	t.ExpectEq(10, graph.Nodes[0].Edges[2].Weight)

	t.ExpectEq(0, graph.Nodes[1].Edges[0].To)
	t.ExpectEq(10, graph.Nodes[1].Edges[0].Weight)
	t.ExpectEq(2, graph.Nodes[1].Edges[1].To)
	t.ExpectEq(10, graph.Nodes[1].Edges[1].Weight)
	t.ExpectEq(3, graph.Nodes[1].Edges[2].To)
	t.ExpectEq(10, graph.Nodes[1].Edges[2].Weight)

	t.ExpectEq(0, graph.Nodes[2].Edges[0].To)
	t.ExpectEq(10, graph.Nodes[2].Edges[0].Weight)
	t.ExpectEq(1, graph.Nodes[2].Edges[1].To)
	t.ExpectEq(10, graph.Nodes[2].Edges[1].Weight)
	t.ExpectEq(3, graph.Nodes[2].Edges[2].To)
	t.ExpectEq(10, graph.Nodes[2].Edges[2].Weight)

	t.ExpectEq(0, graph.Nodes[3].Edges[0].To)
	t.ExpectEq(10, graph.Nodes[3].Edges[0].Weight)
	t.ExpectEq(1, graph.Nodes[3].Edges[1].To)
	t.ExpectEq(10, graph.Nodes[3].Edges[1].Weight)
	t.ExpectEq(2, graph.Nodes[3].Edges[2].To)
	t.ExpectEq(10, graph.Nodes[3].Edges[2].Weight)
}
