// https://www.acmicpc.net/problem/1260
// silver 2
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, _ := reader.ReadString('\n')
	parts := strings.Fields(input)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	v, _ := strconv.Atoi(parts[2])

	graph := NewGraph(n)
	for i := 0; i < m; i++ {
		input, _ = reader.ReadString('\n')
		parts = strings.Fields(input)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		graph.addEdge(u, v)
	}
	graph.dfs(v)
	graph.bfs(v)
	fmt.Fprintln(writer, graph.DfsVisited)
	fmt.Fprintln(writer, graph.BfsVisited)
}

type Queue struct {
	v list.List
}

func NewQueue() *Queue {
	v := list.List{}
	return &Queue{v}
}

func (q *Queue) Push(v int) {
	q.v.PushBack(v)
}

func (q *Queue) Pop() int {
	if q.v.Len() == 0 {
		return 0
	}
	e := q.v.Front()
	q.v.Remove(e)
	return e.Value.(int)
}

type Graph struct {
	adj           [][]int
	dfsVisitedMap map[int]bool
	bfsVisitedMap map[int]bool

	DfsVisited string
	BfsVisited string
}

func NewGraph(n int) *Graph {
	return &Graph{
		make([][]int, n+1),
		make(map[int]bool),
		make(map[int]bool),
		"",
		"",
	}
}

func (g *Graph) addEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
	sort.Sort(sort.IntSlice(g.adj[u]))
	sort.Sort(sort.IntSlice(g.adj[v]))
}

func (g *Graph) dfs(v int) {
	g.dfsVisitedMap[v] = true
	g.DfsVisited += strconv.Itoa(v) + " "
	for _, next := range g.adj[v] {
		if g.dfsVisitedMap[next] {
			continue
		}
		g.dfs(next)
	}
}

func (g *Graph) bfs(v int) {
	q := NewQueue()
	q.Push(v)
	g.bfsVisitedMap[v] = true
	for q.v.Len() > 0 {
		current := q.Pop()
		g.BfsVisited += strconv.Itoa(current) + " "
		for _, next := range g.adj[current] {
			if g.bfsVisitedMap[next] {
				continue
			}
			q.Push(next)
			g.bfsVisitedMap[next] = true
		}
	}

}
