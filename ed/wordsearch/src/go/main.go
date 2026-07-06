package main

import (
	"bufio"
	"fmt"
	"os"
)

var dr = []int{-1, 1, 0, 0}
var dc = []int{0, 0, -1, 1}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func dfs(r, c, index int, grid [][]byte, word string) bool {

	if index == len(word) {
		return true
	}

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != word[index] {
		return false
	}
	temp := grid[r][c]
	grid[r][c] = byte('#')

	for i := 0; i < 4; i++ {
		if dfs(r + dr[i], c + dc[i], index + 1, grid, word) {
			return true
		}
	}

	grid[r][c] = temp

	return false

}

func exist(grid [][]byte, word string) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}

	linhas := len(grid)
	colunas := len(grid[0])

	for r := 0; r < linhas; r++ {
		for c := 0; c < colunas; c++ {
			if grid[r][c] == word[0] {
				if dfs(r, c, 0, grid, word) {
					return true
				}
			}
		}
	}

	return false
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
