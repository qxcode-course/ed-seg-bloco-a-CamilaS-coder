package main
import "fmt"
func main() {
	var N int
	fmt.Scan(&N)

	fila := make([]int, N)

	for i:= 0; i < N; i++ {
		//var num int
		fmt.Scan(&fila[i])
		//fila[i] =  num
	}

	var M int
	fmt.Scan(&M)

	remove := make(map[int]bool)

	for i := 0; i < M; i++ {
		var num int
		fmt.Scan(&num)
		remove[num] = true
	}

	var aux []int

	for _, num := range fila {
		

		if !remove[num] {
			aux = append(aux, num)
		}
		
	}
	for i := range aux {
		fmt.Print(aux[i], " ")
	}
	fmt.Println()

	// for _, v := range remove {
	// 	str := ""
	// 	_, existe := fila[v]

	// 	if !existe {
	// 		str += fmt.Sprintf("%d ", v)
	// 	}

	// 	fmt.Print(str)
	// }


}
