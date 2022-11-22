// this will be the most bigest proyect
// i dot know how to create the game but i going to build the ches game with ia

// i start this proyect in the date  2/10/2022
// i delete the proyect in nov 21 / 2022
// i start the proyect in  nov 21 / 2022

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	N = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	Board = [9][9]string{

		{" 8 ", " " + N[1] + " ", " " + N[3] + " ", " " + N[5] + " ", " " + N[4] + " ", " " + N[2] + " ", " " + N[5] + " ", " " + N[3] + " ", " " + N[1] + " "},
		{" 7 ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " "},
		{" 6 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 4 ", " - ", " - ", " - ", " - ", " " + " " + " ", " - ", " - ", " - "},
		{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 2 ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " "},
		{" 1 ", " " + B[1] + " ", " " + B[3] + " ", " " + B[5] + " ", " " + B[4] + " ", " " + B[2] + " ", " " + B[5] + " ", " " + B[3] + " ", " " + B[1] + " "},
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

//funtion for printBoard
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
func logicChoises(player_name string) (int, int) {

	values := [2]int{}
	for {
		var position1 string

		fmt.Print("Jugador " + player_name + " Please enter the position > ")
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

		if choice == 1{
			colorContrario = N
		}	else{
			colorContrario = B
		}

		if iterator % 2 == 0 {
			fmt.Println(colorContrario)
			Y, X := logicChoises("1")
			fmt.Println(Y, X)
		

		
			colorActual1 := [6]string{}

			for {

				if Board[Y][X] == B[0] || Board[Y][X] == B[1] || Board[Y][X] == B[2] || Board[Y][X] == B[3] || Board[Y][X] == B[4] || Board[Y][X] == B[5]  {
					colorActual1 = B
				}  else {
					colorActual1 = N
			
				}

				fmt.Println(colorContrario)
				fmt.Println(colorActual1)


				if colorContrario != colorActual1{
					fmt.Println("Error color incorrecto")
					Y, X = logicChoises("1")
				}	else{
						break
				}

				

			}

			if choice == 1 {
				choice = 2
			} else {
				choice = 1
			}
			
			llamarMovimientosLogicos("Loro con sal ")
		} else {

			fmt.Println(colorContrario)
			Y, X := logicChoises("2")
			fmt.Println(Y, X)
		


			colorActual2 := [6]string{}
			for {

				if Board[Y][X] == N[0] || Board[Y][X] == N[1] || Board[Y][X] == N[2] || Board[Y][X] == N[3] || Board[Y][X] == N[4] || Board[Y][X] == N[5]  {
					colorActual2 = N
				
				}  else {
					colorActual2 = B
			
				}

				fmt.Println(colorContrario)
				fmt.Println(colorActual2)

				if colorContrario != colorActual2{
					fmt.Println("Error color incorrecto")
					Y, X = logicChoises("2")
				}	else{
						break
				}
			}

			if choice == 2 {
				choice = 1
			} else {
				choice = 2
			}

			llamarMovimientosLogicos("Loro con sal ")
		}

	}

}

func llamarMovimientosLogicos(loro string) {

	fmt.Println("Movimiento permitido " + loro)
}
