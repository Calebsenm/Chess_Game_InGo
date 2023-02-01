// this will be the most bigest proyect
// i dot know how to create the game but i going to build the ches game with ia

// i start this proyect in the date  2/10/2022
// i delete the proyect in nov 21 / 2022
// i start the proyect in  nov 21 / 2022
// refactorizacion del codigo 22 / enero /2023 

package main

import (
	"fmt"
	"os"
	"os/exec"
	"chess/Functions"
)

var (
	N = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	TheYposition int
	TheXposition int

	TheXpositionToMove int
	TheYpositionToMove int


	movimientosPosiblesBlancos [][2]int
	movimientosPosiblesNegros  [][2]int

	Board = [9][9]string{

		{" 8 ", " " + N[1] + " ", " " + N[3] + " ", " " + N[5] + " ", " " + N[2] + " ", " " + N[4] + " ", " " + N[5] + " ", " " + N[3] + " ", " " + N[1] + " "},
		{" 7 ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " "},
		{" 6 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 4 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 2 ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " " ," " + B[0] + " "},
		{" 1 ", " " + B[1] + " ", " " + B[3] + " ", " " + B[5] + " ", " " + B[2] + " ", " " + B[4] + " ", " " + B[5] + " ", " " + B[3] + " ", " " + B[1] + " "},
		{"   ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H "},
	}

	// Board = [9][9]string{

	// 	{" 8 ", " " + N[1] + " ", " - ", " - ", " - ", " " + N[4] + " ", " - ", " - ", " " + N[1] + " "},
	// 	{" 7 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 6 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 4 ", " - ", " - ", " - ", " " + N[2] + " ", " - ", " - ", " - ", " - "},
	// 	{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
	// 	{" 2 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " " + N[2] + " "},
	// 	{" 1 ", " " + B[1] + " ", " - ", " - ", " - ", " " + B[4] + " ", " - ", " - ", " " + B[1] + " "},
	// 	{"   ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H "},
	// }

)

// this funtion going to start the game

func main() {
	fmt.Println("Wecolcome To chess game")
	var choice int

	for {
		fmt.Print("Please input " + "\n 1 for white  pieces \n 2 for black pieces \n -> ")

		fmt.Scan(&choice)

		if choice == 1 || choice == 2 {
			break
		} else {
			fmt.Println("Intentalo de nuevo ")
		}
	}

	iterator := 1
	colorContrario := [6]string{}

	for {
		iterator++

		// this clear the window
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println("-------------------------------------------------------")
		Functions.PrintBoard(choice, Board)
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

			Functions.LlamarMovimientoFichas("1", colorContrario , Board);

		} else {
			if choice == 2 {
				choice = 1
			} else {
				choice = 2
			}
			Functions.LlamarMovimientoFichas("2", colorContrario , Board );
		}

	}

}
