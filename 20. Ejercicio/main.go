package main

import (
	"fmt"
	"math/rand"
)

func init(){
	fmt.Println("this is just the init")
}

func main(){
	x := rand.Intn(400)
	fmt.Println(x)

	switch {
	
	case x <= 100:
		fmt.Println("between 0 and 100")
	
	case x > 100 && x <= 200: 
		fmt.Println("between 101 and 200")
	
	case x > 200 && x <= 250: 
		fmt.Println("between 201 and 250")
	
default:
	fmt.Println("it is more than 250")

}
	
}