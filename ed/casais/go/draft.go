package main
import "fmt"
func main() {

    var n int
    var i int
    fmt.Scan(&n)

    animais := make([]int, n)
    var casais []int
   // var vet  []int

    for i = 0;  i < n; i ++ {
        fmt.Scan(&animais[i])
    }

    for i = 0; i < len(animais); i ++ {
        if animais[i] == animais[i] && animais[i] < 0 {

            casais = append(casais, 0)
            
        }

    }
    fmt.Println(len(casais))
    
}
