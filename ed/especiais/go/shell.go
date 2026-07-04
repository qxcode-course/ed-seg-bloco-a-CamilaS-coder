package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func occurr(vet []int) []Pair {
	ocorre := make(map[int]int)
	var pares []Pair


	for _, num := range vet {
		if num < 0 {
			num = abs(num)
			
		}
		ocorre[num]++
	}

	for num, qtd := range ocorre {
		pair := Pair {One : num, Two : qtd}
		pares = append(pares, pair)

	}

	sort.Slice(pares, func(i int, j int) bool {
		return pares[i].One < pares[j].One
	})	
	
	return pares
	
}
func teams(vet []int) []Pair {
	// 1. Caso base: se o vetor estiver vazio, retorna logo um slice vazio
	if len(vet) == 0 {
		return []Pair{}
	}

	var pares []Pair
	
	// 2. Começamos com o primeiro elemento como referência
	stressAtual := vet[0]
	contador := 0

	// 3. Percorremos o vetor uma única vez (Complexidade O(n))
	for _, num := range vet {
		// Se o stress atual for igual ao do "líder" do time, o time cresce
		if num == stressAtual {
			contador++
		} else {
			// Se o stress mudou, o time anterior "fechou"
			// Guardamos o resultado do time que acabou de passar
			pares = append(pares, Pair{One: stressAtual, Two: contador})
			
			// Resetamos o líder para o novo stress e o contador volta a 1
			stressAtual = num
			contador = 1
		}
	}

	// 4. O "FECHAMENTO": Como o loop acaba, o último time processado 
	// ainda não foi adicionado. Precisamos desse append final:
	pares = append(pares, Pair{One: stressAtual, Two: contador})

	return pares
}

func mnext(vet []int) []int {
	res := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			hasWomanNei := false
			if i > 0 && vet[i - 1] < 0 {
				hasWomanNei = true
			}

			if i < len(vet)- 1 && vet[i + 1] < 0 {
				hasWomanNei = true
			}

			if hasWomanNei {
				res[i] = 1
			}
		}
	}
	
	return res
}

func alone(vet []int) []int {
	res := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			hasWomanNei := false
			if i > 0 && vet[i-1] < 0 {
				hasWomanNei = true
			}

			if i <len(vet) - 1 && vet[i+1] < 0 {
				hasWomanNei = true
			}

			if !hasWomanNei {
				res[i] = 1
			}
		}
	}
	
	return res
}

func couple(vet []int) int {
	man := make(map[int]int)
	woman := make(map[int]int)

	for _, x := range vet {
		if x > 0 {
			man[x]++
		}else if x < 0 {
			woman[-x]++
		}
	}


	totalcasais := 0

	for stress, qtdHomens := range man {
		qtdMulheres := woman[stress]

		if qtdHomens < qtdMulheres {
			totalcasais += qtdHomens
		} else {
			totalcasais += qtdMulheres
		}
	}
	
	return totalcasais
}

func hasSubseq(vet []int, seq []int, pos int) bool {

	if pos + len(seq) > len(vet) {
		return false

	}

	for i := 0; i < len(seq); i++ {
		if vet[pos + i] != seq[i] {
			return false
		}

	}

	return true
}

func subseq(vet []int, seq []int) int {

	for i := 0; i <= len(vet) - len(seq); i++ {
		if hasSubseq(vet, seq, i ) {
			return i
		}
	}
	
	return -1
}

func erase(vet []int, posList []int) []int {

	remover := make(map[int]bool)

	for _, pos := range posList {
		remover[pos] = true
	}

	var res []int

	for i, val := range vet {
		if !remover[i] {
			res = append(res, val)
		}
	}

	return res
}

func clear(vet []int, value int) []int {

	var res []int

	for _, val := range vet {
		if val != value {
			res = append(res, val)
		}
	}
	
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
