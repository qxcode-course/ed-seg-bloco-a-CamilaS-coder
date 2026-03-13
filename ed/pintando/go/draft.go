package main
import "fmt"
import "math"
func main() {
    var a, b, c float32
    fmt.Scan(&a, &b, &c)

    var p = float32(a + b + c) / float32(2)

    var A = p * (p - float32(a)) * (p - float32(b)) * (p - float32(c))

    var Area = math.Sqrt(float64(A))
    fmt.Printf("%.2f\n", Area)
}
