package pieces_

import "fmt"

var (
	Queen_allowedMoves_Queen___ = map[int]Contructor{}
)

type Queen_ struct {


	Y_Queen         int
	X_Queen         int
	Black_pieces_1  [6]string
	White_pieces_1  [6]string
	The_Board_Queen [9][9]string
}

func (b *Queen_) Queen_allowedMoves_Queen_play() {

	fmt.Println(b.The_Board_Queen[b.Y_Queen][b.X_Queen])

	// this is for the white pieces
	if checker(b.White_pieces_1,b.The_Board_Queen[b.Y_Queen][b.X_Queen]){
		// this is for the up moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen, 0 , -1,"UP",b.The_Board_Queen,b.Black_pieces_1)
		//this is for the up right moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,+1, -1, "UPR", b.The_Board_Queen,b.Black_pieces_1)
		// this is for the up left moves 
		queen_Algoritmit(b.Y_Queen,b.X_Queen,-1, -1,"UPL",b.The_Board_Queen,b.Black_pieces_1)
		// this is for the right moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,+1,0,"Right",b.The_Board_Queen,b.Black_pieces_1)
		// this is for the left moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,-1,0,"Left",b.The_Board_Queen,b.Black_pieces_1)
		// this is for the lef down moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,-1,+1,"DownL",b.The_Board_Queen,b.Black_pieces_1)
		// this is for the right dowm moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,+1, +1,"DownR",b.The_Board_Queen,b.Black_pieces_1)
		// this is for the down
		queen_Algoritmit(b.X_Queen,b.X_Queen,0,+1,"Down",b.The_Board_Queen,b.Black_pieces_1)

		// this is for the black pieces
	}	else{
		
		// this is for the up moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen, -1 , 0,"UP",b.The_Board_Queen,b.White_pieces_1)
		//this is for the up right moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,-1, +1, "UPR", b.The_Board_Queen,b.White_pieces_1)
		// this is for the up left moves 
		queen_Algoritmit(b.Y_Queen,b.X_Queen,-1, -1,"UPL",b.The_Board_Queen,b.White_pieces_1)
		// this is for the right moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,0,+1,"Right",b.The_Board_Queen,b.White_pieces_1)
		// this is for the left moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,0,-1,"Left",b.The_Board_Queen,b.White_pieces_1)
		// this is for the lef down moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,+1,-1,"DownL",b.The_Board_Queen,b.White_pieces_1)
		// this is for the right dowm moves
		queen_Algoritmit(b.Y_Queen,b.X_Queen,+1, +1,"DownR",b.The_Board_Queen,b.White_pieces_1)
		// this is for the down
		queen_Algoritmit(b.X_Queen,b.X_Queen,+1,0,"Down",b.The_Board_Queen,b.White_pieces_1)
		
	}


}

func queen_Algoritmit(The__Y int, The__X int, Move_X int , Move_Y int, The_board_edge string , Board___ [9][9] string , TheColor__[6] string) {

	

	var The_onlyMoves_y =  [8] int {0,1,2,3,4,5,6,7}
	var The_onlyMoves_x =  [8]int {1,2,3,4,5,6,7,8}

	counter1 := 0
	counter2 := 0

	for {
		if Move_Y == 1 {  counter1++;  }	else{ counter1--;  }
		if Move_X == 1 { counter2 ++;  }	else{ counter2 --; }
		

		if check__EdgeNumber(The_onlyMoves_y,The_onlyMoves_x,The__Y + counter1,The__X + counter2){

			fmt.Println(Board___[The__Y + counter1 ][The__X + counter2])
			if Board___[The__Y + counter1 ][The__X + counter2] == " - " {
				Queen_allowedMoves_Queen___[len(Queen_allowedMoves_Queen___)+1]= Contructor{The__Y + counter1, The__X + counter2}
			
			
			}	else{

					if checker(TheColor__, Board___[The__Y + counter1][The__X + counter2]){
						Queen_allowedMoves_Queen___[len(Queen_allowedMoves_Queen___)+1]= Contructor{The__Y + counter1, The__X + counter2}
						
						
						break

					}	else{
							break
					}
			}
		}	else{  
				// fmt.Println(counter1 + The__Y ,  " "  ,counter2 + The__X)
				break; 
		}

		
		
	}

	counter1 = 0
	counter2 = 0
}