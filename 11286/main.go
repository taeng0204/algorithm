// https://www.acmicpc.net/problem/11286
// silver 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	data []int
}

func (h *Heap) Init() {
	h.data = make([]int, 1)
}

func (h *Heap) Push(v int) {
	h.data = append(h.data, v)
	h.up()
}

func (h *Heap) Pop() int {
	if len(h.data) == 1 {
		return 0
	}
	result := h.data[1]
	h.data[1] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down()
	return result
}

func (h *Heap) up() {
	for i := len(h.data) - 1; i > 1; i = i / 2 {
		if !isLowAbs(h.data[i], h.data[i/2]) {
			break
		}
		h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
	}
}

func (h *Heap) down() {
	for i := 1; i*2 < len(h.data); {
		child := i*2 + 1
		if i*2+1 >= len(h.data) {
			child = i * 2
		} else if isLowAbs(h.data[i*2], h.data[i*2+1]) {
			child = i * 2
		}

		if !isLowAbs(h.data[child], h.data[i]) {
			break
		}
		h.data[i], h.data[child] = h.data[child], h.data[i]
		i = child
	}
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := getLine(reader)
	heap := Heap{}
	heap.Init()

	for i := 0; i < n; i++ {
		num := getLine(reader)
		if num == 0 {
			fmt.Fprintln(writer, heap.Pop())
		} else {
			heap.Push(num)
		}
	}
}

func getLine(reader *bufio.Reader) int {
	n, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(n))
	return num
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isLowAbs(a, b int) bool {
	if abs(a) < abs(b) {
		return true
	}
	if abs(a) == abs(b) {
		if a < b {
			return true
		}
	}
	return false
}
