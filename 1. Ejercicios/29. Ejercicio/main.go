package main

import "fmt"

func main(){
	for outer := 0 ; outer <5 ; outer ++{
		for inner :=5 ;inner >0 ; inner --{
			fmt.Printf("%v outer loop\t %v inner loop\n", outer, inner)
		}
		fmt.Println("")
	}
}