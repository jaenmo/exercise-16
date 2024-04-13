package main

import (
	"fmt"
	"math/rand"
)

func main(){
	x := rand.Intn(400)
	fmt.Println(x)

	if x <= 100 {
		fmt.Println("between 0 and 100")
	}else if x > 100 && x <= 200 {
		fmt.Println("between 101 and 200")
	}else if x > 200 && x <= 250 {
		fmt.Println("between 201 and 250")
	}else{
		fmt.Println("that was more than 250")

	}

	
	// y := rand.Intn(4)
	// fmt.Println(y)


}