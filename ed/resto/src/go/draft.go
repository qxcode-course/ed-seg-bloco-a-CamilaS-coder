package main
import "fmt"

func div(v int) {
    if v == 0 && v % 2 == 0{
        return
    }
    vAlt := 0
    vAlt = v / 2

    div(vAlt)
    fmt.Print(vAlt, " ")
    fmt.Println(v % 2)
    v = vAlt

   
}

func main() {
    var v int
    fmt.Scan(&v)
    div(v)
}
