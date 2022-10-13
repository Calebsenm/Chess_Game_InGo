package pieces_



// this is for the structure of a diccionary
type Contructor struct{
	This_A int 
	This_B int 
}

// Function to calculate the color of the pice to eat
func checker(Piece_Color [6]string, Color_Position string) bool {

	xddd := false
	for i := 0; i < len(Piece_Color); i++ {
		if  " " + Piece_Color[i] + " " ==   Color_Position  {
			xddd = true
			
			
		}

	}
	return xddd
}


// this is for check the  numers of the edge

func check__EdgeNumber(Only_Y [8] int, Only_X [8] int,value_Y int ,value_X int ) bool{
	XD1 := false
	XD2 := false

	for i := 0; i < len(Only_Y); i++ {
		if Only_Y[i] == value_Y{
			XD1 = true
			
		}
	}

	for j := 0; j < len(Only_X); j++ {
		if Only_X[j] == value_X{
			XD2 = true
		
		}
	}

	return XD1 && XD2

}