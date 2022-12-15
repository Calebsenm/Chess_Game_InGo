package Pieces

import (

	"fmt"
)



var(
	n__ = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b__ = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}
	board__ [9][9] string

	x__ int 
	y__ int

	listOfnumbers = [][2]int{}
)

type Queen_ struct{
	Y_ 				int
	X_				int
	Board_ [9][9]	string
	Color [6]  		string
}

func (ob Queen_ ) MovesQueen()  [][2]int {
	
	x__ = ob.X_
	y__ = ob.Y_
	board__ = ob.Board_
	//this is forUp
	QueenAlgoritmo(-1,0)
	//Up right
	QueenAlgoritmo(-1,+1)
	//Up left
	QueenAlgoritmo(-1 ,-1)
	//right
	QueenAlgoritmo(0, +1)
	//left
	QueenAlgoritmo(0, -1)
	//Dowm
	QueenAlgoritmo(+1, 0)
	//Dowm left
	QueenAlgoritmo(+1, -1)
	//Dowm right
	QueenAlgoritmo(+1, 1)

	
	return listOfnumbers
}


// this is a function for the Queen that will calculate The moves
func QueenAlgoritmo(yCambio_ int , xCambio_ int ){
	
	ataqueFicha1 := board__[y__][x__]
	//variable tha store the position of the tab
	y1_ := y__
	x1_ := x__

	i_ := 0
	for i_ < 9{
		i_++

		// the variables than will make the chance
		y1_ = y1_ +  yCambio_
		x1_ = x1_ + xCambio_

		if fueraRango(y1_ , x1_){
			if board__[y1_][x1_] == " - " {
				listOfnumbers = append(listOfnumbers,[2]int{y1_,x1_})
			}

			//para blancas
			if fichaAtaca(ataqueFicha1) == 1 {
				fmt.Println(board__[y1_][x1_]," ",y1_," ",x1_)
				if ficha(board__[y1_][x1_], n__) {

					listOfnumbers = append(listOfnumbers, [2]int{y1_, x1_})
					break
					
				}else if  ficha(board__[y1_][x1_], b__){
					break
				}
					
			} else if fichaAtaca(ataqueFicha1) == 2 {
				if ficha(board__[y1_][x1_], b__) {
					listOfnumbers = append(listOfnumbers, [2]int{y1_, x1_})
					break
				}	else if  ficha(board__[y1_][x1_], b__){
					break
				}
			}
		}

	}


}


// para controlar el fuera de rango
func fueraRango(y1, x1 int) bool {

	dde := false
	if y1 >= 0 && y1 <= 7 && x1 >= 1 && x1 <= 8 {
		dde = true
	}
	return dde
}

// ver la ficha que ataca
func fichaAtaca(atacaX string) int {

	var numero int

	for i := 0; i < 6; i++ {
		if atacaX == b__[i] {
			numero = 1
			break
		} else if  atacaX == n__[i]{
			numero = 2
			break
		}
	}
	return numero
}


// para controlar el ataque ficha contraria
func ficha(ficha string, list [6]string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == " " + ficha +" "{
			return true
		}
	}
	return false
}
