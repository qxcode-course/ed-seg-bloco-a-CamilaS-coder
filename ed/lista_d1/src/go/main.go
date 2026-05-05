package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	dado int
	next *Node
	prev *Node
}

type List struct {
	root *Node
}

func NewLList() *List {

	list := &List{}
	list.root = &Node{}
	list.root.next = list.root
	list.root.prev = list.root

	return list
}

func inserir (node *Node, value int) {
	novo := &Node {
		dado: value,
		next: node,
		prev: node.prev,
	}

	// novo.next = node
	// novo.prev = node.prev
	node.prev.next = novo
	node.prev = novo

}
func remover(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (n *Node) String() string {
	return fmt.Sprint(n.dado)
} 

func (list *List) String() string {
	sai := ""
	sai += "["

	for x := list.root.next; x != list.root; x = x.next {
		if x.next != list.root{
			sai += fmt.Sprintf("%s, ", x)
		} else {
			sai += fmt.Sprintf("%s", x)
		}

		//sai += fmt.Sprintf("%s", x)
	}

	sai += ("]")
	return sai
}

func (list *List) Size() int {
	var contador = 0
	for x := list.root.next; x != list.root; x = x.next {
		contador++
	}
	return  contador
}

func (list *List) Clear(){
	list.root.next = list.root
	list.root.prev = list.root
}

func (l *List) PushFront(value int) {
	inserir(l.root.next, value)
}

func (l *List) PushBack(value int) {
	inserir(l.root, value)
}

func (l *List) PopFront() {
	remover(l.root.next)
}

func (l *List) PopBack() {
	remover(l.root.prev)
}

// func (l *List) String() string {
// 	var result strings.Builder
// 	result.WriteString("[")e
// 	for x := l.root.next; x != l.root; x = x.next {
// 		result.WriteString(fmt.Sprintf("%s ", x))
// 	}
// 	result.WriteString("]")
// 	return result.String()
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			 fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			 ll.PopFront()
		case "clear":
			 ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
