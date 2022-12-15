package main

import "fmt"

var (
	BL = [6]int{2,  7,  1,  8, 9,  4}
	NL = [6]int{-2,-7, -1, -8,-9, -4}

	matriz = [8][8]int{
		{8 , 0 , 0 , 0 , 0 , 0 , 0  , 0 },
		{0 , 0 , 0 , 0 , 0 , 0 , 0  , 0 },
		{0 , 0 , 0 , 0 , 0 , 0 , 0  , 0 },
		{0 , 0 , 0 , 0 , 0 , 0 , 0  , 0 },
		{-1, 0 , 0 , 0 , 0 , 0 , 0  , 0 },
		{0 , 0 , 0 , 0 , 0 , 0 ,  0 , 0 },
		{0 , 0 , 0 , 0 , 0 , 0 ,  0 , 0 },
		{0 , 0 , 0 , 0 , 0 , 0 ,  0 , 0 },
	}
	matriz1 = [8][8]int{}
)

func main() {

	fmt.Println("hello")
	ivv := algorimoReina(matriz, 0, 0)
	iv := len(ivv)
	fmt.Println(ivv)

	fmt.Println()
	for ip := 0; ip < iv; ip++ {
		matriz1[ivv[ip][1]][ivv[ip][0]] = 1
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			print(matriz1[i][j], " ")
		}
		println()
	}

}

