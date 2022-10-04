package main

import (
	"fmt"
	"strings"
	"strconv"

)
  
func main() {
  
	var position string

	var TheYposition int
	var TheXposition int 

	for {
		fmt.Print("Please enter the position > ")
		fmt.Scanln(&position)

		// this is for separate the string 
		s := position
		v := strings.SplitAfter(s, "")
		
		// this is for become String to a int
		val,err := strconv.Atoi(v[1])
		if err !=  nil{
			
		}	else {
				
				if val > 0 && val <= 9{

					a := [8] string{"a","b","c","d","e","f","g","h"}  
					itero := false

					for io := 0; io < 8; io++ {
						if a[io] == v[0]{
							itero = true
						}

					}
					if itero== true{

						Change_letters_numbers := map[string]int { "a": 1,"b": 2,"c": 3,"d":4,"e" : 5,"f": 6,"g": 7,"h":8,}
						
						TheXposition = Change_letters_numbers[v[0]]
						TheYposition = val
						break	
					}

				
				}
		}
	
		
	}	

	fmt.Println("The position y",TheYposition )
	fmt.Println("The position x",TheXposition)



	// here is the next step
	fmt.Print("Where do you want to play > ")
	var position1 string
	fmt.Scanln(&position1)
	
}