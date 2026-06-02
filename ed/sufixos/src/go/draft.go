package main
import "fmt"

func imprimindo(s string, k int) {
    if len(s) == 0 {
        return
    }

    // fmt.Print(string(s))
    imprimindo(s[1:], k + 1)
    // imprimindo(s[1:], k + 1)
    fmt.Print(string(s))
    fmt.Println()
}
func main() {
    var palavra string
    fmt.Scan(&palavra)

    imprimindo(palavra, 0)
    //fmt.Println("Hello, World!")
}