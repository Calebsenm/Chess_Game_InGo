package pieces_

import "fmt"

var (
	Pawn_Allowed_Moves__ = map[int]Contructor{}
	theKey__             = false
)

type Pawn__ struct {
	Y__     int
	X__     int
	Black__ [6]string
	White__ [6]string

	Board__ [9][9]string
	KEYY    bool
}

func (P *Pawn__) Pawn_Allowed_Play() {
	theKey__ = P.KEYY

	if checker(P.White__, P.Board__[P.Y__][P.X__]) {
		fmt.Println("Loro XD ")
		fmt.Println(P.Y__-1, " ", P.Y__+1)
		fmt.Println(P.X__-1, " ", P.X__+1)

		//  this is for let the first 2 moves
		//  && P.White__[0] == P.Board__ [ P.Y__ ][ P.X__ ]
		if P.Y__ == 6 && check__If_Bouns_out(P.Y__-2, P.X__) {
			Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ - 2, P.X__}
		}

		if P.Board__[P.Y__-1][P.X__] == " - "{

			fmt.Println( check__If_Bouns_out(P.Y__-1, P.X__)  )
			Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ - 1, P.X__}
		}
		if checker(P.Black__, P.Board__[P.Y__-1][P.X__-1]) && check__If_Bouns_out(P.Y__-1, P.X__-1) {

			if theKey__ == true {
				Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ - 1, P.X__ - 1}

			} else {
				Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ - 1, P.X__ - 1}
			}

		}

		if checker(P.Black__, P.Board__[P.Y__][P.X__]) {

			if check__If_Bouns_out(P.Y__-1, P.X__+1) {
				if theKey__ == true {
					Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ - 1, P.X__ + 1}

				} else {
					Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ - 1, P.X__ + 1}
				}
			}

		}

	} else {

		//&&  P.White__ [0] == P.Board__[P.Y__ ][ P.Y__ ]
		if P.Y__ == 1 {
			Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ + 2, P.X__}

		}

		if P.Board__[P.Y__+1][P.X__] == " - " && check__If_Bouns_out(P.Y__+1, P.X__) {
			Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ + 1, P.X__}
		}

		if checker(P.Black__, P.Board__[P.Y__+1][P.X__+1]) && check__If_Bouns_out(P.Y__+1, P.X__+1) {

			if theKey__ == true {
				Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ + 1, P.X__ + 1}

			} else {
				Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ + 1, P.X__ + 1}
			}

		}
		if checker(P.Black__, P.Board__[P.Y__+1][P.X__-1]) && check__If_Bouns_out(P.Y__+1, P.X__-1) {

			if theKey__ == true {
				Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ + 1, P.X__ - 1}

			} else {
				Pawn_Allowed_Moves__[len(Pawn_Allowed_Moves__)+1] = Contructor{P.Y__ + 1, P.X__ - 1}
			}

		}

	}

}
