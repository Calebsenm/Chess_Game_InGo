package Pieces

import "fmt"


var (
	n_____ = [6]string{"\u2659", "\u2656", "\u2655", "\u2658", "\u2654", "\u2657"}
	b_____ = [6]string{"\u265F", "\u265C", "\u265B", "\u265E", "\u265A", "\u265D"}

	x_____ int
	y_____ int

	listOfNumbersKing = [][2]int{}
	board_____           [9][9]string

    listFight1 [][2] int
    listFight2 [][2] int 

    TheKey = false 
)

type King_ struct {
	Y_____     int
	X_____     int
	Board_____ [9][9]string
	Color_____ [6]string
    ListFight1_____[][2] int 
    ListFight2_____[][2] int 
}

func (Kin * King_) MovesKing() [][2]int {

	var vacia [][2]int
	listOfNumbersKing = vacia

	// this save variable
	x_____ = Kin.X_____
	y_____ = Kin.Y_____
	board_____ = Kin.Board_____
    listFight1 = Kin.ListFight1_____ 
    listFight2 = Kin.ListFight2_____ 


	KingAlgoritmo( -1 ,  0, false )
	KingAlgoritmo( -1 , -1, false )
	KingAlgoritmo( -1 , +1, false )
	KingAlgoritmo(  0 , -1, false )
	KingAlgoritmo(  0 , +1, false )
	KingAlgoritmo(  1 , -1, false )
	KingAlgoritmo(  1 ,  1, false )
	KingAlgoritmo(  1 ,  0, false )
		
    // this is for the check rosk 
    KingAlgoritmo(  0 , -2 , true )
    KingAlgoritmo(  0 ,  2 , true ) 

	return listOfNumbersKing
}

func KingAlgoritmo(yChance_ int , xChance_ int , question bool) {

	atactPiece_____ := board_____[y_____][x_____]
	//fmt.Println(atactPiece____)
	y1_____ := y_____
	x1_____ := x_____

	y1_____ = y1_____ + yChance_
	x1_____ = x1_____ + xChance_ 

    TheKey = false 

	if fueraRango_____(y1_____, x1_____) {

        if question == false {
    
		    //fmt.Println(fichaAtaca____(atactPiece____))

		    // white 
		    if fichaAtaca_____(atactPiece_____) == 1 {
			    
				fmt.Println( listNum_____(y1_____,x1_____,listFight1))

				if board_____[y1_____][x1_____] == " - " && listNum_____(y1_____,x1_____,listFight1) == false{
					listOfNumbersKing = append(listOfNumbersKing, [2]int{y1_____, x1_____})
				}
	
				
				if fichas_____(board_____[y1_____][x1_____], n_____) && listNum_____(y1_____,x1_____,listFight1) == false{
		
				    listOfNumbersKing = append(listOfNumbersKing, [2]int{y1_____, x1_____})
			    }
	

		    } 
		    // black 
		    if fichaAtaca_____(atactPiece_____) == 2 {


				if board_____[y1_____][x1_____] == " - " && listNum_____(y1_____,x1_____,listFight2) == false{
					listOfNumbersKing = append(listOfNumbersKing, [2]int{y1_____, x1_____})
				}

				
			    if fichas_____(board____[y1_____][x1_____], b_____)&& listNum_____(y1_____,x1_____,listFight2) == false{
				    listOfNumbersKing = append(listOfNumbersKing, [2]int{y1_____, x1_____})
			
			    }
	
		    }
        } 

        if question == true{

            ay   ,   ax :=  y1_____ ,  x1_____
            validation1 :=  board_____[ay][ax] == " - " 
            validation2 :=  board_____[ay][ax +1] == " - " || board_____[ay][ax +1] == " " +n_____[1] + " " || board_____[ay][ax +1] == " " +b_____[1] + " " 
            validation3 :=  board_____[ay][ax -1] == " - "
            validation4 := false    
                    
    
             // white 
		    if fichaAtaca_____(atactPiece_____) == 1 {
			
                aValidation1 := listNum_____( ay ,ax,listFight1)
                aValidation2 := listNum_____( ay ,ax-1,listFight1)
                
                if board_____[ay][ax+1]== " - "{
                    aValidation3 := listNum_____( ay , ax+1,listFight1)
                    aValidation4 := listNum_____( ay , ax + 2 ,listFight1)
                    
                    if  aValidation1  == false &&  aValidation2 == false &&  aValidation3 == false && aValidation4 == false {
                        validation4 = true
                    }

                } else {
                     aValidation4 := listNum_____( ay , ax - 2 ,listFight1) 

                    if aValidation1 == false &&  aValidation2 == false && aValidation4  == false{
                        validation4 = true
                    }
                } 
                

            } 
		    // black 
		    if fichaAtaca_____(atactPiece_____) == 2 {
			
                 aValidation1 := listNum_____( ay ,ax ,listFight2)
                 aValidation2 := listNum_____( ay ,ax-1,listFight2)
                               
                 if board_____[ay][ax+1]== " - "{
                    aValidation3 := listNum_____( ay , ax+1,listFight2)
                    aValidation4 := listNum_____( ay , ax + 2 ,listFight2) 


                   if  aValidation1 == false &&  aValidation2  == false &&  aValidation3 == false && aValidation4 == false{

                            validation4 = true
                   }
               

                } else {
                    
                    aValidation4 := listNum_____( ay , ax - 2 ,listFight2) 
                    if  aValidation1  == false &&  aValidation2 == false && aValidation4 == false {
                            validation4 = true 
                    }
                } 

		    }

            allValidations := validation1 && validation2 && validation3 && validation4


            if allValidations {
			    listOfNumbersKing = append(listOfNumbersKing, [2]int{y1_____, x1_____})
                TheKey = true 

            }

        }

	}
}

// out of the range
func fueraRango_____(y, x int) bool {
	keyValue := false

	if y >= 0 && y <= 7 && x >= 1 && x <= 8 {
		keyValue = true
	}
	return keyValue
}

// what pice is fithg
func fichaAtaca_____(atacaX_____ string) int {

	var num int

	for i := 0; i < 6; i++ {
		//fmt.Println(atacaX____ ,"  - ", " " + b____[i] + " ")
		if atacaX_____ == " " + b_____[i] + " " {
			num = 1
			break

		} else if atacaX_____ == " " + n_____[i] + " " {
			num = 2
			break

		}
	}

	return num
}

// to control the opponent's chip attack
func fichas_____(ficha_____ string, list_____ [6]string) bool {

	for i := 0; i < len(list_____); i++ {
		//fmt.Println( " "+list____[i]+" ", " - ", ficha____)
		if " "+list_____[i]+" " == ficha_____{

			return true
		}
	}
	return false
}


// para validar si los numeros estan en la lista
func listNum_____( num1_ int , num2_ int , lisNum1_ [][2] int )bool {
    validation := false

    for i := 0; i < len (lisNum1_) ; i ++{
        
        if num1_ == lisNum1_[i][0] && num2_ == lisNum1_ [i][1]{
            validation = true;

			fmt.Println("MataPuercaso")
        }
    }
    return validation;
    
}


