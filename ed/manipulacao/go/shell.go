package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	vetAux := make([]int, 0, len(vet))
	for _, val := range vet {
		if val > 0 {
			vetAux = append(vetAux, val)
		} 
	}
	return vetAux
}

func getCalmWomen(vet []int) []int {
	vetAux := make([]int, 0, len(vet))
	for _, val := range vet {
		if val < 0 && val > -10 {
			vetAux = append(vetAux, val)
		} 
	}
	return vetAux
}

func sortVet(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet); j++ {
			if vet[i] < vet[j] {
				vet[i], vet[j] = vet[j], vet[i]
			}
		}
	}
	return vet
}

func abs (num int) int {
	if num < 0 {
		return -num
	}
	 return num
}

func sortStress(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet); j++ {
			if abs(vet[i]) < abs(vet[j]) {
				vet[i], vet[j] = vet[j], vet[i]
			}
		}
	}
	return vet
}

func reverse(vet []int) []int{
	aux := make([]int, len(vet))

	for i := 0; i < len(vet); i++{
		aux[i] = vet[len(vet)-1-i]
		//aux = append(aux, vet[i])
	}
	//copy(vet, aux)
	return aux
}

func unique(vet []int) []int {
	unicos := make(map[int]bool)
	var aux []int

	for _, valor := range vet {

		_, existe := unicos[valor]

		if !existe {
			unicos[valor] = true
			aux = append(aux, valor)
		}

		//num[valor] = true		
	}
	return aux
	// _ = vet
	// return nil
}

func repeated(vet []int) []int {
	unicos := make(map[int]bool)
	var repitidos []int

	for _, valor := range vet {

		_, existe := unicos[valor]

		if existe {
			repitidos = append(repitidos, valor)
		} else {
			unicos[valor] = true

		}

		//num[valor] = true		
	}
	return repitidos
	// _ = vet
	// return nil
	// _ = vet
	// return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

