package main
import "fmt"

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


func aux_primo(n int, atual int, contador int) int {
    if contador == n {
        return atual - 1
    }

    if eh_primo(atual,  2) {
        return aux_primo(n, atual + 1, contador + 1)
    }
    
    return aux_primo(n, atual + 1, contador)
}

func enesimo_primo(n int) int {
return aux_primo(n, 2, 0)
}

func main() {

    var n int
    fmt.Scan(&n)

    fmt.Println(enesimo_primo(n))

   // fmt.Println("Hello, World!")
}