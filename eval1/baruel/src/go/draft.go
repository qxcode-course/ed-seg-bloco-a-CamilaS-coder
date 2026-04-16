package main
import "fmt"
func main() {
    var qtdfig, figBaruel int
    repetidas := make([]int, 0, figBaruel)
    unicas := make(map[int]bool)


    fmt.Scan(&qtdfig, &figBaruel)
    figVet := make([]int, figBaruel)

    for i := range figVet{
        fmt.Scan(&figVet[i])
        
       //unicas[i] = true
    }

    //fmt.Print(qtdfig, figBaruel, figVet, unicas)

    for _, fig := range figVet {
        
        _, existe := unicas[fig]

        if existe {
            repetidas = append(repetidas, fig)
        } else {

            unicas[fig] = true
             
        }

    }

    

    //fmt.Print(repetidas, unicas)
    sai := fmt.Sprintf("%v", repetidas)

    if sai == "[]" {
        sai = "N"
    } else {
        sai = sai[1:len(sai) - 1]
    }
    fmt.Println(sai)

    for i := 0; i < len(repetidas); i++ {
        
    }

    // if sai == "[]" {
    //     sai = "N"
    //     fmt.Print(sai)
    // } else {
    //     fmt.Print(sai)
    // }

    

    // fmt.Print(sai)
    
}
