// this will be the most bigest proyect
// i dot know how to create the game but i going to build the ches game with ia

// i start this proyect in the date  2/10/2022

// I wish you the best caleb of the future you are capable do not give up. you are going to achieve it you are able

// He terminado la primera parte el programa puede mover las fichas correctamente horas de desarrollo 10

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"main/pieces_"
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

	/// this is for thet let moves
	to_move      int
	allowedMoves = map[int]position_let{}
	// piecesToatac [6]string
)

type position_let struct {
	AAA, BBB int
}

// this is for check the pieces
// Function to calculate the color of the pice to eat
func check(A [6]string, B string) bool {

	xddd := false
	for i := 0; i < len(A); i++ {
		if " "+A[i]+" " == B {
			xddd = true
		}

	}
	return xddd
}

func main() {

	fmt.Print("What is Your Color ->  W or B -> ")
	var second string
	fmt.Scanln(&second)

	it := 0
	for {

	
	
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

	

		if second == "b" {
			for i := 8; i > -1; i-- {
				for j := 8; j > -1; j-- {

					fmt.Print(Board[i][j])
					it++
				}
				fmt.Println()
			}

		} else if second == "w" {

			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {

					fmt.Print(Board[i][j])
				}
				fmt.Println()
			}
		}

		if it%2 == 0 {
			logic_moves(true , B , N  ) 
			second = "b"
		}

		if it%2 == 1 {
			logic_moves( true, N, B )
			second = "w"

		}

		it++
	}

}




