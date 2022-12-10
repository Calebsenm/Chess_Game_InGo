package Pieces



var(
	n__ = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b__ = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}
	board__ [9][9] int

	x__ int 
	y__ int

)

type Queen_ struct{
	Y_ 				int
	X_				int
	Board_ [9][9]	string
	Color [6]  		string
}

func (ob Queen_ ) MovesQueen()  [][]int {
	var listOfnumbers_ = [][]int{}
	x__ = ob.X_
	y__ = ob.Y_


	return listOfnumbers_
}


// this is a function for the Queen that will calculate The moves
func QueenAlgoritmo(yCambio_ int , xCambio_ int ) [][2]int {
	var listOfnumbers = [][2]int{}

	//variable tha store the position of the tab
	y1_ = y__
	x2_ = x__


	i_ := 0
	if i_ < 9{
		i_++

		// the variables than will make the chance


	}

	return listOfnumbers
}