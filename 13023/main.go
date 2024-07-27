// https://www.acmicpc.net/problem/13023
// gold 5
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adj [][]int
}

func NewGraph(n int) *Graph {
	return &Graph{make([][]int, n)}
}

func (g *Graph) addEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) isABCDE() bool {
	for i := 0; i < len(g.adj); i++ {
		var visited map[int]bool = make(map[int]bool)
		if g.dfs(visited, i, 0) {
			return true
		}
	}
	return false
}

func (g *Graph) dfs(visited map[int]bool, current int, depth int) bool {
	if depth == 4 {
		return true
	}

	newVisited := copyMap(visited)
	newVisited[current] = true
	for _, next := range g.adj[current] {
		if newVisited[next] {
			continue
		}
		if g.dfs(newVisited, next, depth+1) {
			return true
		}
	}
	return false
}

func copyMap(original map[int]bool) map[int]bool {
	newMap := make(map[int]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, _ := reader.ReadString('\n')
	parts := strings.Fields(input)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	graph := NewGraph(n)
	for i := 0; i < m; i++ {
		input, _ = reader.ReadString('\n')
		parts = strings.Fields(input)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		graph.addEdge(u, v)
	}

	isABCDE := graph.isABCDE()
	if isABCDE {
		fmt.Fprintln(writer, 1)
	} else {
		fmt.Fprintln(writer, 0)
	}
}
