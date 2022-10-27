package pieces_


var (
	Pawn_Allowed_Moves__ =  map[int]Contructor{}
)

type Pawn__ struct {

	Y__ int 
	X__ int 
	White__ [6] string 
 	Black__ [6] string 

	Board__ [9][9] string 

}

func (P * Pawn__) Pawn_Allowed_Play (){

	if checker(P.White__ , P.Board__[P.Y__][P.X__]) {

	}	else{

	}

}

func checker_Pawn(y int , x int ,moviToA int , moviToB int , board [9][9] string , color [6] string){

	if y + moviToA >= 1 && y {


    
    }
  
}
