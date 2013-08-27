package unionfind

import (
	"graph"
)

type UnionFind struct {
	leaderMap   leaderMap
	childrenMap childrenMap
}

type leaderMap map[interface{}]interface{}     // Maps data to leader
type childrenMap map[interface{}][]interface{} // Maps leader to children

func NodesToEmptyInterfaces(nodes []graph.Node) []interface{} {
	ret := make([]interface{}, len(nodes))
	for i, v := range nodes {
		ret[i] = v
	}

	return ret
}

func New(data []interface{}) *UnionFind {
	uf := new(UnionFind)
	uf.leaderMap = make(leaderMap)
	uf.childrenMap = make(childrenMap)

	for i := 0; i < len(data); i++ {
		d := data[i]

		children := make([]interface{}, 0, 1)
		children = append(children, d)

		uf.leaderMap[d] = d
		uf.childrenMap[d] = children
	}

	return uf
}

func (uf *UnionFind) Union(a1, a2 interface{}) {
	var smaller, larger interface{}
	if len(uf.childrenMap[a1]) < len(uf.childrenMap[a2]) {
		smaller, larger = a1, a2
	} else {
		smaller, larger = a2, a1
	}

	for _, data := range uf.childrenMap[smaller] {
		uf.leaderMap[data] = larger
		uf.childrenMap[larger] = append(uf.childrenMap[larger], data)
	}

	delete(uf.childrenMap, smaller)
}

func (uf *UnionFind) Find(data interface{}) interface{} {
	return uf.leaderMap[data]
}

func (uf *UnionFind) Count(p interface{}) int {
	return len(uf.childrenMap[p])
}
