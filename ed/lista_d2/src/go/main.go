 package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct{
	dado int
	next *Node
	prev *Node
	list *List
}

type List struct {
	root *Node
}

func NewLList() *List {

	list := &List{}
	list.root = &Node{list: list}
	list.root.next =  list.root
	list.root.prev = list.root

	return list
}

func (l *List) Insert (node *Node, value int) {
	novo := &Node {
		dado: value,
		next : node,
		prev: node.prev,
		list : l,
	}
	node.prev.next = novo
	node.prev = novo


}

func (list *List) Remove (node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

//func 

func (list *List) PushBack(value int) {
	list.Insert(list.root, value)
}

func (list *List) PushFront(value int) {
	list.Insert(list.root.next, value)
}

func (node *Node) String() string {
	return fmt.Sprint(node.dado)
}

func (list *List) String() string{

	sai := ""
	sai += "["
	for x := list.root.next ; x != list.root ; x = x.next {

		if x.next != list.root {
			sai += fmt.Sprintf("%s, ", x)
		} else {
			sai += fmt.Sprintf("%s", x)
		}

		//fmt.Print(i)
	}
	
	sai += "]"
	return sai
 }

func (list *List) Size() int {
	var contador = 0

	for x := list.root.next; x != list.root; x =  x.next {

		contador++
		
	}
	return contador
} 

func (list *List) Search(value int) *Node {
	for x := list.root.next; x != list.root; x = x.next {
		if x.dado == value {
			return x
		}
	}

	return nil
}

func (list *List) PopBack() {
	if list.Size() > 0 {
	list.Remove(list.root.prev)
	}
}

func (list *List) PopFront() {
	if list.Size() > 0 {
	list.Remove(list.root.next)
	}
}

func (list *List) Clear() {
	list.root.next = list.root
	list.root.prev = list.root
}

func (list *List) Front() *Node {
	if list.Size() == 0 {
		return nil
	}

	return list.root.Next()
}

func (list *List) Back() *Node {
	if list.Size() == 0 {
		return nil
	}

	return list.root.Prev()
}

func (node *Node) Next() *Node {
	if node.next == node.list.root {
		return nil
	}

	return node.next
}

func (node *Node) Prev() *Node {
	if node.prev == node.list.root {
		return nil
	}

	return node.prev
}


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
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.dado)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.dado)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.dado = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
