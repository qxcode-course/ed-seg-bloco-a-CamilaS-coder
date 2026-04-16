package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(0, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(0, 0) // move o pincel para x 250, y 500
	pen.SetHeading(90)        // coloca o pincel apontando para cima
	pen.DrawRect(500, 500)
	pen.FillSquare(500, 500)
	pen.SetPosition(250, 500) // move o pincel para x 250, y 500
	pen.SetRGB(240, 248, 255)
	pen.Walk(125)

	i := 125.00

	for range 50{
		pen.Left(45)
		pen.Walk(i)
		pen.Right(90)
		i -= 25
	}
	pen.SetHeading(270)
	pen.Walk(125)
	//pen.DrawCircle(250)


	// pen.Walk(90)             // anda 100 pixels
	// pen.Left(35)              // dobra 30 graus para esquerda
	// pen.Walk(150)             // anda 100 pixels
	// pen.DrawCircle(0)        // desenha um círculo de raio 50
	// pen.Right(90)             // gira para direita 60 graus
	// pen.Walk(150)
	

	// for range 10 {

	// 	pen.Up()
	// 	pen.Walk(125) // anda sem riscar
	// 	// pen.Down()
	// 	// pen.Up()
	// 	// pen.FillSquare(10, 100)
	// 	// pen.Walk(-30) // volta sem riscar
	// 	// pen.Down()

	// 	// pen.Left(90) // gira
	// 	//pen.DrawRect(i, j)		
	// }

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
