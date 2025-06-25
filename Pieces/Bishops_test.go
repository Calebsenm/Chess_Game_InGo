package Pieces

import (
	"testing"
)

func TestBishopsMovesBasic(t *testing.T) {
	board := [9][9]string{}
	// Alfil blanco en el centro del tablero
	board[4][4] = " \u2657 "
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if board[i][j] == "" {
				board[i][j] = " - "
			}
		}
	}
	bishop := Bishops_{Y_: 4, X_: 4, Board_: board, Color: n_}
	moves := bishop.MovesBishop()
	if len(moves) == 0 {
		t.Errorf("El alfil debería tener movimientos disponibles en el centro del tablero")
	}
}
