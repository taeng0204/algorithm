// https://www.acmicpc.net/problem/11724
// silver 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

type Graph struct {
	adj [][]int
}

func NewGraph(n int) *Graph {
	return &Graph{make([][]int, n+1)}
}

func (g *Graph) addEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) getCC() int {
	var stack Stack
	var visited map[int]bool = make(map[int]bool)
	count := 0

	for i := 1; i < len(g.adj); i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		stack.Push(i)
		for len(stack) > 0 {
			current := stack.Pop()

			for _, next := range g.adj[current] {
				if !visited[next] {
					stack.Push(next)
					visited[next] = true
				}
			}
		}
		count++
	}
	return count
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

	cc := graph.getCC()

	fmt.Fprintln(writer, cc)
}
