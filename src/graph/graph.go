package graph

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Graph struct {
	Nodes []Node
	Edges []Edge
}

func (g *Graph) NewMst() Mst {
	return make(Mst, 0, len(g.Nodes)-1)
}

func (g *Graph) GetMst(alg MstAlgorithm) (Mst, error) {
	return alg(g)
}

func (g *Graph) NodesAdjacentTo(n Node) []Node {
	nodes := make([]Node, 0)

	for _, edge := range g.Edges {
		if n == edge.Node1 {
			nodes = append(nodes, edge.Node2)
		} else if n == edge.Node2 {
			nodes = append(nodes, edge.Node1)
		}
	}

	return nodes
}

func (g *Graph) EdgesIncidentTo(n Node) []Edge {
	edges := make([]Edge, 0)

	for _, edge := range g.Edges {
		if n == edge.Node1 || n == edge.Node2 {
			edges = append(edges, edge)
		}
	}

	return edges
}

type Node int
type NodeSet map[Node]bool

func NewNodeSet() NodeSet {
	return make(NodeSet)
}

func (s NodeSet) Add(n Node) {
	s[n] = true
}

func (s NodeSet) Remove(n Node) {
	delete(s, n)
}

func (s NodeSet) String() string {
	b := new(bytes.Buffer)

	for n := range s {
		b.WriteString(fmt.Sprintf("%d, ", n))
	}

	return b.String()
}

type Edge struct {
	Node1, Node2 Node
	Weight       int
}

func (e Edge) String() string {
	return fmt.Sprintf("%d-%d (%d)", e.Node1, e.Node2, e.Weight)
}

func (e *Edge) Spans(s1, s2 NodeSet) bool {
	return s1[e.Node1] && s2[e.Node2] || s1[e.Node2] && s2[e.Node1]
}

func New(r io.Reader) (*Graph, error) {
	br := bufio.NewReader(r)
	line, err := br.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var numNodes, numEdges int
	fmt.Sscanf(line, "%d %d", &numNodes, &numEdges)

	graph := &Graph{make([]Node, numNodes), make([]Edge, numEdges)}
	for i := 0; i < numNodes; i++ {
		graph.Nodes[i] = Node(i + 1)
	}

	for i := 0; i < numEdges; i++ {
		line, err := br.ReadString('\n')
		if err != nil {
			return nil, err
		}

		var node1, node2 Node
		var weight int
		fmt.Sscanf(line, "%d %d %d", &node1, &node2, &weight)

		graph.Edges[i] = Edge{node1, node2, weight}
	}

	return graph, nil
}

type Mst []Edge

func (m Mst) Weight() (weight int) {
	for _, edge := range m {
		weight += edge.Weight
	}

	return
}

type MstAlgorithm func(*Graph) (Mst, error)
