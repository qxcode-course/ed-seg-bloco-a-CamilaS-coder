package main

import "fmt"

func matar(elem []int, pos int) ([]int, int) {
    morto := (pos + 1) % len(elem)
    elem = append(elem[:morto], elem[morto+1:]...)
    pos = morto % len(elem)
    //pos = (pos + 1) % len(elem)
    return elem, pos
}

func main() {
    var N, E int
    fmt.Scan(&N, &E)

    vetor := make([]int, N)
    for i := 0; i < N; i++ {
        vetor[i] = i + 1
    }

    pos := E - 1

    for len(vetor) > 1 {
        fmt.Print("[ ")
        for i := 0; i < len(vetor); i++ {
            if i == pos {
                fmt.Print(vetor[i], ">")
            } else {
                fmt.Print(vetor[i])
            }
            if i < len(vetor)-1 {
                fmt.Print(" ")
            }
        }
        fmt.Println(" ]")

        vetor, pos = matar(vetor, pos)
    }

    fmt.Print("[ ")
    fmt.Print(vetor[0], ">")
    fmt.Println(" ]")
}
