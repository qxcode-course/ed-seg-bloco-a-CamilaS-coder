package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ehValido(pos int, d rune, str []rune, L int) bool {

	for i := 1; i <= L; i++ {
		vizinhoEsquerda := pos - i
		if vizinhoEsquerda >= 0 {
			if str[vizinhoEsquerda] == d {
				return false
			}
		} else {
			break
		}
	}


	for i := 1; i <= L; i++ {
		vizinhoDireita := pos + i
		if vizinhoDireita < len(str) {
			if str[vizinhoDireita] == d {
				return false
			}
		} else {
			break
		}
	}

	return true
}

func resolver(pos int, str []rune, L int) bool {

	if pos == len(str) {
		return true
	}

	
	if str[pos] != '.' {
		return resolver(pos+1, str, L)
	}


	for d := 0; d <= L; d++ {
		charDigito := rune('0' + d)

		if ehValido(pos, charDigito, str, L) {
			str[pos] = charDigito 

			if resolver(pos+1, str, L) {
				return true 
			}

			str[pos] = '.' 
		}
	}

	return false 
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)


	if !scanner.Scan() {
		return
	}
	sequencia := scanner.Text()


	if !scanner.Scan() {
		return
	}
	L, _ := strconv.Atoi(scanner.Text())


	strRunes := []rune(sequencia)

	
	if resolver(0, strRunes, L) {
		fmt.Println(string(strRunes))
	}
}