
// this will be the most bigest proyect 
// i dot know how to create the game but i going to build the ches game with ia

// i start this proyect in the date  2/10/2022 

// I wish you the best caleb of the future you are capable do not give up. you are going to achieve it you are able


package main


import(
	"fmt"
)

var(

	Board = [81] string{
		
		" H "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" G "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" F "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" E "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" D "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" C "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" B "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" A "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		"   "," 1 "," 2 "," 3 " ," 4 "," 5 " ," 6 "," 7 "," 8 ",		
	
	}
)

func main()  {
	fmt.Println("I am crazy");

	ii := 0;

	for i := 0; i < 81; i++ {
		fmt.Print("")

		if ii < 9{
			fmt.Print(Board[i])
		}else{
			
			fmt.Println()
			ii = 0;
			fmt.Print(Board[i])
		}
		ii++;

	}
}