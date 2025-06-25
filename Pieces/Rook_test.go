package Pieces

import (
	"testing"
)

func TestRookMovesBasic(t *testing.T) {
	board := [9][9]string{}
	// Torre blanca en el centro del tablero
	board[4][4] = " \u2656 "
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if board[i][j] == "" {
				board[i][j] = " - "
			}
		}
	}
	rook := Rook_{Y___: 4, X___: 4, Board___: board, Color___: n___}
	moves := rook.MovesRook()
	if len(moves) == 0 {
		t.Errorf("La torre debería tener movimientos disponibles en el centro del tablero")
	}
}
