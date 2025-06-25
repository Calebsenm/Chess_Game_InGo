package Pieces

import (
	"testing"
)

func TestKingMovesBasic(t *testing.T) {
	board := [9][9]string{}
	// Rey blanco en el centro del tablero
	board[4][4] = " \u2654 "
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if board[i][j] == "" {
				board[i][j] = " - "
			}
		}
	}
	// Simular que todas las casillas alrededor del rey no están bajo ataque
	var safeMoves [][2]int
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy != 0 || dx != 0 {
				y := 4 + dy
				x := 4 + dx
				if y >= 1 && y <= 8 && x >= 1 && x <= 8 {
					safeMoves = append(safeMoves, [2]int{y, x})
				}
			}
		}
	}
	king := King_{Y_____: 4, X_____: 4, Board_____: board, Color_____: n_____, ListFight1_____: safeMoves, ListFight2_____: safeMoves}
	moves := king.MovesKing()
	if len(moves) == 0 {
		t.Errorf("El rey debería tener movimientos disponibles en el centro del tablero")
	}
}
