package Functions


// esta funcion va a calcular todos los posibles movimientos
/*
Su funcionamiento es un bucle que recorre el tablero  comparando 
el color de cada ficha para luego calcular los movimientos respectivos de cada ficha 

EL primer retorno es de la lista de los movientos de las fichas Blancas 
El segundo retorno es de la lista de los movientos de las fichas Negras 
*/

// se importa el modulo chess/Pieces donde se encuentra cada ficha y su algoritmo el cual retorna 
// los movimientos permitidos de la ficha para cada  posicion y la ficha 

import (
	"chess/Pieces"
)


func calcularTodosLosMovimietos(Board [9][9] string) ([][2]int , [][2]int) {

	var listaMovimientosB [][2]int
	var listaMovimientosN [][2]int
	var movimientosPosiblesBlancos [][2]int
	var movimientosPosiblesNegros  [][2]int

					// pawm     rook       queen    knight     king      bishop
	var N = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	var B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	
	for i := 0; i < len(Board); i++ {
		for j := 0; j < len(Board); j++ {

			y := i
			x := j

			if Board[i][j] == " "+N[0]+" " || Board[i][j] == " "+N[1]+" " || Board[i][j] == " "+N[2]+" " || Board[i][j] == " "+N[3]+" " || Board[i][j] == " "+N[4]+" " || Board[i][j] == " "+N[5]+" " {

				switch Board[i][j] {

				case " " + N[0] + " ":
					pawm := Pieces.Pawm_{y, x, Board, B, true}
					listaMovimientosN = pawm.MovesPawm()

					for _, element := range listaMovimientosN {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, element)
					}

				case " " + N[1] + " ":
					rook  := Pieces.Rook_{y, x, Board, B}
					listaMovimientosN = rook.MovesRook();

					for _, element := range listaMovimientosN {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, element)
					}

				case " " + N[2] + " ":
					queen := Pieces.Queen_{y, x, Board, B}
					listaMovimientosN = queen.MovesQueen()

					for _, element := range listaMovimientosN {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, element)
					}


				case " " + N[3] + " ":
					knight := Pieces.Knight_{y, x, Board, B}
					listaMovimientosN = knight.MovesKnight()

					for _, element := range listaMovimientosN {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, element)
					}

				case " " + N[4] + " ":
					king := Pieces.King_{y, x, Board, B, movimientosPosiblesNegros, movimientosPosiblesBlancos}
					listaMovimientosN = king.MovesKing();

					for _, elem := range listaMovimientosN {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, elem)
					}

				case " " + N[5] + " ":
					bishops := Pieces.Bishops_{y, x, Board, B}
					listaMovimientosN = bishops.MovesBishop();

					for _, element := range listaMovimientosN {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, element)
					}
				}


			} else if Board[i][j] == " "+B[0]+" " || Board[i][j] == " "+B[1]+" " || Board[i][j] == " "+B[2]+" " || Board[i][j] == " "+B[3]+" " || Board[i][j] == " "+B[4]+" " || Board[i][j] == " "+B[5]+" " {

				switch Board[i][j] {

				case " " + B[0] + " ":
					pawm := Pieces.Pawm_{y, x, Board, N, true}
					listaMovimientosB = pawm.MovesPawm();

					for _, element := range listaMovimientosB {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, element)
					}

				case " " + B[1] + " ":
					rook := Pieces.Rook_{y, x, Board, N}
					listaMovimientosB = rook.MovesRook();

					for _, element := range listaMovimientosB {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, element)
					}

				case " " + B[4] + " ":
					king := Pieces.King_{y, x, Board, N, movimientosPosiblesBlancos, movimientosPosiblesNegros}
					listaMovimientosB = king.MovesKing();
					for _, element := range listaMovimientosB {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, element)
					}

				case " " + B[3] + " ":
					knight:= Pieces.Knight_{y, x, Board, N}
					listaMovimientosB = knight.MovesKnight();

					for _, element := range listaMovimientosB{
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, element);
					}

				case " " + B[2] + " ":
					queen := Pieces.Queen_{y, x, Board, N}
					listaMovimientosB = queen.MovesQueen();

					for _, element := range listaMovimientosB {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, element)
					}

				case " " + B[5] + " ":
					bishop := Pieces.Bishops_{y, x, Board, N}
					listaMovimientosB =bishop.MovesBishop();

					for _, elem := range listaMovimientosB {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, elem)
					}
				}

			}
		}
	}

	return movimientosPosiblesBlancos ,  movimientosPosiblesNegros;
}



