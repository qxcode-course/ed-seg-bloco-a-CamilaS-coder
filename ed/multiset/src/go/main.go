package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	//"errors"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(nova_capacity int) *MultiSet {
	return &MultiSet{
		data:  make([]int, 0, nova_capacity),
		size: 0,
		capacity: nova_capacity,
	}

	//return novo_slice
}

func (m *MultiSet) Size() int {
	return m.size
}

func (m *MultiSet) String() string {
	sai := ""
	// if m.Size() == 0 {
	// 	return "[]"
	// }
	sai += "["
	for i := 0; i < m.Size(); i++ {
		if i != m.Size() - 1 {
			sai += fmt.Sprintf("%d, ", m.data[i])
		} else {
			sai += fmt.Sprintf("%d", m.data[i])
		}
	}
	sai += "]"
	return sai
}

func (m *MultiSet) expand() {
		nova_capacity := m.capacity * 2

		if m.Size() == 0 {
			nova_capacity = 1
		}

		novo_slice := make([]int, m.Size(), nova_capacity)

		for i := 0; i < m.Size(); i++ {
			novo_slice[i] = m.data[i]
		}

		m.data = novo_slice
		m.capacity = nova_capacity
}

func (m *MultiSet) insert(value int, index int) error {
	if m.Size() == m.capacity {
		m.expand()
	}

	m.data = append(m.data, 0)

	for i := m.Size(); i > index; i-- {
		m.data[i] = m.data[i - 1]
	}

	m.data[index] = value
	m.size++
	return nil
}

func (m *MultiSet) Search(value int, ) (bool, int) {
	inicio := 0
	fim := m.Size()

	for inicio != fim {
		meio := (inicio + fim) / 2

		if m.data[meio] == value {
			inicio = meio + 1
		} else if value < m.data[meio] {
			fim = meio
		} else if value > m.data[meio] {
			inicio = meio + 1
		}

	}

	if  inicio > 0 && m.data[inicio - 1] == value {
		return true, inicio - 1
	}

	return false, inicio


}

func (m *MultiSet) Insert(value int) {
	_, index := m.Search(value)

	m.insert(value, index)
}

func (m *MultiSet) Contains(value int) bool {
	for i := 0; i < m.Size(); i++ {
		if m.data[i] == value {
			fmt.Println("true")
			return true
		}
	}
	fmt.Println("false")
	return false
}

func (m *MultiSet) erase(index int) error {
	for i := index; i < m.Size() - 1; i++ {
		m.data[i] = m.data[i + 1]
	}
	m.size--
	m.data = m.data[:m.size]
	return nil
}

func (m *MultiSet) Erase(value int) error {
		// if !m.Contains(value) {
		// 	return errors.New("value not found")
		// }

		booleano, index := m.Search(value)

		if booleano == false {
			fmt.Println("value not found")
		}

		m.erase(index)
		return nil
}


func (m *MultiSet) Clear() {
	m.size = 0
}

func (m *MultiSet) Count(value int) int {
	var contador int
	for i := 0; i < m.Size(); i++ {
		if m.data[i] == value {
			contador++
		}
	}
	fmt.Println(contador)
	return contador
}

func (m *MultiSet) Unique() int {

	if m.Size() == 0 {
		fmt.Println(0)
		return 0
	}

	var diferente = m.data[0]
	var contador = 1
	for i := 1; i < m.Size(); i++ {
		if diferente != m.data[i] {
			contador++
		}

		diferente = m.data[i]
		//continue
	}
	fmt.Println(contador)
	return contador
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	 m := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			m = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				m.Insert(value)
			}
		case "show":
			fmt.Println(m)
		case "erase":
			 value, _ := strconv.Atoi(args[1])
			 m.Erase(value)
		case "contains":
			 value, _ := strconv.Atoi(args[1])
			 m.Contains(value)
		case "count":
			 value, _ := strconv.Atoi(args[1])
			 m.Count(value)
		case "unique":
			m.Unique()
		case "clear":
			m.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
