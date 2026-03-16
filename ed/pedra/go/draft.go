package main
import "fmt"
import "math"
func main() {

    var n int
    var i int
    fmt.Scan(&n)
    var pontuacao int
     var melhor = 101
     indice := -1

    for i = 0; i < n; i++ {
        var a, b int
        fmt.Scan(&a, &b)

        if(a < 10 || b < 10) {
            continue
        }
        
        conta := float64(a) - float64(b)

        pontuacao = int(math.Abs(conta))

        if pontuacao < melhor {
            melhor = pontuacao
            indice = i
            //
        }
        
        //fmt.Println("sem ganhador")
        //melhor = pontuacao
        
    }

    if indice == -1 {
        fmt.Println("sem ganhador")
    } else {

        fmt.Println(indice)

    }
    




}
