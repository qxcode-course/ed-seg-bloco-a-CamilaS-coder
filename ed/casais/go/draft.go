package main

import (
	"fmt"
	//"math"
)
func main() {

    var n int
    var i, j int
    
    fmt.Scan(&n)

    var descasados []int
    var casais = 0

    for i = 0; i < n; i++ {
        var animais int
        fmt.Scan(&animais)

        encontrou := false

        for j = 0; j < len(descasados); j++ {
            if descasados[j] == -animais {
                casais++
                descasados[j] = 0 
                encontrou = true
                break
            } 
        }

        if !encontrou {
             descasados = append(descasados, animais)
        }
    }

    //animais := make([]int, n)


    fmt.Println(casais)
    
}
