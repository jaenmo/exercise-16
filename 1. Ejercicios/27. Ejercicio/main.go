package main

import "fmt"

func main(){
i:=100
for {
	// imprime los números excepto 0
	fmt.Printf("Value of i is %v\n", i)
	i-=10
	if i == 0{
		break
	}

	// imprime los números incluido 0
	fmt.Printf("Value of i is %v\n", i)
	if i == 0{
		break
	}
	i-=10
}
fmt.Println("finished loop")

}