func logic_moves(keyColor bool , ColorPiece , ColorPieceOpposite [6] string) {
	var position1 string
	var position2 string

	for {
		fmt.Print("Please enter the position > ")
		fmt.Scanln(&position1)

		// this is for separate the string
		s := position1
		v := strings.SplitAfter(s, "")

		// this is for become String to a int
		val, err := strconv.Atoi(v[1])

		if err != nil {

		} else {

			if val > 0 && val <= 9 {

				a := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
				itero := false

				for io := 0; io < 8; io++ {
					if a[io] == v[0] {
						itero = true
					}

				}
				if itero {

					// this is for x
					Change_letters_numbers := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
					// this is for y
					Change_numbers_numbers := map[string]int{"1": 7, "2": 6, "3": 5, "4": 4, "5": 3, "6": 2, "7": 1, "8": 0}

					TheXposition = Change_letters_numbers[v[0]]
					TheYposition = Change_numbers_numbers[v[1]]

					// this is for the pawn
					////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
					if Board[TheYposition][TheXposition] == " "+B[0]+" " || Board[TheYposition][TheXposition] == " "+ N[0]+" " {
						// ---------------------------------------------------
						for k := range allowedMoves {
							delete(allowedMoves, k)
						}
						// ---------------------------------------------------

						fmt.Println(Board[TheYposition][TheXposition])

						p := pieces_.Pawn__{TheYposition, TheXposition , ColorPieceOpposite, ColorPiece , Board, false}
						p.Pawn_Allowed_Play()

						for i := 0; i < len(pieces_.Pawn_Allowed_Moves__); i++ {
							allowedMoves[len(allowedMoves)+1] = position_let{pieces_.Pawn_Allowed_Moves__[i+1].This_A, pieces_.Pawn_Allowed_Moves__[i+1].This_B}
						}
						fmt.Println(allowedMoves)
						if len(allowedMoves) == 0 {
							fmt.Println("La ficha está bloqueda")
						} else {
							break
						}

					}

					//this is for the Rook
					if Board[TheYposition][TheXposition] == " "+B[1]+" " || Board[TheYposition][TheXposition] == " "+N[1]+" " {

						a := pieces_.Rook_{TheXposition, TheYposition , ColorPieceOpposite, ColorPiece , Board}
						a.Roock_allovedMoves_Play()

						for i := 0; i < len(pieces_.Roock_allovedMoves_); i++ {
							allowedMoves[len(allowedMoves)+1] = position_let{pieces_.Roock_allovedMoves_[i+1].This_A, pieces_.Roock_allovedMoves_[i+1].This_B}
						}

						fmt.Println(allowedMoves)
						if len(allowedMoves) == 0 {
							fmt.Println("La ficha está bloqueda")
						} else {
							break
						}

					}

					// this is for the Queen
					if Board[TheYposition][TheXposition] == " "+B[2]+" " || Board[TheYposition][TheXposition] == " "+N[2]+" " {

						b := pieces_.Queen_{TheYposition, TheXposition,ColorPieceOpposite, ColorPiece , Board}
						b.Queen_allowedMoves_Queen_play()

						for i := 0; i < len(pieces_.Queen_allowedMoves_Queen___); i++ {
							allowedMoves[len(allowedMoves)+1] = position_let{pieces_.Queen_allowedMoves_Queen___[i+1].This_A, pieces_.Queen_allowedMoves_Queen___[i+1].This_B}
						}

						fmt.Println(allowedMoves)
						if len(allowedMoves) == 0 {
							fmt.Println("La ficha está bloqueda")
						} else {
							break
						}

					}

					// this is for the Bishop

					if Board[TheYposition][TheXposition] == " "+B[5]+" " || Board[TheYposition][TheXposition] == " "+N[5]+" " {

						bbbbb := pieces_.Bishop_{TheYposition, TheXposition, ColorPieceOpposite, ColorPiece , Board}
						bbbbb.Bishop_AllowedMoves_bishop_ToPLay()

						for i := 0; i < len(pieces_.Bishop_allowedMoves_); i++ {
							allowedMoves[len(allowedMoves)+1] = position_let{pieces_.Bishop_allowedMoves_[i+1].This_A, pieces_.Bishop_allowedMoves_[i+1].This_B}
						}

						fmt.Println(allowedMoves)
						if len(allowedMoves) == 0 {
							fmt.Println("La ficha está bloqueda")
						} else {
							break
						}

					}
					//this is for the Knight
					if Board[TheYposition][TheXposition] == " "+B[3]+" " || Board[TheYposition][TheXposition] == " "+N[3]+" " {

						XU := pieces_.Knight_{TheYposition, TheXposition,ColorPieceOpposite, ColorPiece , Board}
						XU.Knight_Allowed_Play()

						for i := 0; i < len(pieces_.Knight_allowedMoves__); i++ {
							allowedMoves[len(allowedMoves)+1] = position_let{pieces_.Knight_allowedMoves__[i+1].This_A, pieces_.Knight_allowedMoves__[i+1].This_B}

						}

						fmt.Println(allowedMoves)
						if len(allowedMoves) == 0 {
							fmt.Println("La ficha está bloqueda")
						} else {
							break
						}

					}

					//this is for the King

				}

			}	


			
		}

	}

	fmt.Println("The position y", TheYposition)
	fmt.Println("The position x", TheXposition)

	// here is the next step

	for {
		fmt.Print("Where do you want to play > ")
		fmt.Scanln(&position2)

		// this is for separate the string
		ss := position2
		vv := strings.SplitAfter(ss, "")

		// this is for become String to a int
		vall, errr := strconv.Atoi(vv[1])
		if errr != nil {

		} else {

			if vall > 0 && vall < 9 {

				a := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
				itero := false

				for io := 0; io < 8; io++ {
					if a[io] == vv[0] {
						itero = true
					}

				}
				if itero {
					// this is for x
					Change_letters_numbers2 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
					// this is for y
					Change_numbers_letters2 := map[string]int{"1": 7, "2": 6, "3": 5, "4": 4, "5": 3, "6": 2, "7": 1, "8": 0}

					TheXpositionToMove = Change_letters_numbers2[vv[0]]
					TheYpositionToMove = Change_numbers_letters2[vv[1]]

					the_key := false

					for i := 0; i < len(allowedMoves); i++ {

						if allowedMoves[i+1].AAA == TheYpositionToMove && allowedMoves[i+1].BBB == TheXpositionToMove {
							the_key = true

						}
					}

					if the_key {
						break
					}

					// break

				}

			}
		}

	}

	fmt.Println("The position y", TheYpositionToMove)
	fmt.Println("The position x", TheXpositionToMove)

	// changer := Board[TheYpositionToMove][TheXpositionToMove]
	Board[TheYpositionToMove][TheXpositionToMove] = Board[TheYposition][TheXposition]
	Board[TheYposition][TheXposition] = " - "

	for k := range allowedMoves {
		delete(allowedMoves, k)
	}
}
