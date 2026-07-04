package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	fila := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", 
		"I", "J", "K", "L", "M", "N", "O", "P",
	}

	scanner := bufio.NewScanner(os.Stdin)

	
	for i := 0; i < 15; i++ {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		parts := strings.Fields(line)
		
		
		if len(parts) < 2 {
			i--
			continue
		}

		
		golEsquerda, _ := strconv.Atoi(parts[0])
		golsDireita, _ := strconv.Atoi(parts[1])

		
		teamEsquerda := fila[0]
		teamDireita := fila[1]

	
		fila = fila[2:]

		
		if golEsquerda > golsDireita {
			fila = append(fila, teamEsquerda)
		} else {
			fila = append(fila, teamDireita)
		}
	}


	fmt.Println(fila[0])
}
