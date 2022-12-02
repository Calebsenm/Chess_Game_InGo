package Pieces


type pawm  struct{
	y int
	x int
	board [9][9]string
}

func movesCalculate(objet * pawm) [][]int{

	mv := moves(objet.y , objet.y, objet.board)
	return mv
}

func moves(y int , x int , board [9][9]string ) [][]int{
	
	let_moves := [][] int{}
	let_moves = append (let_moves,[]int{3,4})
	let_moves = append(let_moves, []int{3,4})
	let_moves = append(let_moves, []int{3,4})
	let_moves = append(let_moves, []int{3,4})

	return let_moves
}		