func algorimoReina(tablero [8][8]int, x int, y int) [][2]int {

	fichaAtacante := tablero[y][x]
	x1 := x
	y1 := y
	var iteration = [][2]int{}
	fmt.Println(fichaAtacante)

	i := 0;

	for  i < 8{
		i++ 
		// Esta para ir arriba abajo
		y1 = y1 + 1
		x1 = x1 + 0
		if fueraRango(y1, x1) {

			if matriz[y1][x1] == 0 {
				iteration = append(iteration, [2]int{y1, x1})
			}
			//para blancas
			if fichaAtaca(fichaAtacante) == 1 {
				fmt.Println(matriz[y1][x1]," ",y1," ",x1)
				if ficha(matriz[y1][x1], NL) {
					iteration = append(iteration, [2]int{y1, x1})
					break
				}else if  ficha(matriz[y1][x1], BL){
					break
				}
					
			} else if fichaAtaca(fichaAtacante) == 2 {
				if ficha(matriz[y1][x1], BL) {
					iteration = append(iteration, [2]int{y1, x1})
					break
				}	else if  ficha(matriz[y1][x1], BL){
					break
				}
			}

		}
	}

	// x2 := x
	// y2 := y

	// for i := 0; i < 8; i++ {
	// 	// para abajo derecha
	// 	y2 = y2 + 1
	// 	x2 = x2 + 1

	// 	if fueraRango(y2, x2) {
	// 		if matriz[y2][x2] == 0 {
	// 			iteration = append(iteration, [2]int{y2, x2})
	// 		}
	// 		// para blancas
	// 		if fichaAtaca(fichaAtacante) == 1 {
	// 			if ficha(matriz[y2][x2], NL) {
	// 				iteration = append(iteration, [2]int{y2, x2})
	// 				break
	// 			}
	// 		} else if fichaAtaca(fichaAtacante) == 2 {
	// 			if ficha(matriz[y2][x2], BL) {
	// 				iteration = append(iteration, [2]int{y2, x2})
	// 				break
	// 			}
	// 		}
	// 	}

	// }

	// x3 := x
	// y3 := y

	// for i := 0; i < 8; i++ {

	// 	// para abajo isquierda
	// 	y3 = y3 + 1
	// 	x3 = x3 - 1

	// 	if fueraRango(y3, x3) {
	// 		if matriz[y3][x3] == 0 {
	// 			iteration = append(iteration, [2]int{y3, x3})
	// 		}
	// 		// para blancas
	// 		if fichaAtaca(fichaAtacante) == 1 {
	// 			if ficha(matriz[y3][x3], NL) {
	// 				iteration = append(iteration, [2]int{y3, x3})
	// 				break
	// 			}
	// 		} else if fichaAtaca(fichaAtacante) == 2 {
	// 			if ficha(matriz[y3][x3], BL) {
	// 				iteration = append(iteration, [2]int{y3, x3})
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// x4 := x
	// y4 := y

	// for i := 0; i < 8; i++ {
	// 	// para arriba
	// 	y4 = y4 - 1
	// 	x4 = x4 + 0
	// 	if fueraRango(y4, x4) {
	// 		if matriz[y4][x4] == 0 {
	// 		  iteration = append(iteration, [2]int{y4, x4})
	// 		}
	// 		// para blancas
	// 		if fichaAtaca(fichaAtacante) == 1 {
	// 		  if ficha(matriz[y4][x4], NL) {
	// 			iteration = append(iteration, [2]int{y4, x4})
	// 			break
	// 		  }
	// 		} else if fichaAtaca(fichaAtacante) == 2 {
	// 		  if ficha(matriz[y4][x4], BL) {
	// 			iteration = append(iteration, [2]int{y4, x4})
	// 			break
	// 		  }
	// 		}
	// 	  }
		  
	// }

	// x5 := x
	// y5 := y

	// for i := 0; i <= 7; i++ {
	// 	// para arriba derecha
	// 	y5 = y5 - 1
	// 	x5 = x5 + 1
	// 	if fueraRango(y5, x5) {
	// 		if matriz[y5][x5] == 0 {
	// 		  iteration = append(iteration, [2]int{y5, x5})
	// 		}
	// 		// para blancas
	// 		if fichaAtaca(fichaAtacante) == 1 {
	// 		  if ficha(matriz[y5][x5], NL) {
	// 			iteration = append(iteration, [2]int{y5, x5})
	// 			break
	// 		  }
	// 		} else if fichaAtaca(fichaAtacante) == 2 {
	// 		  if ficha(matriz[y5][x5], BL) {
	// 			iteration = append(iteration, [2]int{y5, x5})
	// 			break
	// 		  }
	// 		}
	// 	  }
		  
	// }

	// x6 := x
	// y6 := y

	// for i := 0; i < 8; i++ {
	// 	// para arriba isquierda
	// 	y6 = y6 - 1
	// 	x6 = x6 - 1
	// 	if fueraRango(y6, x6) {
	// 		if matriz[y6][x6] == 0 {
	// 		  iteration = append(iteration, [2]int{y6, x6})
	// 		}
	// 		// para blancas
	// 		if fichaAtaca(fichaAtacante) == 1 {
	// 		  if ficha(matriz[y6][x6], NL) {
	// 			iteration = append(iteration, [2]int{y6, x6})
	// 			break
	// 		  }
	// 		} else if fichaAtaca(fichaAtacante) == 2 {
	// 		  if ficha(matriz[y6][x6], BL) {
	// 			iteration = append(iteration, [2]int{y6, x6})
	// 			break
	// 		  }
	// 		}
	// 	  }
		  

	// }

	// x7 := x
	// y7 := y

	// for i := 0; i < 8; i++ {
	// 	// para derecha
	// 	y7 = y7 + 0
	// 	x7 = x7 + 1
	// 	if fueraRango(y7, x7) {
	// 		if matriz[y7][x7] == 0 {
	// 		  iteration = append(iteration, [2]int{y7, x7})
	// 		}
	// 		// para blancas
	// 		if fichaAtaca(fichaAtacante) == 1 {
	// 		  if ficha(matriz[y7][x7], NL) {
	// 			iteration = append(iteration, [2]int{y7, x7})
	// 			break
	// 		  }
	// 		} else if fichaAtaca(fichaAtacante) == 2 {
	// 		  if ficha(matriz[y7][x7], BL) {
	// 			iteration = append(iteration, [2]int{y7, x7})
	// 			break
	// 		  }
	// 		}
	// 	  }
	// }
	// x8 := x
	// y8 := y
	// for i := 0; i < 8; i++ {
	// 	// para isquierda
	// 	y8 = y8 + 0
	// 	x8 = x8 - 1
	// 	if fueraRango(y8, x8) {
	// 		if matriz[y8][x8] == 0 {
	// 		  iteration = append(iteration, [2]int{y8, x8})
	// 		}
	// 		// para blancas
	// 		if fichaAtaca(fichaAtacante) == 1 {
	// 		  if ficha(matriz[y8][x8], NL) {
	// 			iteration = append(iteration, [2]int{y8, x8})
	// 			break
	// 		  }
	// 		} else if fichaAtaca(fichaAtacante) == 2 {
	// 		  if ficha(matriz[y8][x8], BL) {
	// 			iteration = append(iteration, [2]int{y8, x8})
	// 			break
	// 		  }
	// 		}
	// 	  }

	// }

	return iteration
}

// para controlar el ataque ficha contraria
func ficha(ficha int, list [6]int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == ficha {
			return true
		}
	}

	return false
}

// para controlar el fuera de rango
func fueraRango(y1, x1 int) bool {

	dde := false
	if y1 >= 0 && y1 <= 7 && x1 >= 0 && x1 <= 7 {
		dde = true
	}
	return dde

}

// ver la ficha que ataca
func fichaAtaca(atacaX int) int {

	var numero int

	for i := 0; i < 6; i++ {
		if atacaX == BL[i] {
			numero = 1
			break
		} else if  atacaX == NL[i]{
			numero = 2
			break
		}
	}
	return numero
}
