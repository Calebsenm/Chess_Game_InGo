package Pieces

// this is for check if color is equals

func colorChecker(piece string, color [6] string) bool{
	veri := false
	for i := 0; i < len(color); i++ {
		
		if color[i] == " " + piece + " " {
			veri = true
		} 
	}
	return veri
}


// Para verificar si el calculo se salìo de rango 
func checkEdge( y_ , x_ int ) (bool,bool){

	var valuesY [8] int;
	var valuesX [8] int;
	var key1 		bool
	var key2 		bool

	for i := 0; i < 8; i++ {
		valuesY[i] = i + 1
		valuesX[i] = i
	}

	for i := 0; i < 8; i++ {
		if y_ == valuesY[i]{
			key1 = true
		}
	}

	for i := 0; i < 8; i++ {
		if y_ == valuesX[i]{
			key2 = true
		}
	}

	return key1 , key2
}