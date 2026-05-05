package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move para a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.End() { //Se o cursor não está no final da linha
		e.cursor = e.cursor.Next() // move o cursor pra DIREITA
		return
		//e.cursor = e.line.Value.Front()
	}
	// Estamos no final da linha
	if e.line != e.lines.Back() { // Se não está na ultima linha
		e.line = e.line.Next() // move para a próxima linha
		e.cursor = e.line.Value.Front() // move o cursor para o inicio da proxima linha
	}
	
	
}

func (e *Editor) KeyUp() {
	//Se a linha não é a primeira linha, ou seja não é a primerira linha da lista de linhas
	if e.line != e.lines.Front() {
		e.line = e.line.Prev() // move para linha anterior
		e.cursor = e.line.Value.Front() //move o cursor para o inicio da linha anterior
		//return

	}

	// if e.cursor == e. line.Value.Back() { // O cursor está no final da linha
	// 	e.line = e.line.Prev()
	// 	e.cursor = e.line.Value.Front()
	// 	//return
	// }
}

func (e *Editor) KeyDown() {
	//Se a linha não é a ultima, ou seja não é a ultima linha da lista de linhas
	if e.line != e.lines.Back() {
		e.line = e.line.Next() // move para proxima linha linha debaixo
		e.cursor = e.line.Value.Front() // move o cursor para o começo da linha
	}


}

func (e *Editor) KeyEnter() {
	nova := NewList[rune]() //Cria uma linha nova

	e.lines.Insert(e.lines.root, nova) // SE não tiver proxima linha ou seja é a primeira linha

	//Se o cursor estiver no final da linha, ou seja depois do ultimo nó válido
	if e.cursor == e.line.Value.End()  { 
		e.line = e.line.Next() // Move para a proxima linha ou seja  linha debaixo
		e.cursor = e.line.Value.Front() // Move o cursor para o inicio da proxima linha
	}

	
}


func (e *Editor) KeyBackspace() {
	//Se o cursor não estiver no começo da linha
	if e.cursor != e.line.Value.Front() {
		e.line.Value.Erase(e.cursor.prev)
		return //apaga o caractere que está antes do cursor
	}

	// O cursor está no começo da linha
	if e.line != e.lines.Front(){
	//linhaAnterior := e.line.Prev()
	e.cursor = e.line.Prev().Value.End()
	}




}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.Back() {
		e.line.Value.Erase(e.cursor.next)
	}

}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}

