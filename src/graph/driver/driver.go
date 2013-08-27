package main

import (
	"fmt"
	"graph"
	"graph/kruskals"
	"graph/prims"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g, err := graph.New(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Prims")
	mst, err := g.GetMst(prims.Impl1)
	if err != nil {
		panic(err)
	}
	fmt.Println(mst.Weight())

	fmt.Println("Prims2")
	mst, err = g.GetMst(prims.Impl2)
	if err != nil {
		panic(err)
	}
	fmt.Println(mst.Weight())

	fmt.Println("Kruskals")
	mst, _ = g.GetMst(kruskals.Impl1)
	if err != nil {
		panic(err)
	}
	fmt.Println(mst.Weight())

	_, nextEdge, err := kruskals.MaxSpacingKClustering(g, 4)
	if err != nil {
		panic(err)
	}

	if nextEdge != nil {
		fmt.Printf("Max spacing 4-clustering: %d\n", nextEdge.Weight)
	}
}
