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
	
	
)

var (
	N = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	Board = [9][9]string{

		{" 8 ", " " + N[1] + " ", " " + N[3] + " ", " " + N[5] + " ", " " + N[4] + " ", " " + N[2] + " ", " " + N[5] + " ", " " + N[3] + " ", " " + N[1] + " "},
		{" 7 ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " ", " " + N[0] + " "},
		{" 6 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 5 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 4 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 3 ", " - ", " - ", " - ", " - ", " - ", " - ", " - ", " - "},
		{" 2 ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " ", " " + B[0] + " "},
		{" 1 ", " " + B[1] + " ", " " + B[3] + " ", " " + B[5] + " ", " " + B[4] + " ", " " + B[2] + " ", " " + B[5] + " ", " " + B[3] + " ", " " + B[1] + " "},
		{"   ", " A ", " B ", " C ", " D ", " E ", " F ", " G ", " H "},
	}

	TheYposition int
	TheXposition int

	TheXpositionToMove int
	TheYpositionToMove int
)


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
				}
				fmt.Println()
			}

		} else {

			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {

					fmt.Print(Board[i][j])
				}
				fmt.Println()
			}
		}

		if it%2 == 0 {
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
						if itero == true {

							// this is for x
							Change_letters_numbers := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
							// this is for y
							Change_numbers_numbers := map[string]int{"1": 7, "2": 6, "3": 5, "4": 4, "5": 3, "6": 2, "7": 1, "8": 0}

							
							TheXposition = Change_letters_numbers[v[0]]
							TheYposition = Change_numbers_numbers[v[1]]
							
							
							algoritmoXDDDD := The_pawn1{TheXposition,TheYposition,B,N,Board}
							algoritmoXDDDD.Allowed_moves()

							if len(ReturnValues()) == 0{
								fmt.Println("La ficha está bloqueda")
							}	else{
								break
							}

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
						if itero == true {
							// this is for x
							Change_letters_numbers2 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
							// this is for y
							Change_numbers_letters2 := map[string]int{"1": 7, "2": 6, "3": 5, "4": 4, "5": 3, "6": 2, "7": 1, "8": 0}

							TheXpositionToMove = Change_letters_numbers2[vv[0]]
							TheYpositionToMove = Change_numbers_letters2[vv[1]]

							
						
							algoritmoXDDDD := The_pawn1{TheXpositionToMove,TheYpositionToMove,B,N,Board}
							algoritmoXDDDD.Allowed_moves()

							Bread_cheess := ReturnValues()

							the_key := false
							
							for i := 0; i < len(Bread_cheess); i++ {
								if Bread_cheess[i].A == TheYpositionToMove && Bread_cheess[i].B == TheXpositionToMove {
									the_key = true
								}
							}

							if the_key == true{
								break
							}

							
						}

					}
				}

			}

			fmt.Println("The position y", TheYpositionToMove)
			fmt.Println("The position x", TheXpositionToMove)

			changer := Board[TheYpositionToMove][TheXpositionToMove]
			Board[TheYpositionToMove][TheXpositionToMove] = Board[TheYposition][TheXposition]
			Board[TheYposition][TheXposition] = changer
			
			

			

			

		} 
		
		if it%2 == 1  {

			// the algoritmo Hacker
			
			fmt.Print("El moviento de la maquina es ? ")
			var second string
			fmt.Scanln(&second)
		}

		it++
	}

}


