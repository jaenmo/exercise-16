package main

import ("fmt"
	"math/rand"
)


func main(){

	for i:=0 ; i<10 ; i++{
		x:= rand.Intn(5)
		//fmt.Printf("loop number %v\t", i)
	switch{
	case x == 0:
		fmt.Printf("es la %v vez, x is %v\n", i, x)
	case x == 1:
		fmt.Printf("es la %v vez, x is %v\n", i, x)
	case x == 2:
		fmt.Printf("es la %v vez, x is %v\n", i, x)
	case x == 3:
		fmt.Printf("es la %v vez, x is %v\n", i, x)
	case x == 4:
		fmt.Printf("es la %v vez, x is %v\n", i, x)

	}
}

}