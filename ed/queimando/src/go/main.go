package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	L, C int
}

// type Stack[T any] struct {
// 	data []T
// }

// func NewStack[T any]() *Stack[T] {
// 	return &Stack[T] {
// 		data: []T{},
// 	}
// }

// func (s *Stack[T]) Push(v T) {
// 	s.data = append(s.data, v)
// }

// func (s *Stack[T]) Pop() T {
// 	index := len(s.data) - 1
// 	v := s.data[index]
// 	s.data = s.data[:index]
// 	return v
// }

// func (s *Stack[T]) IsEmpty() bool {
// 	return len(s.data) == 0
// }


func burnTrees(grid [][]rune, l, c int) {

	nl := len(grid)
	if nl == 0 {
		return
	}

	nc := len(grid[0])

	stack := NewStack[Pos]()
	stack.Push(Pos{L:l, C: c})

	for !stack.IsEmpty() {
		curr := stack.Pop()

		if curr.L < 0 || curr.L >= nl || curr.C < 0 || curr.C >= nc {
			continue
		}

		if grid[curr.L][curr.C] == '#' {
			grid[curr.L][curr.C] = 'o'

			stack.Push(Pos{L: curr.L - 1, C: curr.C})
			stack.Push(Pos{L: curr.L + 1, C: curr.C})
			stack.Push(Pos{L: curr.L, C: curr.C - 1})
			stack.Push(Pos{L: curr.L, C: curr.C + 1})
		}
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
