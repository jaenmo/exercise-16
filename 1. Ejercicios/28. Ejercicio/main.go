package main

// ejercicio para usar modulus (%) y continue en un loop para núm impares

/*imprimir números impares múltiplos de 3 decreciendo desde 100
si el número divisible por 3 es par, imprimir, pero destacar
*/

import "fmt"

func main() {
	for x := 50; x > 0; x-- {
		if x%3 == 0 && x%2 != 0 {
			fmt.Printf("%v es divisible por 3 y es impar\n", x)
		}
		if x%3 == 0 && x%2 == 0 {
			fmt.Printf("%v es divisible por 3 y es par\n", x)

		}

	}

	// x := 50
	// for x >= 0 {
	// 	if x%3 == 0 && x%2 !=0{
	// 		fmt.Printf("%v es divisible por 3 y es impar\n", x)
	// 		x--
	// 		// if x%2 == 0 {
	// 		// 	fmt.Printf("%v además, es par\n", x)
	// 		// 	continue
	// 		// }

	// 	}
	// }

}

//fmt.Println("fin del proceso")
