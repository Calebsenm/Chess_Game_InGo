
package main

import(
	"fmt"

)
var (
	allowedMoves = map[int] position1{  }
)
type position1 struct{
	A, B ,h , a int
	p string 
}
func main(){

	allowedMoves[len(allowedMoves)+1] = position1{ 2 , 3,7,7,"XD"}
	allowedMoves[len(allowedMoves)+1] = position1{ 3 , 3,7,7,"XD"}
	allowedMoves[len(allowedMoves)+1] = position1{ 4 , 3,7,7,"XD"}

	for i := 0; i < len(allowedMoves) +5; i++ {
		fmt.Println(allowedMoves[i])
	}
}