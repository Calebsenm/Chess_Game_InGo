package Functions

import (
	"fmt"
	"strconv"
	"strings"
)

// this funtion going return the coordinates x y changing the words for numbers 

/*


*/

func CambioLetraNumero(player_name string , option string ) (int, int) {

	values := [2]int{}
	for {
		// for example this going to save  a2 
		var jugada string

		fmt.Print("Jugador " + player_name + option + " > ")
		fmt.Scanln(&jugada);


		if len(jugada) == 2 {

			// this is for separate the string
			listaJugada := strings.SplitAfter(jugada, "");
			numero, _ := strconv.Atoi(listaJugada[1]);

			ListaDeletrasPermitidas  := [8]  string{"a", "b", "c", "d", "e", "f", "g", "h"}
			ListaDeNumerosPermitidos := [8]  int   { 1, 2, 3, 4, 5, 6, 7, 8 }

			var keyListaletras bool
			var keyListanumeros bool

			// comapara la lista de letras permitidas con la jugada

			for indice := 0; indice < len ( ListaDeletrasPermitidas ) ; indice ++ {
				if ListaDeletrasPermitidas[indice] == listaJugada[0] {

					keyListaletras = true
				}

			}

			for indice := 0 ; indice < len(ListaDeNumerosPermitidos); indice++ {
				if ListaDeNumerosPermitidos[indice] == numero {
					keyListanumeros = true
				}

			}


			//compara los requisitos de la seleccion de la posicion 
			if keyListaletras && keyListanumeros {
				//  para  las letras en x
				Change_letters_numbers := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
				// para los numeros  en la lista  y
				Change_numbers_numbers := map[string]int{"1": 7, "2": 6, "3": 5, "4": 4, "5": 3, "6": 2, "7": 1, "8": 0}

				values[0] = Change_numbers_numbers[listaJugada[1]];
				values[1] = Change_letters_numbers[listaJugada[0]];

				break
			}
		}

	}
	return values[0], values[1];

}




