package main

import (
	"fmt"
	"math/rand"
)


func main(){
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Println(x)
	fmt.Println(y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are greater than 6")
	case x >= 4 && x >= 6 && y !=5:
		fmt.Println("x is greater than or equal to 4 and less or equal to 6 and y is not 5")
	default:
		fmt.Println("nothing of the previous cases were met")
	
}

}