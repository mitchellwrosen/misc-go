package kruskals

import (
	"container/unionfind"
	"fmt"
	"graph"
	"sort"
)

// TODO: move this to graph?
type edgeSlice []graph.Edge

func (e edgeSlice) Len() int {
	return len(e)
}

func (e edgeSlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e edgeSlice) Less(i, j int) bool {
	return e[i].Weight < e[j].Weight
}

// O(m log n) implemenation.
func Impl1(g *graph.Graph) (graph.Mst, error) {
	ret := g.NewMst()
	sort.Sort(edgeSlice(g.Edges))

	uf := unionfind.New(unionfind.NodesToEmptyInterfaces(g.Nodes))

	// |g.Nodes|-1 edges in a MST
	cur := 0
	for i := 0; i < len(g.Nodes)-1; i++ {
		edge := g.Edges[cur]
		l1 := uf.Find(edge.Node1)
		l2 := uf.Find(edge.Node2)
		for l1 == l2 {
			cur++
			edge = g.Edges[cur]
			l1 = uf.Find(edge.Node1)
			l2 = uf.Find(edge.Node2)
		}

		ret = append(ret, edge)
		uf.Union(l1, l2)

		cur++
	}

	return ret, nil
}

// Max-spacing k-clustering implementation.
func MaxSpacingKClustering(g *graph.Graph, k int) ([]graph.Edge,
	*graph.Edge, error) {

	if k > len(g.Nodes) {
		return nil, nil, fmt.Errorf("%d clusters requested of graph with %d "+
			"nodes", k, len(g.Nodes))
	}

	edges := g.NewMst()
	sort.Sort(edgeSlice(g.Edges))

	uf := unionfind.New(unionfind.NodesToEmptyInterfaces(g.Nodes))

	// Add edges until there are k clusters.
	cur := 0
	for i := 0; i < len(g.Nodes)-k; i++ {
		edge := g.Edges[cur]
		l1 := uf.Find(edge.Node1).(graph.Node)
		l2 := uf.Find(edge.Node2).(graph.Node)
		for l1 == l2 {
			cur++
			edge = g.Edges[cur]
			l1 = uf.Find(edge.Node1).(graph.Node)
			l2 = uf.Find(edge.Node2).(graph.Node)
		}

		edges = append(edges, edge)
		uf.Union(l1, l2)

		cur++
	}

	if k != 1 {
		edge := g.Edges[cur]
		l1 := uf.Find(edge.Node1).(graph.Node)
		l2 := uf.Find(edge.Node2).(graph.Node)
		for l1 == l2 {
			cur++
			edge = g.Edges[cur]
			l1 = uf.Find(edge.Node1).(graph.Node)
			l2 = uf.Find(edge.Node2).(graph.Node)
		}

		return edges, &edge, nil
	}

	return edges, nil, nil
}
