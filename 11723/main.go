// https://www.acmicpc.net/problem/11723
// silver 5
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data map[int]struct{}
}

func (s *Set) Add(v int) {
	s.data[v] = struct{}{}
}

func (s *Set) Remove(v int) {
	delete(s.data, v)
}

func (s *Set) Check(v int) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set) Toggle(v int) {
	if s.Check(v) {
		s.Remove(v)
	} else {
		s.Add(v)
	}
}

func (s *Set) All() {
	for i := 1; i <= 20; i++ {
		s.Add(i)
	}
}

func (s *Set) Empty() {
	s.data = make(map[int]struct{})
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := getLine(reader)
	set := Set{data: make(map[int]struct{})}
	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		token := strings.Split(strings.TrimSpace(input), " ")
		switch token[0] {
		case "add":
			num, _ := strconv.Atoi(token[1])
			set.Add(num)
		case "remove":
			num, _ := strconv.Atoi(token[1])
			set.Remove(num)
		case "check":
			num, _ := strconv.Atoi(token[1])
			if set.Check(num) {
				fmt.Fprint(writer, "1\n")
			} else {
				fmt.Fprint(writer, "0\n")
			}
		case "toggle":
			num, _ := strconv.Atoi(token[1])
			set.Toggle(num)
		case "all":
			set.All()
		case "empty":
			set.Empty()
		}
	}
}

func getLine(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(input))
	return num
}
