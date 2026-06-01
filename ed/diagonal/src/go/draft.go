package main
import "fmt"

func impEsp(k int) {
    if k == 0 {
        return
    }

    fmt.Print(" ")
    impEsp(k - 1)
    //println()
}
func diagonal(s string, k int) {
    if len(s) == 0 {
        return
    }

    impEsp(k)
    fmt.Print(string(s[0]))
    fmt.Println()
    diagonal(s[1:], k + 1)
    //fmt.Println()
    

}
func main() {
    var palavra string
    fmt.Scan(&palavra)

    diagonal(palavra, 0)
    // fmt.Println("Hello, World!")
}