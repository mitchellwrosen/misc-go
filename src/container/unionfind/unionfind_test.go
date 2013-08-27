package unionfind

import (
	"mytesting"
	"testing"
)

func TestUnionFind(te *testing.T) {
	t := mytesting.T{te}

	data := [...]int{1, 2, 3, 4, 5}

	data2 := make([]interface{}, len(data))
	for i, v := range data {
		data2[i] = v
	}
	uf := New(data2)

	// Initially, each node's leader is itself.
	for i := 0; i < 5; i++ {
		t.ExpectEq(data[i], uf.Find(i+1))
	}

	uf.Union(data[0], data[1])
	t.ExpectEq(data[0], uf.Find(1))
	t.ExpectEq(data[0], uf.Find(2))

	uf.Union(data[0], data[2])
	t.ExpectEq(data[0], uf.Find(3))
}
