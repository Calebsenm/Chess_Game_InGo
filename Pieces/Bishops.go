package Pieces

import "fmt"


var(
	n_ = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b_ = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}
	let_moves1 = [][]int{}
	Color111 = [6] string {}



)
type Bishops_ struct{
	Y_ 				int
	X_				int
	Board_ [9][9]	string
	Color [6]  		string
}


func (objet Bishops_) MovesCalculate() [][]int{
	Color111 = objet.Color

	mv := movesBishop(objet.Y_ , objet.X_, objet.Board_)
	return mv
}
func movesBishop(y_ int , x_ int , board [9][9]string ) [][]int{
	
	
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
				let_moves1 = append(let_moves1, []int{ay + in1 , ax +in2})
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

