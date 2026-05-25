package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	//"errors"
)

type Set struct {
	data []int
	size int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data: make([]int, 0, capacity),
		size: 0,
		capacity: capacity,
	}
}
func (s *Set) Size() int {
	return s.size
}

func (s *Set) String() string {
	sai := ""
	if s.size == 0 {
		return"[]"
	}
	sai += "["
	for i := 0; i < s.Size(); i++ {
		if i != s.Size() - 1 {
			sai += fmt.Sprintf("%d, ", s.data[i])
		} else {
			sai += fmt.Sprintf("%d", s.data[i])
		}
	}
	sai += "]"
	return sai
}

func (s *Set) binarysearch(value int) int {
	//var meio int
	var inicio = 0
	var fim = s.Size() - 1
	
	for inicio <= fim{

		meio := (inicio + fim) / 2

		if s.data[meio] == value {
			//s.data[meio] = value
			return meio
		}
		if value < s.data[meio] {
			fim = meio - 1
		}
		if value > s.data[meio] {
			inicio = meio + 1
		}
	}
	return -1
}

func (s  *Set) reserve(nova_capacity int) {
	var novo_slice = make([]int, nova_capacity)

	for i := 0; i < s.Size(); i++ {
		novo_slice[i] = s.data[i]
	}

	s.data = novo_slice
	s.capacity = nova_capacity
}

func (s *Set) insert(value int, index int) {

	if s.Size() == s.capacity {
		nova_capacity := s.capacity * 2

		if s.capacity == 0 {
			nova_capacity = 1
		}

		s.reserve(nova_capacity)
	}

	s.data = append(s.data, 0)

	for i := s.Size(); i > index; i-- {
		s.data[i] = s.data[i - 1] 
	}
	
	s.data[index] = value
	s.size++
	//return nil
}

func (s *Set) Insert(value int) {
	
	if s.binarysearch(value) != -1{
		return
	}

	// if s.Size() == 0 {
	// 	s.insert(value, 0)
	// }

	for i := 0; i < s.Size(); i++ {
		if s.data[i] < value {
			continue
			//return
		}
		
		s.insert(value, i)
		return
	}

	s.insert(value, s.Size())

}

func (s *Set) Contains(value int) bool {
	for i := 0; i < s.Size(); i++ {
		if s.data[i] == value {
			fmt.Println("true")
			return true
		}
	}
	fmt.Println("false")
	return false
}

func (s *Set) erase(index int) error {
	// if !s.Contains(s.data[index]) {
	// 	return errors.New("value not found")
	// }

	for i := index; i < s.Size() - 1; i++ {
		s.data[i] = s.data[i + 1]
	}
	s.size--
	return nil

}


func (s *Set) Erase(value int) bool {
	index := s.binarysearch(value)
	
	if index == -1 {
        fmt.Println("value not found") 
        return false
    }

	// if index == -1 {
	// 	//fmt.Println("value not found")
	// 	return false
	// }
	
	s.erase(index)

	return true
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	 s := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			s = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				s.Insert(value)
			}
		case "show":
			fmt.Println(s)
		case "erase":
			 value, _ := strconv.Atoi(parts[1])
			 s.Erase(value)

		case "contains":
			 value, _ := strconv.Atoi(parts[1])
			 s.Contains(value)
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
