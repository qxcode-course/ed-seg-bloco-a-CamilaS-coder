package main
import "fmt"
func main() {

    var figAlb, figBaruel int
    fmt.Scan(&figAlb, &figBaruel)
    //valor := 0
    figVet := make([]int, figBaruel)
    unicas := make(map[int]bool)
    repetidos := make([]int, 0, figBaruel)
    
    for i := range figVet {
        fmt.Scan(&figVet[i])
    }
    
    for _, fig := range figVet {

        _, existe := unicas[fig]
        if existe {
            repetidos = append(repetidos, fig)
        } else {
            unicas[fig] = true
        }

    }

    saida := fmt.Sprintf("%v", repetidos)
    if saida == "[]" {
        fmt.Println("N")
    } else {

        fmt.Println(saida[1: len(saida) - 1])
    }


    faltantes := make([]int, 0, figAlb)

    for i := 1; i <= figAlb; i++ {
        if !unicas[i] {
            faltantes = append(faltantes, i)
        }
    }

    
    sai := fmt.Sprintf("%v", faltantes)

    if sai == "[]" {
        fmt.Println("N")
    } else {

        fmt.Println(sai[1:len(sai) - 1])
    }


    

    // var figAlb, figBaruel int
    // fmt.Scan(&figAlb, &figBaruel)
    // var i int
    // var j int
    // var k int
    // var apareceu bool
    // var figVet []int
    // primeiro := true
    // var contador int
    // var aux []int
    // //var aux2 []int
    
    // for i = 0; i < figBaruel; i++{
    //     var fig int
    //     fmt.Scan(&fig)
    //     figVet = append(figVet, fig)
        
    // }
    
    // for i = 0; i < len(figVet); i++ {
    //     apareceu = false
    //     for  j = 0; j <= i - 1; j++{
    //         if figVet[i] == figVet[j] {
    //             apareceu = true
    //             break
    //         }
            
    //     } 
            
    //     if !apareceu{
    //         contador := 0

    //         for k = 0; k < len(figVet); k++{
    //         if figVet[i] == figVet[k] {
    //             //aux = append(aux, figVet[k])
    //              contador++
    //             }
    //         }
    //         extras := contador - 1

    //          for x := 0; x < extras; x++ {

    //             if !primeiro {
    //                 fmt.Print(" ")
    //             }

    //             fmt.Print(figVet[i])
    //             primeiro = false
    //         }

    //     }
        
        
    // }

    // fmt.Println()

    // for y := 0; y < figAlb; y++ {
    //     contador++
    //     aux = append(aux, contador)
    // }

    // for i := 0; i < len(figVet); i++ {
    //     for j := 0; j < len(aux); j++ {
    //         if figVet[i] != aux[j] {
    //             fmt.Print(aux[j])
    //             break
    //         }
    //     }

    // }



}
