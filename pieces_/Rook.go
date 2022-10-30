package pieces_

var(

	Roock_allovedMoves_ = map[int]Contructor{}
)


type Rook_ struct {

	Y_ int
	X_ int

	Black_pieces___[6] string
	White_pieces___[6] string

	The_Board_Roock [9][9] string

}

func (a * Rook_ ) Roock_allovedMoves_Play(){

	if checker_Bishop(a.White_pieces___ ,a.The_Board_Roock[a.Y_][a.X_]){

		// up 
		checher_Roock_Moves(a.Y_ , a.X_ , -1 , 0, a.The_Board_Roock, a.Black_pieces___)
		//this is left 
		checher_Roock_Moves(a.Y_ , a.X_ , 0 , -1, a.The_Board_Roock, a.Black_pieces___)
		// this is for righ
		checher_Roock_Moves(a.Y_ , a.X_  , 0 , +1,  a.The_Board_Roock,a.Black_pieces___)
		// dowm
		checher_Roock_Moves(a.Y_ , a.X_  , +1, 0 ,  a.The_Board_Roock,a.Black_pieces___)
  

	}	else{
		
		 // this is UP right
		 checher_Roock_Moves(a.Y_ , a.X_ , +1, 0 ,  a.The_Board_Roock, a.White_pieces___)
		 //this is UP left
		 checher_Roock_Moves(a.Y_ , a.X_  , 0, -1, a.The_Board_Roock,a.White_pieces___ )
		 // this is for the Down left
		 checher_Roock_Moves(a.Y_ , a.X_ , 0 , +1,  a.The_Board_Roock ,a.White_pieces___)
		 // this is for the  Down right
		 checher_Roock_Moves(a.Y_ , a.X_  , -1, 0, a.The_Board_Roock , a.White_pieces___)
	  
	}
}

// Function to calculate the color of the pice to eat
func checker_Roock(Piece_Color [6]string, Color_Position string) bool {

	xddd := false
	for i := 0; i < len(Piece_Color); i++ {
		if  " " + Piece_Color[i] + " " ==   Color_Position  {
			xddd = true
		}

	}
	return xddd
}
func checher_Roock_Moves(Y_BYY int , X_BYY int, MoveA int , MoveB int ,Board_ [9][9] string, TheColor_ [6] string)  {
 
    
    iter_1 := MoveA
    iter_2 := MoveB

    for{
      // for the bouns out of range
    //   if  Y_BYY + iter_1 >= 0 && Y_BYY + iter_1 <=  7 && X_BYY +  iter_2 <= 0 &&  X_BYY + iter_2 <= 8{
    //    break 
    //   }
	if  Y_BYY + iter_1 <0 || Y_BYY + iter_1 >  8 || X_BYY +  iter_2 < 0 &&  X_BYY + iter_2 > 8 {
		break 
	   }

      if Board_[Y_BYY + iter_1][X_BYY + iter_2] == " - " || checker_Bishop(TheColor_,Board_[Y_BYY + iter_1][X_BYY + iter_2]){
        Bishop_allowedMoves_[len(Bishop_allowedMoves_)+1] = Contructor{Y_BYY + iter_1 , X_BYY + iter_2 }
      
      } else {
          break
      }

      // this is for increase of the move 
      if MoveA == 1 {
        iter_1 ++
      } else{
        iter_1 --

      }
  
      if MoveB == 1{
        iter_2 ++
      } else {
        iter_2 --
      }
    }
   
}