package Pieces


var (
	n_____ = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b_____ = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	x_____ int
	y_____ int

	listOfNumbersKing = [][2]int{}
	board_____           [9][9]string
)

type King_ struct {
	Y_____     int
	X_____     int
	Board_____ [9][9]string
	Color_____ [6]string
}

func (Kin * King_) MovesKing() [][2]int {

	var vacia [][2]int
	listOfNumbersKing = vacia

	// this save variable
	x_____ = Kin.X_____
	y_____ = Kin.Y_____
	board_____ = Kin.Board_____


	KingAlgoritmo( -1 , 0 )
	KingAlgoritmo( -1 , -1)
	KingAlgoritmo( -1 , +1)
	KingAlgoritmo(  0 , -1)
	KingAlgoritmo(  0 , +1)
	KingAlgoritmo(  1 , -1)
	KingAlgoritmo(  1 ,  1)
	KingAlgoritmo(  1 ,  0)
		
	return listOfNumbersKing
}

func KingAlgoritmo(yChance_, xChance_ int) {

	atactPiece_____ := board_____[y_____][x_____]
	//fmt.Println(atactPiece____)
	y1_____ := y_____
	x1_____ := x_____

	y1_____ = y1_____ + yChance_
	x1_____ = x1_____ + xChance_ 


	if fueraRango_____(y1_____, x1_____) {
		if board_____[y1_____][x1_____] == " - " {
			listOfNumbersKing = append(listOfNumbersKing, [2]int{y1_____, x1_____})
		}

		//fmt.Println(fichaAtaca____(atactPiece____))

		// white 
		if fichaAtaca_____(atactPiece_____) == 1 {
			if fichas_____(board_____[y1_____][x1_____], n_____){
		
				listOfNumbersKing = append(listOfNumbersKing, [2]int{y1_____, x1_____})
			}
	

		} 
		// black 
		if fichaAtaca_____(atactPiece_____) == 2 {
			if fichas_____(board____[y1_____][x1_____], b_____){
				listOfNumbersKing = append(listOfNumbersKing, [2]int{y1_____, x1_____})
			
			}
	
		}

	}
}

// out of the range
func fueraRango_____(y, x int) bool {
	keyValue := false

	if y >= 0 && y <= 7 && x >= 1 && x <= 8 {
		keyValue = true
	}
	return keyValue
}

// what pice is fithg
func fichaAtaca_____(atacaX_____ string) int {

	var num int

	for i := 0; i < 6; i++ {
		//fmt.Println(atacaX____ ,"  - ", " " + b____[i] + " ")
		if atacaX_____ == " " + b_____[i] + " " {
			num = 1
			break

		} else if atacaX_____ == " " + n_____[i] + " " {
			num = 2
			break

		}
	}

	return num
}

// to control the opponent's chip attack
func fichas_____(ficha_____ string, list_____ [6]string) bool {

	for i := 0; i < len(list_____); i++ {
		//fmt.Println( " "+list____[i]+" ", " - ", ficha____)
		if " "+list_____[i]+" " == ficha_____{

			return true
		}
	}
	return false
}
