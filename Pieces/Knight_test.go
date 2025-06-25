package Pieces

import (
	"testing"
)

func TestKnightMovesBasic(t *testing.T) {
	board := [9][9]string{}
	// Caballo blanco en el centro del tablero
	board[4][4] = " \u2658 "
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if board[i][j] == "" {
				board[i][j] = " - "
			}
		}
	}
	knight := Knight_{Y____: 4, X____: 4, Board____: board, Color____: n____}
	moves := knight.MovesKnight()
	if len(moves) == 0 {
		t.Errorf("El caballo debería tener movimientos disponibles en el centro del tablero")
	}
}
