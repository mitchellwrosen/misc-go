package prims

import (
	"container/heap"
	"fmt"
	"graph"
	"math"
)

// O(mn) implemenation.
func Impl1(g *graph.Graph) (graph.Mst, error) {
	ret := g.NewMst()

	// Spanned set initialized with one node.
	X := graph.NewNodeSet()
	X.Add(g.Nodes[0])

	// Un-spanned set initialized with the rest.
	Y := graph.NewNodeSet()
	for i := 1; i < len(g.Nodes); i++ {
		Y.Add(g.Nodes[i])
	}

	// At each iteration, add the minimum edge that crosses the spanned and
	// unspanned sets. Continue until all nodes are spanned.
	for len(Y) != 0 {
		minEdge := graph.Edge{-1, -1, math.MaxInt32}
		for _, edge := range g.Edges {
			if edge.Spans(X, Y) && edge.Weight < minEdge.Weight {
				minEdge = edge
			}
		}

		if minEdge.Weight == math.MaxInt32 {
			return nil, fmt.Errorf("No edge spanned %s and %s", X, Y)
		}

		ret = append(ret, minEdge)

		// Remove from Y and add to X.
		if Y[minEdge.Node1] {
			Y.Remove(minEdge.Node1)
			X.Add(minEdge.Node1)
		} else {
			Y.Remove(minEdge.Node2)
			X.Add(minEdge.Node2)
		}
	}

	return ret, nil
}

type edgeHeap []graph.Edge

// sort.Interface implementation.
func (e edgeHeap) Len() int {
	return len(e)
}

func (e edgeHeap) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e edgeHeap) Less(i, j int) bool {
	return e[i].Weight < e[j].Weight
}

// heap.Interface implementation
func (e *edgeHeap) Push(x interface{}) {
	n := len(*e)
	(*e) = (*e)[0 : n+1]
	(*e)[n] = x.(graph.Edge)
}

func (e *edgeHeap) Pop() interface{} {
	n := len(*e)
	ret := (*e)[n-1]
	(*e) = (*e)[0 : n-1]
	return ret
}

// O(m log n) implemenation.
func Impl2(g *graph.Graph) (graph.Mst, error) {
	ret := g.NewMst()

	// Spanned set initialized with one node.
	X := graph.NewNodeSet()
	X.Add(g.Nodes[0])

	// Un-spanned set initialized with the rest.
	Y := graph.NewNodeSet()
	for i := 1; i < len(g.Nodes); i++ {
		Y.Add(g.Nodes[i])
	}

	// Push onto a heap (EdgeHeap) all the edges.
	edgeHeap := make(edgeHeap, 0, len(g.Edges))
	for _, edge := range g.Edges {
		heap.Push(&edgeHeap, edge)
	}

	// At each iteration, add the minimum edge that crosses the spanned and
	// unspanned sets. Continue until all nodes are spanned.
	for len(Y) != 0 {
		// The min edge that spans X and Y is of course not necessarily the min
		// edge overall. Thus we must save off edges that we pop, so we can push
		// them back on once we find the first edge that spans X and Y.
		tempEdges := make([]graph.Edge, 0)
		minEdge := heap.Pop(&edgeHeap).(graph.Edge)
		for !minEdge.Spans(X, Y) {
			//fmt.Printf("%s doesn't span %s and %s\n", minEdge, X, Y)
			tempEdges = append(tempEdges, minEdge)
			minEdge = heap.Pop(&edgeHeap).(graph.Edge)
		}

		for _, edge := range tempEdges {
			heap.Push(&edgeHeap, edge)
		}

		//fmt.Printf("%s spans %s and %s\n", minEdge, X, Y)
		ret = append(ret, minEdge)

		// Remove from Y and add to X.
		if Y[minEdge.Node1] {
			Y.Remove(minEdge.Node1)
			X.Add(minEdge.Node1)
		} else {
			Y.Remove(minEdge.Node2)
			X.Add(minEdge.Node2)
		}
	}

	return ret, nil

}

// A graph.Node, but with extra information for this heap implementation.
type heapNode struct {
	node   graph.Node
	edge   graph.Edge // The cheapest edge to the set of spanned nodes.
	weight int
}

type nodeHeap []heapNode

// sort.Interface implementation.
func (n nodeHeap) Len() int {
	return len(n)
}

func (n nodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nodeHeap) Less(i, j int) bool {
	return n[i].weight < n[j].weight
}

// heap.Interface implementation
func (n *nodeHeap) Push(x interface{}) {
	l := len(*n)
	(*n) = (*n)[0 : l+1]
	(*n)[l] = x.(heapNode)
}

func (n *nodeHeap) Pop() interface{} {
	l := len(*n)
	ret := (*n)[l-1]
	(*n) = (*n)[0 : l-1]
	return ret
}

func (h nodeHeap) Update(n heapNode) {

}

// O(m log n) implemenation.
/*
func prims3(g *graph.Graph) (graph.Mst, error) {
	ret := g.NewMst()

	// Spanned set initialized with one node.
	X := graph.NewNodeSet()
	X.Add(g.Nodes[0])

	// Un-spanned set initialized with the rest.
	Y := graph.NewNodeSet()
	for i := 1; i < len(g.Nodes); i++ {
		Y.Add(g.Nodes[i])
	}

	// Push onto a heap (NodeHeap) all the Nodes in Y, with values equal to
	// their minimum (only) edge to X.
	nodeHeap := make(NodeHeap, 0, len(Y))
	for _, edge := range g.EdgesIncidentTo(g.Nodes[0]) {
		if Y[edge.Node1] {
			heap.Push(&nodeHeap,
				HeapNode{node: edge.Node1, weight: edge.Weight})
		} else {
			heap.Push(&nodeHeap,
				HeapNode{node: edge.Node2, weight: edge.Weight})
		}
	}

	// At each iteration, add the minimum edge that crosses the spanned and
	// unspanned sets. Continue until all nodes are spanned.
	for len(Y) != 0 {
		// Pop the min node, and add its min edge that spans X and Y.
		node := heap.Pop(&nodeHeap).(HeapNode)
		ret = append(ret, node.edge)

		// Remove from Y and add to X.
		Y.Remove(node.node)
		X.Add(node.node)

		// Nodes adjacent to the node just added to |X| that are still in |Y|
		// possibly now have a cheaper edge from |Y| to |X| (via the node just
		// added).
		for _, adjNode := range g.NodesAdjacentTo(node.node) {
			if Y[adjNode] {

				nodeHeap.Update(adjNode)
			}
		}

		tempEdges := make(EdgeHeap, 0)
		minEdge := heap.Pop(&edgeHeap).(graph.Edge)
		for !minEdge.Spans(X, Y) {
			//fmt.Printf("%s doesn't span %s and %s\n", minEdge, X, Y)
			tempEdges = append(tempEdges, minEdge)
			minEdge = heap.Pop(&edgeHeap).(graph.Edge)
		}

		for _, edge := range tempEdges {
			heap.Push(&edgeHeap, edge)
		}

		//fmt.Printf("%s spans %s and %s\n", minEdge, X, Y)
		ret = append(ret, minEdge)

	}

	return ret, nil

}
*/
