package main

import "fmt"

var (
	matriz = [8][8]int{
		{-4, -7, -5, -9, -8, -5, -7, -4},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{4, 7, 5, 9, 8, 5, 7, 4},
	}
    matriz1 = [8][8]int{}
)

func main() {

	fmt.Println("hello")
    ivv := algorimoReina(matriz , 5, 4)
    iv := len(ivv)
    fmt.Println(ivv)

    for ip := 0 ; ip  < iv  ; ip++{
        matriz1[ivv[ip][1]][ivv[ip][0]] = 1
    }

    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            print(matriz1[i][j])
        }
        println()
    }

}

func algorimoReina(tablero [8][8]int, x int, y int) [][2]int {

	var letMoves = [][2]int{}

	x1 := x
	y1 := y
	x2 := x
	y2 := y
	x3 := x
	y3 := y
	x4 := x
	y4 := y
	x5 := x
	y5 := y
	x6 := x
	y6 := y
    x7 := x
	y7 := y
    x8 := x
	y8 := y

	for i := 0; i < 7; i++ {
		var iteration = [][2]int{}
		// para abajo
		y1 = y1 + 1
		x1 = x1 + 0
        iteration = append(iteration, [2]int{y1, y2})
		// para abajo derecha
		y2 = y2 + 1
		x2 = x2 + 1
        iteration = append(iteration, [2]int{y1, y2})
		// para abajo isquierda
		y3 = y3 + 1
		x3 = x3 - 1
        iteration = append(iteration, [2]int{y1, y2})
		// para arriba
		y4 = y4 - 1
		x4 = x4 + 0
        iteration = append(iteration, [2]int{y1, y2})
		// para arriba derecha
		y5 = y5 - 1
		x5 = x5 + 1
        iteration = append(iteration, [2]int{y1, y2})
		// para arriba isquierda
		y6 = y6 - 1
		x6 = x6 - 1
        iteration = append(iteration, [2]int{y1, y2})
		// para derecha
		y7 = y7 + 0
		x7 = x7 + 1
        iteration = append(iteration, [2]int{y1, y2})
		// para isquierda
		y8 = y8 + 0
		x8 = x8 - 1
        iteration = append(iteration, [2]int{y1, y2})

        for io := 0 ; io < len(iteration); io++{
            if matriz [iteration[io][0]][iteration[io][1]] == 0{
                letMoves = append(letMoves, [2]int{iteration[io][0],iteration[io][1]})
            }  
        }

	}
	return letMoves
}
