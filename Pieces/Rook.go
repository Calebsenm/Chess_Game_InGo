package Pieces;



var(

	n___ = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b___ = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	x___ int 
	y___ int 

	listOfNumbersRook = [][2]int { }
	board___  [9][9]	string
)

type Rook_ struct {

	Y___   				int 
	X___   				int
	Board___ [9][9] 	string
	Color___ [6]		string 
}

func (Rk * Rook_) MovesRook() [][2] int {
	var vacia  [][2]int
 	listOfNumbersRook = vacia

	// this save the variable 
	x___ = Rk.X___
	y___ = Rk.Y___
    board___ = Rk.Board___


	// this is for the Up moves
	RookAlgoritmo( -1 , 0 )
	// this is for the Left moves
	RookAlgoritmo( 0 , -1 )
	// this is for the right moves
	RookAlgoritmo( 0 , +1 )
	// this is for the Dowm moves
	RookAlgoritmo( +1 , 0 )

	return listOfNumbersRook
} 



func RookAlgoritmo(yChange , xChange int  ){
 
	atactPiece := board___[y___][x___];

	y111 := y___
	x111 := x___


	i := 0;
	for i < 9{
		i++

		y111 = y111 + yChange;
		x111 = x111 + xChange;

		if fueraRango___(y111,x111){

			if board___[y111][x111] == " - "{
				listOfNumbersRook = append(listOfNumbersRook, [2]int{y111 ,x111} )
			}	

			if fichaAtaca___( atactPiece ) == 1{
				if fichas___(board___[y111][x111], n___) {
					listOfNumbersRook = append(listOfNumbersRook, [2]int{y111, x111})
					break
				} else if fichas___(board___[y111][x111], b___) {
					break

				}

			}	else if fichaAtaca___(atactPiece) == 2 {
				if fichas___(board___[y111][x111], b___) {
					listOfNumbersRook = append(listOfNumbersRook, [2]int{y111, x111})
					break

				} else if fichas___(board___[y111][x111], n___) {
					break
				}
			}


		}
	}
	
}

// out of the range
func fueraRango___(y , x int ) bool {
	keyValue := false;

	if y >= 0 && y <= 7 && x >= 1 && x <= 8 {
		keyValue = true;
	}
	return keyValue;
}

// what pice is fithg
func fichaAtaca___(atacaX___ string) int{

	var num int;

	for i := 0; i < 6; i++ {
		if atacaX___ == " " + b___[i] + " "{
			num = 1;
			break;

		} else if atacaX___ == " " + n___[i] + " " {
			num = 2;
			break;

		}
	}

	return num;
}

// to control the opponent's chip attack
func fichas___(ficha___ string , list___ [6] string) bool {
	
	for i := 0; i < len(list___); i++ {
	
		if " " +list___[i] + " " == ficha___ {
			return true
		}
	}
	return false
}
