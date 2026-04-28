package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	if x == 1 {
		return false
	}

	if x % div == 0 && x != div {
		return false
	}

	if x == div {
		return true
	}

	//var num = x % div

	// if num != 0 {
	// 	div++
	// }

	return eh_primo(x, div + 1)
	// _, _ = x, div
	// return false;
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
