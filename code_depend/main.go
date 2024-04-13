package main

import (
	"fmt"

	"github.com/jaenmo/Ejercicios/gopherguides/greet"
	"github.com/jaenmo/puppy/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2:= puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)


	greet.Hello()

}
