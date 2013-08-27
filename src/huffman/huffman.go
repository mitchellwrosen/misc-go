package main

//package huffman

import (
	"fmt"
	"sort"
)

type Pair struct {
	key byte
	val int
}

type PairSlice []Pair

// sort.Interface implementation.
func (p PairSlice) Len() int {
	return len(p)
}
func (p PairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PairSlice) Less(i, j int) bool {
	return p[i].val < p[j].val
}

func sortMapByValue(m map[byte]int) PairSlice {
	p := make(PairSlice, 0, len(m))

	for k, v := range m {
		p = append(p, Pair{k, v})
	}

	sort.Sort(p)

	return p
}

func Encode(data []byte) []byte {
	freq := make(map[byte]int)

	for _, b := range data {
		freq[b]++
	}

	sortedFreq := sortMapByValue(freq)

	for _, f := range sortedFreq {
		fmt.Printf("%c: %d\n", f.key, f.val)
	}

	return data
}

func main() {
	data := []byte{'a', 'a', 'b', 'b', 'b', 'c'}

	Encode(data)
}
