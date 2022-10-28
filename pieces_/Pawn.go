package pieces_

var (
	Pawn_Allowed_Moves__ =  map[int]Contructor{}
	theKey__ = false;
)

type Pawn__ struct {

	Y__ int 
	X__ int 
	Black__ [6] string 
	White__ [6] string 
 
	Board__ [9][9] string 
	KEYY bool
	
}

func (P * Pawn__) Pawn_Allowed_Play (){
	theKey__ = P.KEYY;


	if checker(P.White__ , P.Board__[P.Y__][P.X__]) {

		checker_Pawn(P.Y__ , P.X__ , -2 , 0, P.Board__ ,P.White__, P.Black__, false)
		checker_Pawn(P.Y__ , P.X__ , -1 , 0, P.Board__ ,P.White__, P.Black__, false)
		checker_Pawn(P.Y__ , P.X__ , -1 , -1, P.Board__ ,P.White__, P.Black__, true)
		checker_Pawn(P.Y__ , P.X__ , -1 , +1, P.Board__ ,P.White__, P.Black__, true)

	}	else{

		checker_Pawn(P.Y__ , P.X__ , +2 , 0, P.Board__ ,P.Black__,P.White__, false)
		checker_Pawn(P.Y__ , P.X__ , +1 , 0, P.Board__ ,P.Black__,P.White__, false)
		checker_Pawn(P.Y__ , P.X__ , +1 , -1, P.Board__,P.Black__,P.White__, true)
		checker_Pawn(P.Y__ , P.X__ , +1 , +1, P.Board__  ,P.Black__,P.White__, true)

	}

}

func checker_Pawn(y int , x int ,moviToA int , moviToB int , board [9][9] string , color1 [6] string , color2 [6] string , key bool){

	if y + moviToA >= 1 && y  + moviToA <= 7 && x + moviToB >= 1 && x + moviToB <= 8 {

		if y  + moviToA == 6 && color1 [0] == board [ y + moviToA][x + moviToB]  {
			Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__) +1] = Contructor{y -2 ,x}

		} 
		if y + moviToB == 1 &&  color2 [0] == board[y + moviToA][x + moviToB] {
			Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__) +1] = Contructor{ y + 2 , x }
		} 

        if board [y + moviToA ][x + moviToB ] == " - " && key  == false  {
			Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__) +1] = Contructor{y + moviToA, x + moviToB}            

		}
		if  key == true  && checker(color2, board[y + moviToA][x + moviToB])  && theKey__ == true && key == true {
            Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__) + 1] = Contructor{y + moviToA, x + moviToB}

        }

    }
  
}
