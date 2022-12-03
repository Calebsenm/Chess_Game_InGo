package Pieces

var(
	B = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}
)  


type Pawm_ struct{
	Y 				int
	X 				int
	Board_ [9][9]	string
	Color [6]  		string
}

func (objet Pawm_) MovesCalculate() [][]int{

	mv := movesPawm(objet.Y , objet.X, objet.Board_)
	return mv
}

func movesPawm(y int , x int , board [9][9]string ) [][]int{
	
	let_moves := [][] int{}
	let_moves = append (let_moves,[]int{y,x})

	// chek color is white or black
	colorChecker(board[y][x],B)

	return let_moves
}		
