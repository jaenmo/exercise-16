package main

import "fmt"

func main() {
m := map[string] int{
	"James":		42,
	"MoneyPenny":	32,
}
age1 := m["James"]

fmt.Println(age1)

fmt.Println("---------------")

for i, _ := range m{
	fmt.Println(i)
}

fmt.Println("---------------")

age2 := m["Q"]
fmt.Println(age2)


if k, ok := m["Q"]; ok{
	fmt.Printf("Q value is %v",k)

}
}