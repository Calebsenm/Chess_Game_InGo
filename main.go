// this will be the most bigest proyect
// i dot know how to create the game but i going to build the ches game with ia

// i start this proyect in the date  2/10/2022
// i delete the proyect in nov 21 / 2022
// i start the proyect in  nov 21 / 2022

package main

import (
	"chess/Pieces"
	"fmt"
	"strconv"
	"strings"
)

var (
	N = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	// Board = [9][9]string{

	// 	{" 8 ", " " + N[1] + " ", " " + N[3] + " ", " " + N[5] + " ", " " + N[2] + " ", " " + N[4] + " ", " " + N[5] + " ", " " + N[3] + " ", " " + N[1] + " "},
	// 	{" 7 ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " "},
	// 	{" 6 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 4 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 2 ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " " ," " + B[0] + " "},
	// 	{" 1 ", " " + B[1] + " ", " " + B[3] + " ", " " + B[5] + " ", " " + B[2] + " ", " " + B[4] + " ", " " + B[5] + " ", " " + B[3] + " ", " " + B[1] + " "},
	// 	{"   ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H "},
	// }

	Board = [9][9]string{

		{" 8 ", " " + N[1] + " ", " - ", " - ", " - "," " + N[4] + " ", " - ", " - ", " " + N[1] + " "},
		{" 7 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 6 ", " - ", " - ", " - ", " " + B[2] + " ", " - ", " - ", " - ", " - "},
		{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 4 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 2 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 1 ",  " " + B[1] + " ", " - ", " - ", " - "," " +B[4]+" ", " - ", " - ",  " " + B[1] + " "},
		{"   ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H "},
	}
	Board__ = [9][9]string{

		{" 8 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 7 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 6 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 4 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 2 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 1 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{"   ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H "},
	}

	TheYposition int
	TheXposition int

	TheXpositionToMove int
	TheYpositionToMove int

	allowedMoves = map[int]positions{}

	movimientosPosiblesBlancos [][2]int
	movimientosPosiblesNegros  [][2]int
)

// this is the strct for the moves
type positions struct {
	A, B int
}

// funtion for printBoard
func printBoard(lastChoise int) {

	if lastChoise == 2 {
		for i := 8; i > -1; i-- {
			for j := 8; j > -1; j-- {
				fmt.Print(Board[i][j])
			}
			fmt.Println()
		}
	} else if lastChoise == 1 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {

				fmt.Print(Board[i][j])
			}
			fmt.Println()
		}
    
	}

}

// this funtion going return the coordinates x y
func logicChoises(player_name, option string) (int, int) {

	movimientosPosiblesBlancos  = nil
	movimientosPosiblesNegros  =  nil
	values := [2]int{}
	for {
		var position1 string

		fmt.Print("Jugador " + player_name + option + " > ")
		fmt.Scanln(&position1)

		if len(position1) == 2 {
			// this is for separate the string
			s := position1
			v := strings.SplitAfter(s, "")

			a := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
			b := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

			choise, _ := strconv.Atoi(v[1])

			var key bool
			var key1 bool

			for io := 0; io < 8; io++ {
				if a[io] == v[0] {

					key = true
				}

			}

			for io := 0; io < 8; io++ {
				if b[io] == choise {

					key1 = true
				}

			}

			if key == true && key1 == true {
				// this is for x
				Change_letters_numbers := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
				// this is for y
				Change_numbers_numbers := map[string]int{"1": 7, "2": 6, "3": 5, "4": 4, "5": 3, "6": 2, "7": 1, "8": 0}

				values[0] = Change_numbers_numbers[v[1]]
				values[1] = Change_letters_numbers[v[0]]

				break
			}
		}

	}

	return values[0], values[1]

}

// funcion para verificar la ficha permitida

// this funtion going to start the game

func main() {
	fmt.Println("Wecolcome To chess game")
	var choice int

	for {
		fmt.Print("Please input " + "\n 1 for white  pieces \n 2 for black pieces \n -> ")

		fmt.Scan(&choice)

		if choice == 1 || choice == 2 {
			break
		}
	}

	iterator := 1

	colorContrario := [6]string{}

	for {
		iterator++

		fmt.Println("-------------------------------------------------------")
		printBoard(choice)
		fmt.Println("-------------------------------------------------------")

		if choice == 1 {
			colorContrario = N
		} else if choice == 2 {
			colorContrario = B
		}

		if iterator%2 == 0 {

			if choice == 1 {
				choice = 2
			} else {
				choice = 1
			}

			llamarMovimientosLogicos("1", colorContrario)

		} else {
			if choice == 2 {
				choice = 1
			} else {
				choice = 2
			}
			llamarMovimientosLogicos("2", colorContrario)
		}

	}

}

