package main

import (
	"bufio"
	"fmt"
	"os"
)

type knapsack struct {
	capacity int
}

type item struct {
	weight int
	value  int
}

func max(i, j int64) int64 {
	if i > j {
		return i
	}

	return j
}

func printValues(v [][]int) {
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[0]); j++ {
			fmt.Printf("%d ", v[i][j])
		}
		fmt.Println()
	}
}

func pack(k knapsack, items []item) (int64, []item) {
	values := make([][]int64, len(items))
	for i := 0; i < len(items); i++ {
		values[i] = make([]int64, k.capacity+1)
	}

	// Initialize row 1
	for i := 0; i < k.capacity+1; i++ {
		if items[0].weight > i {
			values[0][i] = 0
		} else {
			values[0][i] = int64(items[0].weight)
		}
	}

	for i := 1; i < len(items); i++ {
		for j := 1; j <= k.capacity; j++ {
			// This sub-solution can't build upon the case that item i is in
			// the knapsack if item i is too heavy to fit in by itself.
			if j-items[i].weight >= 0 {
				values[i][j] = max(values[i-1][j],
					values[i-1][j-items[i].weight]+int64(items[i].value))
			} else {
				values[i][j] = values[i-1][j]
			}
		}
	}

	// Reconstruct the items selected.
	curCap := k.capacity
	selectedItems := make([]item, 0)
	for curItem := len(items) - 1; curItem > 0; curItem-- {
		if values[curItem-1][curCap] != values[curItem][curCap] {
			selectedItems = append(selectedItems, items[curItem])

			curCap -= items[curItem].weight
		}
	}
	if values[0][curCap] > 0 {
		selectedItems = append(selectedItems, items[0])
	}

	return values[len(items)-1][k.capacity], selectedItems
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var capacity, numItems int
	fmt.Sscanf(line, "%d %d", &capacity, &numItems)

	knapsack := knapsack{capacity}
	items := make([]item, numItems)
	for i := 0; i < numItems; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		var val, weight int
		fmt.Sscanf(line, "%d %d", &val, &weight)
		items[i] = item{weight, val}
	}

	fmt.Printf("Solving cap. %d knapsack, %d items\n", capacity, numItems)
	fmt.Println(pack(knapsack, items))
}
