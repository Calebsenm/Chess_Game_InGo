
package pieces_

var (
  Bishop_allowedMoves_ = map[int]Contructor{}
)

type Bishop_ struct {
  Y_BY int
  X_BY int 
  Black_pieces_2  [6]string
  White_pieces_2  [6]string
  The_Board_Bishop  [9][9]string
}




func (iop * Bishop_) Bishop_AllowedMoves_bishop_ToPLay(){
    
    if checker_Bishop(iop.White_pieces_2,iop.The_Board_Bishop[iop.Y_BY][iop.X_BY]){
      // this is UP right
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.Black_pieces_2);
      //this is UP left
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, -1, iop.The_Board_Bishop, iop.Black_pieces_2)
      // this is for the Down left
      checher_Bishop(iop.Y_BY , iop.X_BY , +1, +1, iop.The_Board_Bishop, iop.Black_pieces_2)
      // this is for the  Down right
      checher_Bishop(iop.Y_BY , iop.X_BY , +1, -1, iop.The_Board_Bishop, iop.Black_pieces_2)


    } else{
      // this is UP right
      checher_Bishop(iop.Y_BY , iop.X_BY , +1, +1, iop.The_Board_Bishop, iop.White_pieces_2);
      //this is UP left
      checher_Bishop(iop.Y_BY , iop.X_BY , +1, -1, iop.The_Board_Bishop, iop.White_pieces_2)
      // this is for the Down left
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.White_pieces_2)
      // this is for the  Down right
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, -1, iop.The_Board_Bishop, iop.White_pieces_2)
    }
  } 


  func checher_Bishop(Y_BYY int , X_BYY int, MoveA int , MoveB int ,Board_ [9][9] string, TheColor_ [6] string)  {
 
    
    iter_1 := MoveA
    iter_2 := MoveB

    for{
      // for the bouns out of range
      if  Y_BYY + iter_1 == -1 || Y_BYY + iter_1 == 8 || X_BYY +  iter_2 == 0 || X_BYY + iter_2 == 9{
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

// Function to calculate the color of the pice to eat
func checker_Bishop(Piece_Color [6]string, Color_Position string) bool {

	xddd := false
	for i := 0; i < len(Piece_Color); i++ {
		if  " " + Piece_Color[i] + " " ==   Color_Position  {
			xddd = true
			
			
		}

	}
	return xddd
}