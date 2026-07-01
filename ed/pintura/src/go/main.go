package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	corOriginal := image[sr][sc]

	if corOriginal != color {
		dfs(image, sr, sc, corOriginal, color)
	}

	// _ := image
	return image
}

func dfs(image [][]int, l int, c int, corOriginal int, novaCor int) {
	if l < 0 || l >= len(image) || c < 0 || c >= len(image[0]) || image[l][c] != corOriginal {
		return
	}

	image[l][c] = novaCor

	dfs(image, l + 1, c, corOriginal, novaCor)
	dfs(image, l - 1, c, corOriginal, novaCor)
	dfs(image, l, c + 1, corOriginal, novaCor)
	dfs(image, l, c - 1, corOriginal, novaCor)
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
