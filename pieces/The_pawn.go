
package pieces
// Function to calculate the color of the pice to eat
func check(A [6]string, B string) bool {

	xddd := false
	for i := 0; i < len(A); i++ {
		if A[i] == B {
			xddd = true
		}

	}
	return xddd
}

var (
	piecesToatac [6]string
	to_move      int

	allowedMoves = map[int]position1{}
)

type position1 struct {
	A, B int
}

type The_pawn struct {
	X1    int
	Y1    int
	white [6]string
	black [6]string

	board [9][9]string
}

func (moves *The_pawn) Allowed_moves() {

	boardI := moves.board

	// if the pice position is white
	if boardI[moves.Y1][moves.X1] == moves.white[0] {
		piecesToatac = moves.black
		to_move = -1
	}
	// if the pice position is black
	if boardI[moves.Y1][moves.X1] == moves.black[0] {
		piecesToatac = moves.white

		to_move = 1

	}

	//the first move
	if moves.Y1+to_move == 8 && boardI[moves.Y1][moves.X1] == moves.white[0] {
		// the position 2 front
		if boardI[moves.Y1+(to_move*2)][moves.X1] == " - " {
			allowedMoves[len(allowedMoves)] = position1{moves.Y1 + (to_move * 2), moves.X1}
		}
	} else if moves.Y1+to_move == 1 && boardI[moves.Y1][moves.X1] == moves.black[0] {
		// the position 2 front
		if boardI[moves.Y1+(to_move*2)][moves.X1] == " - " {
			allowedMoves[len(allowedMoves)] = position1{moves.Y1 + (to_move * 2), moves.X1}
		}
	}

	// the other movements   the next
	//the first move  for the whithe pieces
	if boardI[moves.Y1][moves.X1] == moves.white[0] {
		// the position 2 front
		if boardI[moves.Y1+(to_move*1)][moves.X1] == " - " {
			allowedMoves[len(allowedMoves)+1] = position1{moves.Y1 + (to_move * 1), moves.X1}

		}

		if moves.X1+(to_move*1) >= 0 {
			// left
			if check(moves.black, boardI[moves.Y1+(to_move*1)][moves.X1+(to_move*1)]) == true {
				allowedMoves[len(allowedMoves)+1] = position1{moves.Y1 + (to_move * 1), moves.X1 + (to_move * 1)}

			}
		}

		if moves.X1+(to_move*1) <= 8 {
			// left
			// right
			if check(moves.black, boardI[moves.Y1+(to_move*1)][moves.X1+(to_move*-1)]) {
				allowedMoves[len(allowedMoves)+1] = position1{moves.Y1 + (to_move * 1), moves.X1 + (to_move * -1)}

			}

		}



		// this is for the black pieces
	} else if boardI[moves.Y1][moves.X1] == moves.black[0] {
		// the position front
		if boardI[moves.Y1+(to_move*1)][moves.X1] == " - " {
			allowedMoves[len(allowedMoves)+1] = position1{moves.Y1 + (to_move * 1), moves.X1}
		}

		if moves.X1+(to_move*1) >= 0 {
			// left
			if check(moves.white, boardI[moves.Y1+(to_move*1)][moves.X1+(to_move*1)]) == true {
				allowedMoves[len(allowedMoves)+1] = position1{moves.Y1 + (to_move * 1), moves.X1 + (to_move * 1)}

			}
		}

		if moves.X1+(to_move*1) <= 8 {
			// left
			// right
			if check(moves.white, boardI[moves.Y1+(to_move*1)][moves.X1+(to_move*-1)]) {
				allowedMoves[len(allowedMoves)+1] = position1{moves.Y1 + (to_move * 1), moves.X1 + (to_move * -1)}

			}

		}

	}

}
