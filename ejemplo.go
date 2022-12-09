package main

import "fmt"

var (

    matriz = [8][8]int{
        {-4, -7, -5, -9, -8, -5, -7, -4},
        {-1, -1, -1, -1, -1, -1, -1, -1},
        {0 ,  0,  0,  0,  0,  0,  0,  0},
        {0 ,  0,  0,  0,  0,  0,  0,  0},
        {0 ,  0,  0,  0,  0,  0,  0,  0},
        {0 ,  0,  0,  0,  0,  0,  0,  0},
        {1 , 1 , 1 , 1 , 1 , 1 , 1 , 1 },
        { 4,  7,  5,  9,  8,  5,  7,  4},
      }
)



func main() {

    fmt.Println("hello")
}

func algorimoReina(tablero [8][8] int , x int ,y int ) [][2] int {

    var  letMoves = [][2] int{} 

    x1 := x; y1 := y 
    x2 := x; y2 := y
    x3 := x; y3 := y
    x4 := x; y4 := y
    x5 := x; y5 := y
    x6 := x; y6 := y
 
    for i := 0; i < 8;i++{
 
        if x1 >= 0 && x1 <= 8 && y1 >= 0 && y1 <= 8{
            
        }
 
        y1++;y2--

    }
    return letMoves
}
  

