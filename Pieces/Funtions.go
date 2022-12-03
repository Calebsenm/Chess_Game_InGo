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