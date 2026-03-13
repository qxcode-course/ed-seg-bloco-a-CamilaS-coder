package main
import "fmt"
func main() {
    var H, P, F, D int
    
    fmt.Scan(&H, &P, &F, &D)

    var i = F

    for {
        i += D

        var p = (i + 16) % 16

        if p == H {
            fmt.Println("S")
            break
        } else if p == P {
            fmt.Println("N")
            break
        }



    }
    
    // if D == 1 {
    //     F++
    //     if F == 16 {
    //         F = 0
    //     }

    //     if F == P {
    //         fmt.Println("N")
    //     } else {
    //         fmt.Println("S")
    //     }
        
    // } else if D == -1{
    //     F--
    //     if F == 0 {
    //         F = 16
    //     }

    //     if F == P {
    //         fmt.Println("N")
    //     } else {
    //         fmt.Println("S")
    //     }
    // }

    //fmt.Println("qxcode")
}
