
// this will be the most bigest proyect 
// i dot know how to create the game but i going to build the ches game with ia

// i start this proyect in the date  2/10/2022 

// I wish you the best caleb of the future you are capable do not give up. you are going to achieve it you are able


package main


import(
	"fmt"
)

var(

	B = [6] string {"\u2659","\u2656","\u2655","\u2658","\u2654","\u2657"}
	N = [6] string {"\u265F","\u265C","\u265B","\u265E","\u265A","\u265D"}


	Board = [81] string{
		
		" H "," "+B[1]+" "," "+B[3]+" "," "+B[5]+" "," "+B[4]+" "," "+B[2]+" "," "+B[5]+" ", " "+B[3]+" ", " "+B[1]+" ",
		" G "," "+B[0]+" "," "+B[0]+" "," "+B[0]+" "," "+B[0]+" "," "+B[0]+" "," "+B[0]+" ", " "+B[0]+" ", " "+B[0]+" ",
		" F "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" E "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" D "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" C "," - "," - "," - "," - "," - "," - ", " - ", " - ",
		" B "," "+N[0]+" "," "+N[0]+" "," "+N[0]+" "," "+N[0]+" "," "+N[0]+" "," "+N[0]+" ", " "+N[0]+" ", " "+N[0]+" ",
		" A "," "+N[1]+" "," "+N[3]+" "," "+N[5]+" "," "+N[4]+" "," "+N[2]+" "," "+N[5]+" ", " "+N[3]+" ", " "+N[1]+" ",
		"   "," 1 "," 2 "," 3 " ," 4 "," 5 " ," 6 "," 7 "," 8 ",		
	
	}
)

func main()  {

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

 
	// iterator := 0;

	// for {

		


	// 	iterator++
	// }

	
}