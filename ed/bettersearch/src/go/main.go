package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {

	var low = 0;
	var high = len(slice)
	//var meio = 0

	for low != high {
		meio := (low + high) / 2

		if slice[meio] == value {
			return true, meio
		}

		if value < slice[meio] {
			high = meio
		}

		if  value > slice[meio] {
			low = meio + 1
		}

	}


	return false, low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
