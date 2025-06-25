package Pieces

import (
	"testing"
)

func TestPawmMovesBasic(t *testing.T) {
	board := [9][9]string{}
	// Inicializar la fila 6 (índice 6) con peones blancos y la fila 5 y 4 vacías
	for i := 1; i <= 8; i++ {
		board[6][i] = " \u2659 " // Peón blanco
		board[5][i] = " - "
		board[4][i] = " - "
	}
	// Peón blanco en posición inicial (6,4)
	pawm := Pawm_{Y_1: 6, X_1: 4, Board_1: board, Color_1: n_1, Key_1: false}
	moves := pawm.MovesPawm()
	if len(moves) == 0 {
		t.Errorf("El peón debería tener movimientos disponibles en la posición inicial")
	}
}
