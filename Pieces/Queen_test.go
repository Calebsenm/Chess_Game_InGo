package Pieces

import (
	"testing"
)

func TestQueenMovesBasic(t *testing.T) {
	board := [9][9]string{}
	// Reina blanca en el centro del tablero
	board[4][4] = " \u2655 "
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if board[i][j] == "" {
				board[i][j] = " - "
			}
		}
	}
	queen := Queen_{Y_: 4, X_: 4, Board_: board, Color: n__}
	moves := queen.MovesQueen()
	if len(moves) == 0 {
		t.Errorf("La reina debería tener movimientos disponibles en el centro del tablero")
	}
}
