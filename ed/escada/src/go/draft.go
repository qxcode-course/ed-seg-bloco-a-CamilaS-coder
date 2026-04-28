package main
import "fmt"

func op_degraus(a int) int {

    if a == 1 || a == 2 {
        return 1
    }

    if a == 3 {
        return 2
    }


    return op_degraus(a - 1) + op_degraus(a - 3)
}

func main() {
    var a int
    fmt.Scan(&a)

    //op_degraus(a)
    fmt.Println(op_degraus(a))
}
