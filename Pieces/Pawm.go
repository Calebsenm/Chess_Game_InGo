package Pieces

var (
	n_1 = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b_1 = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	x_1 int
	y_1 int

	listOfNumbersPawm = [][2]int{}
	board_1           [9][9]string
)

type Pawm_ struct {
	Y_1     int
	X_1     int
	Board_1 [9][9]string
	Color_1 [6]string
	Key_1   bool
}

func (pw *Pawm_) MovesPawm() [][2]int {

	var vacia [][2]int
	listOfNumbersPawm = vacia

	// this save variable
	x_1 = pw.X_1
	y_1 = pw.Y_1
	board_1 = pw.Board_1

	if pw.Color_1 == n_1 {
		PawmAlgoritmo(-1, 0)
	} else {
		PawmAlgoritmo(1, 0)
	}

	if pw.Key_1 == true {
		if pw.Color_1 == n_1 {
			PawmAlgoritmo(-1, -1)
			PawmAlgoritmo(-1, +1)
		} else {
			PawmAlgoritmo(1, -1)
			PawmAlgoritmo(1, +1)
		}

	}

	return listOfNumbersPawm
}

func PawmAlgoritmo(yChance_, xChance_ int) {

	atactPiece_1 := board_1[y_1][x_1]

	//fmt.Println(atactPiece____)
	y1_1 := y_1
	x1_1 := x_1

	y1_1 = y1_1 + yChance_
	x1_1 = x1_1 + xChance_

	if fueraRango__2(y1_1, x1_1) {

		// para los movimientos de ataque

		if y_1 == 1 || y_1 == 6 {

			if y_1 == 1 {
				if board_1[2][x1_1] == " - " && board_1[3][x1_1] == " - " {
					listOfNumbersPawm = append(listOfNumbersPawm, [2]int{2, x1_1})
					listOfNumbersPawm = append(listOfNumbersPawm, [2]int{3, x1_1})

				}
			}

			if y_1 == 6 {

				if board_1[5][x1_1] == " - " && board_1[4][x1_1] == " - " {
					listOfNumbersPawm = append(listOfNumbersPawm, [2]int{5, x1_1})
					listOfNumbersPawm = append(listOfNumbersPawm, [2]int{4, x1_1})

				}
			}
		}

		if board_1[y1_1][x1_1] == " - " {
			listOfNumbersPawm = append(listOfNumbersPawm, [2]int{y1_1, x1_1})
		}

		//fmt.Println(fichaAtaca____(atactPiece____))

		// white
		if fichaAtaca_1(atactPiece_1) == 1 {
			if fichas_1(board_1[y1_1][x1_1], n_1) {

				listOfNumbersPawm = append(listOfNumbersPawm, [2]int{y1_1, x1_1})
			}

		}
		// black
		if fichaAtaca_1(atactPiece_1) == 2 {
			if fichas_1(board_1[y1_1][x1_1], b_1) {
				listOfNumbersPawm = append(listOfNumbersPawm, [2]int{y1_1, x1_1})

			}

		}

	}
}

// out of the range
func fueraRango__2(y, x int) bool {
	keyValue := false

	if y >= 0 && y <= 7 && x >= 1 && x <= 8 {
		keyValue = true
	}
	return keyValue
}

// what pice is fithg
func fichaAtaca_1(atacaX_1 string) int {

	var num int

	for i := 0; i < 6; i++ {
		//fmt.Println(atacaX____ ,"  - ", " " + b____[i] + " ")
		if atacaX_1 == " "+b_1[i]+" " {
			num = 1
			break

		} else if atacaX_1 == " "+n_1[i]+" " {
			num = 2
			break

		}
	}

	return num
}

// to control the opponent's chip attack
func fichas_1(ficha_1 string, list_1 [6]string) bool {

	for i := 0; i < len(list_1); i++ {
		//fmt.Println( " "+list____[i]+" ", " - ", ficha____)
		if " "+list_1[i]+" " == ficha_1 {

			return true
		}
	}
	return false
}

// package Pieces

// var (
// 	N = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
// 	B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}
// )

// type Pawm_ struct {
// 	Y      int
// 	X      int
// 	Board_ [9][9]string
// 	Color  [6]string
// 	Llave  bool
// }

// func (objet Pawm_) MovesCalculate() [][2]int {

// 	mv := movesPawm(objet.Y, objet.X, objet.Board_ , objet.Llave )
// 	return mv
// }

// func movesPawm(y int, x int, board [9][9]string , keyPawm bool ) [][2]int {

// 	let_moves := [][2]int{}

// 	// chek color is white or black

// 	if colorChecker(board[y][x], B) {
// 		if y == 6 && board[y-1][x] == " - " && board[y-2][x] == " - " {

// 			let_moves = append(let_moves, [2]int{y - 1, x})
// 			let_moves = append(let_moves, [2]int{y - 2, x})
// 		} else if board[y-1][x] == " - " {
// 			let_moves = append(let_moves, [2]int{y - 1, x})
// 		}

// 		// para almacenar el calulo del ataque derecho
// 		_, k2 := checkEdge(y, x)

// 		if k2 == true && colorChecker(board[y-1][x-1], N)  {
// 			let_moves = append(let_moves, [2]int{y - 1, x - 1})
// 		}

// 		if k2 == true && colorChecker(board[y-1][x+1], N)  {
// 			let_moves = append(let_moves, [2]int{y + 1, x - 1})
// 		}

// 		if   k2 && keyPawm{
// 			let_moves = append(let_moves, [2]int{y - 1, x - 1})
// 			let_moves = append(let_moves, [2]int{y + 1, x - 1})
// 		}

// 		// el calculo de las negras
// 	} else {
// 		if y == 1 && board[y+1][x] == " - " && board[y+2][x] == " - " {
// 			let_moves = append(let_moves, [2]int{y + 1, x})
// 			let_moves = append(let_moves, [2]int{y + 2, x})

// 		} else if y == 6 && board[y+1][x] == " - " {
// 			let_moves = append(let_moves, [2]int{y + 1, x})
// 		}

// 		_, k2 := checkEdge(y, x)
// 		if k2 == true && colorChecker(board[y+1][x+1], B) {
// 			let_moves = append(let_moves, [2]int{y + 1, x + 1})
// 		}
// 		if k2 == true && colorChecker(board[y-1][x+1], N)  {
// 			let_moves = append(let_moves, [2]int{y - 1, x + 1})
// 		}

// 		if   k2 && keyPawm{
// 			let_moves = append(let_moves, [2]int{y - 1, x - 1})
// 			let_moves = append(let_moves, [2]int{y + 1, x - 1})
// 		}
// 	}

// 	return let_moves
// }
