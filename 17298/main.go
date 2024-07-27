// https://www.acmicpc.net/problem/17298
// gold 4
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	n := getLine(reader)
	numbers := getNumbers(reader, n)

	fmt.Fprintln(writer, RightBigNumber(n, numbers))
}

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s Stack) Top() int {
	if len(s) == 0 {
		return -1
	}
	return s[len(s)-1]
}

func RightBigNumber(n int, inputs []int) string {
	result := make([]string, n)
	stack := make(Stack, 1, n)

	for i := 0; i < n-1; i++ {
		for stack.Top() != -1 && inputs[stack.Top()] < inputs[i+1] {
			result[stack.Pop()] = strconv.Itoa(inputs[i+1])
		}
		stack.Push(i + 1)
	}

	for stack.Top() != -1 {
		result[stack.Pop()] = "-1"
	}
	return strings.Join(result, " ")
}

func getLine(reader *bufio.Reader) int {
	n, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(n))
	return num
}

func getNumbers(reader *bufio.Reader, line int) []int {
	n, _ := reader.ReadString('\n')
	nums := strings.Split(strings.TrimSpace(n), " ")
	numbers := make([]int, line, line)
	wg := sync.WaitGroup{}
	wg.Add(line)

	for i, num := range nums {
		go func() {
			number, _ := strconv.Atoi(num)
			numbers[i] = number
		}()
	}
	wg.Wait()
	return numbers
}
