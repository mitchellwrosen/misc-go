package kruskals

import (
	"graph"
	"os"
	"testing"
)

func BenchmarkKruskals(b *testing.B) {
	b.StopTimer()
	file, err := os.Open("../driver/clustering1.txt")
	if err != nil {
		b.Fatal(err)
	}

	g, err := graph.New(file)
	if err != nil {
		b.Fatal(err)
	}

	b.StartTimer()
	for i := 0; i < 10; i++ {
		Impl1(g)
	}
}
