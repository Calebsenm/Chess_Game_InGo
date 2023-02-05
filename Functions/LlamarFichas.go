package Functions

import ( 
	"chess/Pieces"
	"fmt"

);

var(
	N = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}
	Board [9][9] string

);


// this is for calls the moves to other pieces 
func LlamarMovimientoFichas(player string, colorContrario [6]string , board [9][9]string ) ( [9][9] string ) {

	// update the board
	Board =  board;

	//llama el  movimiento de juego 
	y, x := CambioLetraNumero(player, " Escoge una posicion > ");
	
	veri := false

	
	ix2 := false
	var movimientosPosiblesBlancos, movimientosPosiblesNegros   = calcularTodosLosMovimietos(Board);
	
	NumeroIteraciones := 0;
	for {

		for {
			if Board[y][x] == " - " || Board[y][x] == " "+colorContrario[0]+" " || Board[y][x] == " "+colorContrario[1]+" " || Board[y][x] == " "+colorContrario[2]+" " || Board[y][x] == " "+colorContrario[3]+" " || Board[y][x] == " "+colorContrario[4]+" " || Board[y][x] == " "+colorContrario[5]+" " {
				fmt.Println("Has elegido un movimiento no permitido")
				y, x = CambioLetraNumero("1 ", "Escoge una posicion")
			} else {
				break
			}
		}

		var listaDePosiciones [][2]int
		//Pawm
		if Board[y][x] == " "+B[0]+" " || Board[y][x] == " "+N[0]+" " {
			fmt.Println(Board[y][x])

			v1 := Pieces.Pawm_{y, x, Board, colorContrario, false}
			listaDePosiciones = v1.MovesPawm()

		}
		//Bishops
		if Board[y][x] == " "+B[5]+" " || Board[y][x] == " "+N[5]+" " {
			fmt.Println(Board[y][x])
			v2 := Pieces.Bishops_{y, x, Board, colorContrario}
			listaDePosiciones = v2.MovesBishop()

		}

		//Queen
		if Board[y][x] == " "+B[2]+" " || Board[y][x] == " "+N[2]+" " {

			fmt.Println(Board[y][x])
			v3 := Pieces.Queen_{y, x, Board, colorContrario}

			listaDePosiciones = v3.MovesQueen()
		}

		//Rook
		if Board[y][x] == " "+B[1]+" " || Board[y][x] == " "+N[1]+" " {

			fmt.Println(Board[y][x])
			v4 := Pieces.Rook_{y, x, Board, colorContrario}
			listaDePosiciones = v4.MovesRook()
		}
		// Knight
		if Board[y][x] == " "+B[3]+" " || Board[y][x] == " "+N[3]+" " {

			fmt.Println(Board[y][x])
			v5 := Pieces.Knight_{y, x, Board, colorContrario}
			listaDePosiciones = v5.MovesKnight()
		}

		llave_Rey := false
		// King
		if Board[y][x] == " "+B[4]+" " || Board[y][x] == " "+N[4]+" " {

			fmt.Println(Board[y][x])
			v6 := Pieces.King_{y, x, Board, colorContrario, movimientosPosiblesNegros, movimientosPosiblesBlancos}

			listaDePosiciones = v6.MovesKing()
			llave_Rey = Pieces.TheKey
		}


		if len(listaDePosiciones) == 0 {
			fmt.Println("La ficha està bloqueada")
			y, x = CambioLetraNumero(" - ", "Escoge una posicion -> ")

		}

		if len(listaDePosiciones) != 0 {
			
			Y, X := CambioLetraNumero("", " Donde deseas moverte > ");

			for i := 0; i < len(listaDePosiciones); i++ {
				if listaDePosiciones[i][0] == Y && listaDePosiciones[i][1] == X {
					veri = true
					break
				}
			}

			if veri{
				// para controlar el enrrosque , esto va a verificar si la ficha es un rey y el moviminto

				if (Board[y][x] == " "+B[4]+" " || Board[y][x] == " "+N[4]+" ") && (llave_Rey == true) && (X == 3 || X == 7) {
					Board[Y][X] = Board[y][x]

					if X == 3 {
						Board[y][x-1] = Board[y][x-4]
						Board[y][x] = " - "
						Board[y][x-4] = " - "

					} else if X == 7 {

						Board[y][x+1] = Board[y][x+3]
						Board[y][x] = " - "
						Board[y][x+3] = " - "
					}

				} else {

					Board[Y][X] = Board[y][x]
					Board[y][x] = " - "
				}
				ix2 = true
			}

			if ix2 == true {
				break
			}
			
			NumeroIteraciones ++;
		}

	}



	
	return Board;
}






// // para validar si los numeros estan en la lista
// func verificarNumeroPosicion(num1_ int, num2_ int, lisNum1_ [][2]int) bool {
// 	validation := false

// 	for i := 0; i < len(lisNum1_); i++ {
// 		if num1_ == lisNum1_[i][1] && num2_ == lisNum1_[i][0] {
// 			validation = true
// 		}
// 	}
// 	return validation

// }