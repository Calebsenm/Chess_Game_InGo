package main

import (
	
	"fmt"
)
	
func main() {
		
	fmt.Println( "hello")	

	var list [] int =  []int { 1 ,4 , 3 , 5 }
	
	for uno , _ := range list {
	
		println(uno)
	}

	//list = nil
	list = list [:0]
	
	println()
	for uno , _ := range list {
	
		println(uno)
	}
}
