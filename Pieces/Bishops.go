package Pieces

type Bishops_ struct {

	Y int 
	X int
	Board_ [9][9]string
	Color [6]  		string
}

func (objet Bishops_) MovesCalculate() [][]int{
	mv := movesBishop(objet.Y , objet.X, objet.Board_)
	return mv
}
func movesBishop(y int , x int , board [9][9]string ) [][]int{
	
	let_moves := [][] int{}
	let_moves = append (let_moves,[]int{y,x})

	

	return let_moves
}		
