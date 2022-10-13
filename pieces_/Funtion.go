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
