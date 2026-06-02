package main

import "fmt"


func eh_primo(n int, div int) bool {
	if n == 1 {
		return false
	}

	if n%div == 0 && n != div {
		return false
	}

	if n == div {
		return true
	}

	return eh_primo(n, div+1)
}


func aux_primo(n int, atual int, primos []int) []int {
	
	if len(primos) == n {
		return primos
	}

	
	if eh_primo(atual, 2) {
		return aux_primo(n, atual+1, append(primos, atual))
	}

	
	return aux_primo(n, atual+1, primos)
}


func priminhos(n int) []int {
	return aux_primo(n, 2, []int{})
}

func main() {
	var n int
	fmt.Scan(&n)
	
	resultado := priminhos(n)
    fmt.Print("[")
	for i, v := range resultado {
		fmt.Print(v)
		if i < len(resultado)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
	
	
}