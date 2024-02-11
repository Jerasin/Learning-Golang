package tutorial

import "fmt"

func sumTotal1(number1 int, number2 int) (int, string) {
	total := number1 + number2
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}

	// fmt.Println("total", total)
	return total, status
}

func sumTotal2(number1, number2 int) (int, string) {
	total := number1 + number2
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}

	// fmt.Println("total", total)
	return total, status
}

func variadicFunction(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total = total + v
	}

	return total
}

func Function() {
	// a, b := sumTotal2(5, 9)
	// fmt.Println("a", a)
	// fmt.Println("b", b)

	// x, y := sumTotal1(5, 9)
	// fmt.Println("x", x)
	// fmt.Println("y", y)

	z := variadicFunction(90, 80, 50)
	fmt.Println("z", z)

}
