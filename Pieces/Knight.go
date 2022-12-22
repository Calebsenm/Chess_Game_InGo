package Pieces


var (
	n____ = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b____ = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	x____ int
	y____ int

	listOfNumbersKnight = [][2]int{}
	board____           [9][9]string
)

type Knight_ struct {
	Y____     int
	X____     int
	Board____ [9][9]string
	Color____ [6]string
}

func (Kn *Knight_) MovesKnight() [][2]int {

	var vacia [][2]int
	listOfNumbersKnight = vacia

	// this save variable
	x____ = Kn.X____
	y____ = Kn.Y____
	board____ = Kn.Board____


	KnightAlgoritmo( 2  , -1 )
	KnightAlgoritmo( 2  ,  1)
	KnightAlgoritmo( -1 , -2)
	KnightAlgoritmo( -1 , +2)

	KnightAlgoritmo( -2 , 1 )
	KnightAlgoritmo( -2 , -1)
	KnightAlgoritmo(  1 ,  2)
	KnightAlgoritmo(  1 , -2)
		
	return listOfNumbersKnight
}

func KnightAlgoritmo(yChance, xChance int) {

	atactPiece____ := board____[y____][x____]
	//fmt.Println(atactPiece____)
	y1____ := y____
	x1____ := x____

	y1____ = y1____ + yChance
	x1____ = x1____ + xChance 


	if fueraRango____(y1____, x1____) {
		if board____[y1____][x1____] == " - " {
			listOfNumbersKnight = append(listOfNumbersKnight, [2]int{y1____, x1____})
		}

		//fmt.Println(fichaAtaca____(atactPiece____))
		if fichaAtaca____(atactPiece____) == 1 {
			if fichas____(board____[y1____][x1____], n____){
		
				listOfNumbersKnight = append(listOfNumbersKnight, [2]int{y1____, x1____})
			}
	

		} 
		if fichaAtaca____(atactPiece____) == 2 {
			if fichas____(board____[y1____][x1____], b____){
				listOfNumbersKnight = append(listOfNumbersKnight, [2]int{y1____, x1____})
			
			}
	
		}

	}
}

// out of the range
func fueraRango____(y, x int) bool {
	keyValue := false

	if y >= 0 && y <= 7 && x >= 1 && x <= 8 {
		keyValue = true
	}
	return keyValue
}

// what pice is fithg
func fichaAtaca____(atacaX____ string) int {

	var num int

	for i := 0; i < 6; i++ {
		//fmt.Println(atacaX____ ,"  - ", " " + b____[i] + " ")
		if atacaX____ == " " + b____[i] + " " {
			num = 1
			break

		} else if atacaX____ == " " + n____[i] + " " {
			num = 2
			break

		}
	}

	return num
}

// to control the opponent's chip attack
func fichas____(ficha____ string, list____ [6]string) bool {

	for i := 0; i < len(list____); i++ {
		//fmt.Println( " "+list____[i]+" ", " - ", ficha____)
		if " "+list____[i]+" " == ficha____{

			return true
		}
	}
	return false
}
