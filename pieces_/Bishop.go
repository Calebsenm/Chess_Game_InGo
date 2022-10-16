
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



func (iop * Bishop_) Bishop_AllowedMoves_bishop(){
    
    if checker(iop.White_pieces_1,iop.The_Board_Queen[iop.Y_Queen][iop.X_Queen]) == true{
      // this is UP right
      ichecher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.Black_pieces_2);
      //this is UP left
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.Black_pieces_2)
      // this is for the Down left
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.Black_pieces_2)
      // this is for the  Down right
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.Black_pieces_2)


    } else{
      // this is UP right
      ichecher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.White_pieces_2);
      //this is UP left
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.White_pieces_2)
      // this is for the Down left
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.White_pieces_2)
      // this is for the  Down right
      checher_Bishop(iop.Y_BY , iop.X_BY , -1, +1, iop.The_Board_Bishop, iop.White_pieces_2)
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

      if Board_[Y_BYY + iter_1][X_BYY + iter_2] == " - " || checker(TheColor_,Y_BYY + iter_1 , X_BYY + iter_2) == true{
        Bishop_allowedMoves_Queen__[len(Bishop_allowedMoves_Queen__)+1] = Contructor{Y_BYY + iter_1 , X_BYY + iter_2 }
      
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
