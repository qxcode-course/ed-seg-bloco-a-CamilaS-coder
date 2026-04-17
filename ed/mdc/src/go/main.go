package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	var res, quo int
	quo = a / b
	res = a % b
	a = b * quo + res
	return mdc(b, res)
	
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
