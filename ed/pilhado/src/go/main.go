package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ponto struct {
	r, c int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	linhaDimensoes := scanner.Text()
	partes := strings.Fields(linhaDimensoes)
	if len(partes) < 2 {
		return
	}

	linhas, _ := strconv.Atoi(partes[0])
	colunas, _ := strconv.Atoi(partes[1])

	labirinto := make([][]rune, linhas)
	var inicio, fim Ponto

	// Leitura da matriz e busca por 'I' e 'F'
	for i := 0; i < linhas; i++ {
		if scanner.Scan() {
			linhaText := scanner.Text()

			// Garante que linhas menores por falta de caracteres no fim sejam preenchidas com espaços
			for len(linhaText) < colunas {
				linhaText += " "
			}
			labirinto[i] = []rune(linhaText)

			for j := 0; j < colunas; j++ {
				if labirinto[i][j] == 'I' {
					inicio = Ponto{r: i, c: j}
				} else if labirinto[i][j] == 'F' {
					fim = Ponto{r: i, c: j}
				}
			}
		}
	}

	caminho := NewStack[Ponto]()
	becos := NewStack[Ponto]()

	caminho.Push(inicio)

	// Direções: Cima, Baixo, Esquerda, Direita
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	// DFS Iterativa
	for !caminho.IsEmpty() {
		atual := caminho.Top()

		labirinto[atual.r][atual.c] = '.'

		if atual == fim {
			break
		}

		var validos []Ponto
		for i := 0; i < 4; i++ {
			nr, nc := atual.r+dr[i], atual.c+dc[i]

			if nr >= 0 && nr < linhas && nc >= 0 && nc < colunas {
				vizinhoChar := labirinto[nr][nc]

				// Válido se não for parede e nem tiver sido visitado ainda
				if vizinhoChar != '#' && vizinhoChar != '.' {
					validos = append(validos, Ponto{r: nr, c: nc})
				}
			}
		}

		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else {
			becos.Push(atual)
			caminho.Pop()
		}
	}

	// Limpando os becos sem saída (voltando a ser espaço vazio)
	for !becos.IsEmpty() {
		beco := becos.Pop()
		labirinto[beco.r][beco.c] = ' '
	}

	// Impressão do resultado final
	for i := 0; i < linhas; i++ {
		fmt.Println(string(labirinto[i]))
	}
}

