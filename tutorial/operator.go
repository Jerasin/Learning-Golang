package tutorial

import "fmt"

func Operator() {
	var number1 int = 30
	var number2 int = 45
	fmt.Println("ผลบวก = ", number1+number2)
	fmt.Println("ผลลบ = ", number1-number2)
	fmt.Println("ผลคูณ = ", number1*number2)
	fmt.Println("ผลหาร = ", number1/number2)
	fmt.Println("เศษ = ", number1%number2)
	fmt.Println("number1 == number2 = ", number1 == number2)
	fmt.Println("number1 >= number2 = ", number1 >= number2)
	fmt.Println("number1 <= number2 = ", number1 <= number2)
}
