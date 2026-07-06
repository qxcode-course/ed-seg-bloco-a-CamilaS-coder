package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func podeSomar (index int, alvo int, elementos []int) bool {
    if alvo == 0 {
        return true
    }

    if alvo < 0 || index >= len(elementos) {
        return false
    }

    if podeSomar(index + 1, alvo - elementos[index], elementos) {
        return true
    }

    return podeSomar(index + 1, alvo, elementos)

}


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if !scanner.Scan() {
        return
    }

    linha1 := scanner.Text()
    partes1 := strings.Fields(linha1)

    if len(partes1) < 2 {
        return
    }

    n, _ := strconv.Atoi(partes1[0])
    k, _ := strconv.Atoi(partes1[1])

    if !scanner.Scan() {
        return
    }

    linha2 := scanner.Text()
    partes2 := strings.Fields(linha2)

    elementos := make([]int, n) 
    for i := 0; i < n && i < len(partes2); i++ {
        elementos[i], _ = strconv.Atoi(partes2[i])
    }

    if podeSomar(0, k, elementos) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
   // fmt.Println("Hello, World!")
}