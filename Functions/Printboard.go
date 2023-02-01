
package Functions


import (
	"fmt"	
)


// funtion for printBoard
func PrintBoard(lastChoise int, Board [9][9]string) {

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