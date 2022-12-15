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

	// 	{" 8 ", " " + N[1] + " ", " " + N[3] + " ", " " + N[5] + " ", " " + N[4] + " ", " " + N[2] + " ", " " + N[5] + " ", " " + N[3] + " ", " " + N[1] + " "},
	// 	{" 7 ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " "},
	// 	{" 6 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 4 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 2 ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " _ ", " " + B[0] + " "},
	// 	{" 1 ", " " + B[1] + " ", " " + B[3] + " ", " " + B[5] + " ", " " + B[4] + " ", " " + B[2] + " ", " " + B[5] + " ", " " + B[3] + " ", " " + B[1] + " "},
	// 	{"   ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H "},
	// }

	Board = [9][9]string{

		{" 8 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 7 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 6 ", " - ", " - ", " - ", " " + N[2] + " ", " - ", " - ", " - ", " - "},
		{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 4 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 2 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 1 ", " - ", " - ", " - ",  " " + B[2] + " ", " - ", " - ", " " + B[2] + " ", " - "},
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

		//Pawm 
		if Board[y][x] == " "+B[0]+" " || Board[y][x] == " "+N[0]+" " {
			fmt.Println(Board[y][x])

			v1 := Pieces.Pawm_{y , x , Board, colorContrario}
			a = v1.MovesCalculate()

		}
		//Bishops
		if Board[y][x] == " "+B[5]+" " || Board[y][x] == " "+N[5]+" " {
			fmt.Println(Board[y][x])
			v2 := Pieces.Bishops_{y , x , Board , colorContrario} 
			a = v2.MovesCalculate()
			
		}

		//Queen
		if Board[y][x] == " "+B[2]+" " || Board[y][x] == " "+N[2]+" " {

			fmt.Println(Board[y][x])
			v3 := Pieces.Queen_{y , x , Board , colorContrario}

			a = v3.MovesQueen()
		}

		for i := 0; i < len(a); i++ {	
			Board__[a[i][0]][a[i][1] ] = " o "
		}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {

				fmt.Print(Board__[i][j])
			}
			fmt.Println()
		}


		if len(a) == 0 {
			fmt.Println("La ficha està bloqueada"); 
			y, x = logicChoises(" - ", "Escoge una posicion -> ")

		} 

		if len(a) != 0{
			Y, X := logicChoises("", " Donde deseas moverte")
			for i := 0; i < len(a); i++ {

				if a[i][0] == Y && a[i][1] == X {
					veri = true
					break
				}
			}


            

			if veri == true {
		
				Board[Y][x] = Board[y][x]
				Board[y][x] = " - "
	
				ix2 = true
			}
		
	
			if ix2 == true{
				break
			}
			ix++
		}
		
	}

	
}
