package tutorial

import "fmt"

func Print() {
	name := "Mo"
	age := 27
	score := 99.99
	isPass := false
	fmt.Println("My name is", name)
	fmt.Printf("My age %v\n", age)
	fmt.Printf("Type age %T\n", age)
	fmt.Printf("Type score %T\n", score)
	fmt.Printf("Type isPass %T\n", isPass)
}
