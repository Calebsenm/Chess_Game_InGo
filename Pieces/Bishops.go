package Pieces

var (
	n_ = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b_ = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	y_ int
	x_ int

	listNumbers_ = [][2]int{}
	board_       [9][9]string
)

type Bishops_ struct {
	Y_     int
	X_     int
	Board_ [9][9]string
	Color  [6]string
}

func (bh Bishops_) MovesBishop() [][2]int {

	// this save the variable
	x_ = bh.X_
	y_ = bh.Y_
	board_ = bh.Board_

	// this for up right
	BishopsAlgorimo(-1, 1)
	//Up left
	BishopsAlgorimo(-1, -1)
	//Dowm lef
	BishopsAlgorimo(1, -1)
	// Dowm right
	BishopsAlgorimo(1, 1)

	return listNumbers_
}

// this is a function for the Bishop that will calculate The moves

func BishopsAlgorimo(yCambio__, xCambio__ int) {

	ataquesFicha := board_[y_][x_]

	y11 := y_
	x11 := x_

	i_ := 0

	for i_ < 9 {
		i_++

		y11 = y11 + yCambio__
		x11 = x11 + xCambio__

		if fueraRango_1(y11, x11) {

			if board_[y11][x11] == " - " {
				listNumbers_ = append(listNumbers_, [2]int{y11, x11})
			}

			if fichaAtaca_(ataquesFicha) == 1 {

				if ficha_(board_[y11][x11], n_) {
					listNumbers_ = append(listNumbers_, [2]int{y11, x11})
					break
				} else if ficha_(board_[y11][x11], b_) {
					break

				}

			} else if fichaAtaca_(ataquesFicha) == 2 {

				if ficha_(board_[y11][x11], b_) {
					listNumbers_ = append(listNumbers_, [2]int{y11, x11})
					break

				} else if ficha_(board_[y11][x11], n_) {
					break
				}
			}

		}
	}

}

// para controlar el fuera de rango
func fueraRango_1(y1, x1 int) bool {

	dde := false
	if y1 >= 0 && y1 <= 7 && x1 >= 1 && x1 <= 8 {
		dde = true
	}
	return dde
}

// ver la ficha que ataca
func fichaAtaca_(atacaX string) int {

	var numero int

	for i := 0; i < 6; i++ {

		if atacaX == " "+b__[i]+" " {
			numero = 1
			break
		} else if atacaX == " "+n__[i]+" " {
			numero = 2
			break
		}
	}
	return numero
}

// para controlar el ataque ficha contraria
func ficha_(ficha string, list [6]string) bool {
	for i := 0; i < len(list); i++ {
		if " "+list[i]+" " == ficha {
			return true
		}
	}
	return false
}






/*
func (objet Bishops_) MovesCalculate() [][2]int{
	Color111 = objet.Color

	mv := movesBishop(objet.Y_ , objet.X_, objet.Board_)
	return mv


}
func movesBishop(y_ int , x_ int , board [9][9]string ) [][2]int{


	//let_moves = append (let_moves,[]int{y_,x_})

	if colorChecker(board[y_ ][x_] , b_ ){

		moves_( -1 , -1 , y_ , x_ , board)
		moves_( -1 , +1 , y_ , x_ , board)
		moves_( +1 , +1 , y_ , x_ , board)
		moves_( +1 , -1 , y_ , x_ , board)

	}	else{

		moves_( -1 , -1 , y_ , x_ , board)
		moves_( -1 , +1 , y_ , x_ , board)
		moves_( +1 , +1 , y_ , x_ , board)
		moves_( +1 , -1 , y_ , x_ , board)
	}

	return let_moves1
}
func  moves_(a  int , b int , ay int , ax int , aboard [9][9]string){

	in1 := 0
	in2 := 0



	for {
		fmt.Println("Loro")

		verif1 := ay + in1
		verif2 := ax + in2

		choser1 := verif1 >= 0 && verif1 <= 7
		choser2 := verif2  >= 1 && verif1 <= 8

		if choser1 == false || choser2 == false{
			break
		}

		Col := [6] string {}
		if Color111 == n_ {
			Col = b_
		}	else{
			Col = n_
		}

		if aboard[ay + in1][ax +in2]  != " - " || checkerHacker(aboard[ay + in1][ax +in2] , Col) == true {
			fmt.Println(aboard[ay + in1][ax +in2]  )
			break

		}

		if choser1 == true && choser2 == true{
			if aboard[ay + in1][ax +in2]  == " - "{
				let_moves1 = append(let_moves1, [2]int{ay + in1 , ax +in2})
				fmt.Println(ay + in1 ," ", ax +in2)

			}
		}

		if a > 0{
			in1 ++
		}
		if a < 0{
			in1 --
		}
		if b > 0{
			in2 ++
		}
		if b < 0{
			in2 --
		}
	}
}

*/