// esta funcion va a calcular todos los posibles movimientos
func calcularTodosLosMovimietos( ) {

	var val [][2]int
	var val1 [][2]int

	for i := 0; i < len(Board); i++ {
		for j := 0; j < len(Board); j++ {

			y := i
			x := j
			
			if Board[i][j] == " "+N[0]+" " || Board[i][j] == " "+N[1]+" " || Board[i][j] == " "+N[2]+" " || Board[i][j] == " "+N[3]+" " || Board[i][j] == " "+N[4]+" " || Board[i][j] == " "+N[5]+" " {
				

				switch Board[i][j] {

				case " " + N[0] + " ":
					value_ := Pieces.Pawm_{y, x, Board, B, true}
					val = value_.MovesPawm()

					for _, elem := range val {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, elem)
					}

					

				case " " + N[1] + " ":
					value_ := Pieces.Rook_{y, x, Board, B}
					val = value_.MovesRook()

					for _, elem := range val {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, elem)
					}

				case " " + N[2] + " ":
					value_ := Pieces.King_{y, x, Board, B}
					val = value_.MovesKing()
					for _, elem := range val {
						movimientosPosiblesNegros = append(movimientosPosiblesNegros, elem)
					}

				case " " + N[3] + " ":
					value_ := Pieces.Knight_{y, x, Board, B}
                    val = value_.MovesKnight()

                    for _, elem := range val {
                        movimientosPosiblesNegros = append(movimientosPosiblesNegros, elem)
                    }


				case " " + N[4] + " ":
					value_ := Pieces.King_{y, x, Board, B}
                    val = value_.MovesKing()

                    for _, elem := range val {
                        movimientosPosiblesNegros = append(movimientosPosiblesNegros, elem)
                    }

				case " " + N[5] + " ":
					value_ := Pieces.Bishops_{y, x, Board, B}
                    val = value_.MovesBishop()
					
                    for _, elem := range val {
                        movimientosPosiblesNegros = append(movimientosPosiblesNegros, elem)
                    }

				}

				 


			} else if Board[i][j] == " "+B[0]+" " || Board[i][j] == " "+B[1]+" " || Board[i][j] == " "+B[2]+" " || Board[i][j] == " "+B[3]+" " || Board[i][j] == " "+B[4]+" " || Board[i][j] == " "+B[5]+" " {


				

				switch Board[i][j] {

				case " " + B[0] + " ":
					value_ := Pieces.Pawm_{y, x, Board, N, true}
					val1 = value_.MovesPawm()
			
					for _, elem := range val1 {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, elem)
					}

			
				case " " + B[1] + " ":
					value_ := Pieces.Rook_{y, x, Board, N}
					val1 = value_.MovesRook()
			
					for _, elem := range val1 {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, elem)
					}
			
				case " " + B[2] + " ":
					value_ := Pieces.King_{y, x, Board, N}
					val1 = value_.MovesKing()
					for _, elem := range val1 {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, elem)
					}
			
				case " " + B[3] + " ":
					value_ := Pieces.Knight_{y, x, Board, N}
					val1 = value_.MovesKnight()
			
					for _, elem := range val1 {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, elem)
					}
			
			
				case " " + B[4] + " ":
					value_ := Pieces.King_{y, x, Board, N}
					val1 = value_.MovesKing()
			
					for _, elem := range val1 {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, elem)
					}
			
				case " " + B[5] + " ":
					value_ := Pieces.Bishops_{y, x, Board, N}
					val1 = value_.MovesBishop()
					
					for _, elem := range val1 {
						movimientosPosiblesBlancos = append(movimientosPosiblesBlancos, elem)
					}
				}
				 
			}
		}
	}

}

// this is for calls the moves to other pieces

