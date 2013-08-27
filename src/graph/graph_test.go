package graph

import (
	"container/heap"
	"mytesting"
	"testing"
)

func TestEdgeHeap(te *testing.T) {
	t := mytesting.T{te}

	edgeHeap := make(EdgeSlice, 0, 5)
	for i := 0; i < 5; i++ {
		heap.Push(&edgeHeap, Edge{0, 0, i})
	}

	for i := 0; i < 5; i++ {
		top := heap.Pop(&edgeHeap).(Edge)
		t.ExpectEq(i, top.Weight)
	}
}
