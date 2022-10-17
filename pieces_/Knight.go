package pieces_

import "fmt"

var (
	Knight_allowedMoves__ = map[int]Contructor{}
)

type Knight_ struct {
	Y_K              int
	X_K              int
	B_Pie            [6]string
	W_Pie            [6]string
	The_board_Knight [9][9]string
}

func (K *Knight_) Knight_Allowed_Play() {

	if checker(K.W_Pie, K.The_board_Knight[K.Y_K][K.X_K]) {

		fmt.Println("No sea pendejo")
		// this is for up A
		checher_Knight(K.Y_K, K.X_K, -2, +1, K.The_board_Knight, K.B_Pie)
		checher_Knight(K.Y_K, K.X_K, -1, +2, K.The_board_Knight, K.B_Pie)

		// this is for up B
		checher_Knight(K.Y_K, K.X_K, -2, -1, K.The_board_Knight, K.B_Pie)
		checher_Knight(K.Y_K, K.X_K, -1, -2, K.The_board_Knight, K.B_Pie)

		// this is for left Down  A
		checher_Knight(K.Y_K, K.X_K, +2, +1, K.The_board_Knight, K.B_Pie)
		checher_Knight(K.Y_K, K.X_K, +1, +2, K.The_board_Knight, K.B_Pie)

		// this is for left Down  B
		checher_Knight(K.Y_K, K.X_K, +2, -1, K.The_board_Knight, K.B_Pie)
		checher_Knight(K.Y_K, K.X_K, +1, -2, K.The_board_Knight, K.B_Pie)

	} else {

		// this is for up A
		checher_Knight(K.Y_K, K.X_K, -2, +1, K.The_board_Knight, K.B_Pie)
		checher_Knight(K.Y_K, K.X_K, -1, +2, K.The_board_Knight, K.B_Pie)

		// this is for up B
		checher_Knight(K.Y_K, K.X_K, -2, -1, K.The_board_Knight, K.B_Pie)
		checher_Knight(K.Y_K, K.X_K, -1, -2, K.The_board_Knight, K.B_Pie)

		// this is for left Down  A
		checher_Knight(K.Y_K, K.X_K, +2, +1, K.The_board_Knight, K.B_Pie)
		checher_Knight(K.Y_K, K.X_K, +1, +2, K.The_board_Knight, K.B_Pie)

		// this is for left Down  B
		checher_Knight(K.Y_K, K.X_K, +2, -1, K.The_board_Knight, K.B_Pie)
		checher_Knight(K.Y_K, K.X_K, +1, -2, K.The_board_Knight, K.B_Pie)
	}
}

func checher_Knight(Y_ int, X_ int, M_A int, M_B int, BD_ [9][9]string, Color_ [6]string) {

	it_1 := M_A
	it_2 := M_B

	if Y_ + it_1 >= 1 && Y_ + it_1 <= 7 && X_ + it_2 >= 1 && X_ + it_2 <=  8 {
		if BD_[Y_+it_1][X_+it_2] == " - " || checker(Color_, BD_[Y_+it_1][X_+it_2]) {
		Knight_allowedMoves__[len(Knight_allowedMoves__)+1] = Contructor{Y_ + it_1, X_ + it_2}
	}

	}

	
}
