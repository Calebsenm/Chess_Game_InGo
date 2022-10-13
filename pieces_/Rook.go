package pieces_

import (
	"fmt"
)

var (
	Rook_allowedMoves_Rook_play = map[int]mapp{}
)

type mapp struct {
	ASSS int
	BSSS int
}

type Rook_ struct {
	X_Rook_      int
	Y_Rook_      int
	Black_pieces [6]string
	White_pieces [6]string

	The_Board_Rook [9][9]string
}

func (v *Rook_) Rook_allowedMovesRook() {

	fmt.Println(v.The_Board_Rook[v.Y_Rook_][v.X_Rook_])

	//this is for up moves

	iterator := 1
	for {

		if v.Y_Rook_-iterator >= 0 {
			if v.The_Board_Rook[v.Y_Rook_-iterator][v.X_Rook_] == " - "  || check(v.Black_pieces, v.The_Board_Rook[v.Y_Rook_-iterator][v.X_Rook_]) == true{
				
				Rook_allowedMoves_Rook_play[len(Rook_allowedMoves_Rook_play)+1] = mapp{v.Y_Rook_-iterator , v.X_Rook_ }
				
				if check(v.Black_pieces, v.The_Board_Rook[v.Y_Rook_-iterator][v.X_Rook_]) == true {
					break
				}

			}  else {
				break
			}

		}else {
			break
		}

		iterator++
	}

	// the right
	iterator1 := 1
	for {
		if v.X_Rook_+iterator1 <= 8 {
			if v.The_Board_Rook[v.Y_Rook_][v.X_Rook_+iterator1] == " - " || check(v.Black_pieces, v.The_Board_Rook[v.Y_Rook_][v.X_Rook_+iterator1]) == true {

				Rook_allowedMoves_Rook_play[len(Rook_allowedMoves_Rook_play)+1] = mapp{v.Y_Rook_, v.X_Rook_ + iterator1}
				if check(v.Black_pieces, v.The_Board_Rook[v.Y_Rook_][v.X_Rook_+iterator1]) == true {
					break
				}

			} else {
				break
			}

		}else {
			break
		}
		iterator1++
	}

	// the lefht
	iterator2 := 1
	for {
		if v.Y_Rook_-iterator2 >= 0 {
			if v.The_Board_Rook[v.Y_Rook_][v.X_Rook_-iterator2] == " - " || check(v.Black_pieces, v.The_Board_Rook[v.Y_Rook_][v.X_Rook_-iterator2]) == true {

				Rook_allowedMoves_Rook_play[len(Rook_allowedMoves_Rook_play)+1] = mapp{v.Y_Rook_, v.X_Rook_ - iterator2}
				if check(v.Black_pieces, v.The_Board_Rook[v.Y_Rook_][v.X_Rook_-iterator2]) == true {
					break
				}

			} else {
				break
			}

		}else {
			break
		}

		iterator2++
	}

	// down
	iterator3 := 1
	for {
		if v.Y_Rook_+iterator3 <= 8 {
			if v.The_Board_Rook[v.Y_Rook_+iterator3][v.X_Rook_] == " - " || check(v.Black_pieces, v.The_Board_Rook[v.Y_Rook_+iterator3][v.X_Rook_]) == true {

				Rook_allowedMoves_Rook_play[len(Rook_allowedMoves_Rook_play)+1] = mapp{v.Y_Rook_ + iterator3, v.X_Rook_}
				if check(v.Black_pieces, v.The_Board_Rook[v.Y_Rook_ + iterator3][v.X_Rook_]) == true {
					break
				}
			} else {
				break
			}

		}else {
			break
		}

		iterator3++
	}

	

}

// Function to calculate the color of the pice to eat
func check(A [6]string, B string) bool {

	xddd := false
	for i := 0; i < len(A); i++ {
		if  " " + A[i] + " " ==   B  {
			xddd = true
			
			
		}

	}
	return xddd
}

func (v *Rook_) Running() {

	fmt.Println("The Rook is running in ", v.X_Rook_, v.Y_Rook_)

}

func Hello() {

	fmt.Println("Welcome to my game")

}