func llamarMovimientosLogicos(player string, colorContrario [6]string) {

	var a [][2]int
	y, x := logicChoises(player, "Escoge una posicion")
	veri := false

	ix := 0
	ix2 := false

	for {

		for {
			if Board[y][x] == " - " || Board[y][x] == " "+colorContrario[0]+" " || Board[y][x] == " "+colorContrario[1]+" " || Board[y][x] == " "+colorContrario[2]+" " || Board[y][x] == " "+colorContrario[3]+" " || Board[y][x] == " "+colorContrario[4]+" " || Board[y][x] == " "+colorContrario[5]+" " {
				fmt.Println("Has elegido un movimiento no permitido")
				y, x = logicChoises("1 ", "Escoge una posicion")
			} else {
				break
			}
		}

		var vacia [][2]int
		a = vacia
		Board__ = [9][9]string{

			{" 8 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
			{" 7 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
			{" 6 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
			{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
			{" 4 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
			{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
			{" 2 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
			{" 1 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
			{"   ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H "},
		}

		//Pawm
		if Board[y][x] == " "+B[0]+" " || Board[y][x] == " "+N[0]+" " {
			fmt.Println(Board[y][x])

			v1 := Pieces.Pawm_{y, x, Board, colorContrario,false}
			a = v1.MovesPawm()

		}
		//Bishops
		if Board[y][x] == " "+B[5]+" " || Board[y][x] == " "+N[5]+" " {
			fmt.Println(Board[y][x])
			v2 := Pieces.Bishops_{y, x, Board, colorContrario}
			a = v2.MovesBishop()

		}

		//Queen
		if Board[y][x] == " "+B[2]+" " || Board[y][x] == " "+N[2]+" " {

			fmt.Println(Board[y][x])
			v3 := Pieces.Queen_{y, x, Board, colorContrario}

			a = v3.MovesQueen()
		}

		//Rook
		if Board[y][x] == " "+B[1]+" " || Board[y][x] == " "+N[1]+" " {

			fmt.Println(Board[y][x])
			v4 := Pieces.Rook_{y, x, Board, colorContrario}
			a = v4.MovesRook()
		}
		// Knight
		if Board[y][x] == " "+B[3]+" " || Board[y][x] == " "+N[3]+" " {

			fmt.Println(Board[y][x])
			v5 := Pieces.Knight_{y, x, Board, colorContrario}
			a = v5.MovesKnight()
		}

        llave_Rey := false
		// King
		if Board[y][x] == " "+B[4]+" " || Board[y][x] == " "+N[4]+" " {

			fmt.Println(Board[y][x])
			v6 := Pieces.King_{y, x, Board, colorContrario}
			a = v6.MovesKing()


            // This is for the rosk

            // para ver si se puede hacer hacia la isquierda
            if Board [y][ x -1 ] == " - " && Board [y][x - 2]== " - " && y == 0 || y == 7 {
                a = append(a,[2] int{ y , x - 2 } );
                a = append(a,[2] int{ y , x + 2 } );
                llave_Rey = true;            
            }
		}

		calcularTodosLosMovimietos( )

		//fmt.Println(movimientosPosiblesBlancos)
		//fmt.Println(movimientosPosiblesNegros)

		for i := 0; i < len(movimientosPosiblesNegros); i++ {
			Board__[movimientosPosiblesNegros[i][0]][movimientosPosiblesNegros[i][1]] = " o "
		}
		for i := 0; i < len(movimientosPosiblesBlancos); i++ {
			Board__[movimientosPosiblesBlancos[i][0]][movimientosPosiblesBlancos[i][1]] = " p "
		}
		


		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {

				fmt.Print(Board__[i][j])
			}
			fmt.Println()
		}

		if len(a) == 0 {
			fmt.Println("La ficha està bloqueada")
			y, x = logicChoises(" - ", "Escoge una posicion -> ")

		}

		if len(a) != 0 {
			Y, X := logicChoises("", " Donde deseas moverte")
			for i := 0; i < len(a); i++ {
				if a[i][0] == Y && a[i][1] == X {
					veri = true
					break
				}
			}


			if veri == true {
                // para controlar el enrrosque , esto va a verificar si la ficha es un rey y el moviminto

				// fmt.Println(( X == 3 || X == 7 ))
				// fmt.Println( X )
                
                if (Board[y][x] == " " + B[4] + " "|| Board[y][x] == " " + N[4] + " " ) && ( llave_Rey == true ) && ( X == 3 || X == 7 ) && verificarNumeroPosicion(y , x , movimientosPosiblesBlancos) &&   verificarNumeroPosicion(y , x , movimientosPosiblesNegros){
                    Board[Y][X] = Board[y][x] 

                    if X == 3{
                        Board[y][ x - 1] = Board[y][x - 4]
                        Board[y][x] = " - "
                        Board[y][x -4] = " - "
                    
                    }	else if X == 7{
						
                            Board[y][x +1] = Board[y][x+3]
                        	Board[y][x] = " - "
                        	Board[y ][x+3] = " - "
                    }

                } else{
					
				    Board[Y][X] = Board[y][x]
				    Board[y][x] = " - "
                }
				ix2 = true
			}

			if ix2 == true {

				break
			}
			ix++
		}

	}

}





// para validar si los numeros estan en la lista
func verificarNumeroPosicion( num1_ int , num2_ int , lisNum1_ [][2] int )bool {
    validation := false

    for i := 0; i < len (lisNum1_) ; i ++{
        
        if num1_ == lisNum1_[i][1] && num2_ == lisNum1_ [i][0]{
            validation = true;
        }
    }
    return validation;
    
}








