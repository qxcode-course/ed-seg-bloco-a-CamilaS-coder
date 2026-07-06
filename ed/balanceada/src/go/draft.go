package main
import (

    "fmt"
    "bufio"
    "os"
)

type Stack struct {
    data []rune
}

func (s *Stack) Push(val rune) {
    s.data = append(s.data, val)
}

func (s *Stack) Pop() rune {
    if len(s.data) == 0 {
        return 0
    }

    val := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data) - 1]
    return val
}

func (s *Stack) IsEmpty() bool {
    return len(s.data) == 0
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }

    text := scanner.Text()

    pilha := &Stack{}
    balanceado := true

    for _, char := range text {
        if char == '(' || char == '[' {
            pilha.Push(char)
        } else if char == ')' || char == ']' {
            if pilha.IsEmpty() {
                balanceado = false
                break
            }

            topo := pilha.Pop()

            if (char == ')' && topo != '(') || (char == ']' && topo != '[') {
                balanceado = false
                break
            }
        }
    }

    if !pilha.IsEmpty() {
        balanceado = false
    }

    if balanceado {
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }
}