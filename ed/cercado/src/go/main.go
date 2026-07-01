package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func dfs(board [][]byte, l, c int) {

	if(l < 0 || l >= len(board) || c < 0 || c >= len(board[0]) || board[l][c] != 'O') {
		return
	}

	board[l][c] = 'T'

	dfs(board, l + 1, c)
	dfs(board, l - 1, c)
	dfs(board, l, c + 1)
	dfs(board, l, c - 1)

	// _ = board
}

func solve(board [][]byte){

	nlinhas, ncolunas := len(board), len(board[0])
	for i := 0; i < nlinhas; i++ {
		for j := 0; j < ncolunas; j++ {
			if(i == 0 || i == nlinhas - 1 || j == 0 || j == ncolunas - 1 && board[i][j] == 'O') {
				dfs(board, i, j)
			}
		}
	}

	for i := 0; i < nlinhas; i++ {
		for j := 0; j < ncolunas; j++ {
			if(board[i][j] == 'O') {
				board[i][j] = 'X'
			} else if board[i][j] == 'T'{
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
