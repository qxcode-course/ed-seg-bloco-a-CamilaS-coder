package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var elementos []string

	for n := l.Front(); n != l.End(); n = n.next {
		if n == sword {
			elementos = append(elementos, fmt.Sprint(n.Value) + ">")
		} else {
			elementos = append(elementos, fmt.Sprint(n.Value))
		}
	}
	return "[ " + strings.Join(elementos, " ") + " ]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it == nil {
		return nil
	}

	proximo := it.next

	if proximo == l.End() {
		proximo = proximo.next
	}
	return proximo
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
