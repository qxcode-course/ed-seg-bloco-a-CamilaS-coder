package main
import "fmt"
func main() {

    var figAlb, figBaruel int
    fmt.Scan(&figAlb, &figBaruel)
    var i int
    var j int
    var k int
    var apareceu bool
    var figVet []int
    primeiro := true
    var contador int
    var aux []int
    //var aux2 []int
    
    for i = 0; i < figBaruel; i++{
        var fig int
        fmt.Scan(&fig)
        figVet = append(figVet, fig)
        
    }
    
    for i = 0; i < len(figVet); i++ {
        apareceu = false
        for  j = 0; j <= i - 1; j++{
            if figVet[i] == figVet[j] {
                apareceu = true
                break
            }
            
        } 
            
        if !apareceu{
            contador := 0

            for k = 0; k < len(figVet); k++{
            if figVet[i] == figVet[k] {
                //aux = append(aux, figVet[k])
                 contador++
                }
            }
            extras := contador - 1

             for x := 0; x < extras; x++ {

                if !primeiro {
                    fmt.Print(" ")
                }

                fmt.Print(figVet[i])
                primeiro = false
            }

        }
        
        
    }

    fmt.Println()

    for y := 0; y < figAlb; y++ {
        contador++
        aux = append(aux, contador)
    }

    //fmt.Print(aux)


    // for i = 0; i < len(figVet); i++ {
    //     for j = i+1; j < len(figVet); j++ {
    //         if figVet[i] == figVet[j] {
    //             aux = append(aux, figVet[i])
    //         }
    //     }
    // }



}
