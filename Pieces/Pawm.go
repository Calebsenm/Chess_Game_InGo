package Pieces

var (
	N = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}
)

type Pawm_ struct {
	Y      int
	X      int
	Board_ [9][9]string
	Color  [6]string
}

func (objet Pawm_) MovesCalculate() [][2]int {
	
	mv := movesPawm(objet.Y, objet.X, objet.Board_)
	return mv
}

func movesPawm(y int, x int, board [9][9]string) [][2]int {

	let_moves := [][2]int{}

	// chek color is white or black

	if colorChecker(board[y][x], B) {
		if y == 6 && board[y-1][x] == " - " && board[y-2][x] == " - " {

			let_moves = append(let_moves, [2]int{y - 1, x})
			let_moves = append(let_moves, [2]int{y - 2, x})
		} else if board[y-1][x] == " - " {
			let_moves = append(let_moves, [2]int{y - 1, x})
		}

		// para almacenar el calulo del ataque derecho
		_, k2 := checkEdge(y, x)

		if k2 == true && colorChecker(board[y-1][x-1], N) {
			let_moves = append(let_moves, [2]int{y - 1, x - 1})
		}
		if k2 == true && colorChecker(board[y-1][x+1], N) {
			let_moves = append(let_moves, [2]int{y + 1, x - 1})
		}

		// el calculo de las negras
	} else {
		if y == 1 && board[y+1][x] == " - " && board[y+2][x] == " - " {
			let_moves = append(let_moves, [2]int{y + 1, x})
			let_moves = append(let_moves, [2]int{y + 2, x})

		} else if y == 6 && board[y+1][x] == " - " {
			let_moves = append(let_moves, [2]int{y + 1, x})
		}

		_, k2 := checkEdge(y, x)
		if k2 == true && colorChecker(board[y+1][x+1], B) {
			let_moves = append(let_moves, [2]int{y + 1, x + 1})
		}
		if k2 == true && colorChecker(board[y-1][x+1], N) {
			let_moves = append(let_moves, [2]int{y - 1, x + 1})
		}

	}

	return let_moves
}
