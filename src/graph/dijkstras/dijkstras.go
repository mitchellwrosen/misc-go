package dijkstras

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Graph struct {
	Nodes []Node
}

type Node struct {
	Id    int
	Edges []Edge
}

type Edge struct {
	To     int
	Weight int
}

func NewGraph(r io.Reader) (*Graph, error) {
	graph := new(Graph)

	reader := bufio.NewReader(r)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		nums := strings.Split(string(line[:len(line)-1]), " ") // Strip newline.

		// Each line is: node [edgeTo weight]+
		if len(nums)%2 == 0 {
			return nil, fmt.Errorf("Invalid line %s", line)
		}

		nodeId, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}

		node := Node{Id: nodeId}
		for i := 1; i < len(nums); i += 2 {
			edgeTo, err := strconv.Atoi(nums[i])
			if err != nil {
				return nil, err
			}

			edgeWeight, err := strconv.Atoi(nums[i+1])
			if err != nil {
				return nil, err
			}

			node.Edges = append(node.Edges, Edge{edgeTo, edgeWeight})
		}

		graph.Nodes = append(graph.Nodes, node)
	}

	return graph, nil
}

func (g *Graph) ShortestPath(from, to int) int {
	// TODO
	return 0
}

//func main() {
//if len(os.Args) != 2 {
//fmt.Println("Usage: %s filename\n", os.Args[0])
//os.Exit(1)
//}

//file, err := os.Open(os.Args[1])
//if err != nil {
//panic(err)
//}

//graph, err := NewGraph(file)
//}
