 package main

import (
	"bufio"
	"fmt"
	"os"
)
func dfs(grid [][]byte, i int, j int) {
	if (i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0') {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i + 1, j)
	dfs(grid, i - 1, j)
	dfs(grid, i, j + 1)
	dfs(grid, i, j - 1)
}
  
// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	
	var contador int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if (grid[i][j] == '1') {
				dfs(grid, i, j)
				contador++
			}
		}
	}

	return contador



	// _ := grid
	// return 0
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
