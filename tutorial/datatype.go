package tutorial

import "fmt"

func dataTypeConst() {
	const name = "Jame"
	fmt.Println("Data Type Const name", name)
}

func DataType() {
	// fmt.Println("Data Type String")
	var name1 string
	fmt.Println("Data Type String name1", name1)
	name1 = "gg"
	name2 := "jin"
	fmt.Println("Data Type String name1", name1)
	fmt.Println("Data Type String name2", name2)

	// fmt.Println("Data Type Int")
	var int1 int
	fmt.Println("Data Type Int int1", int1)
	int1 = 4
	int2 := 10
	fmt.Println("Data Type Int int1", int1)
	fmt.Println("Data Type Int int2", int2)

	var score1 float64
	fmt.Println("Data Type Float score1", score1)
	score1 = 100.30
	score2 := 99.99
	fmt.Println("Data Type Float score1", score1)
	fmt.Println("Data Type Float score2", score2)

	var isPass1 bool
	isPass2 := true
	isPass1 = false
	fmt.Println("Data Type Bool isPass1", isPass1)
	fmt.Println("Data Type Bool isPass2", isPass2)

	score3, score4, score5 := 2, 3, 4
	fmt.Println("Data Type Int score3", score3)
	fmt.Println("Data Type Int score4", score4)
	fmt.Println("Data Type Int score5", score5)

	dataTypeConst()
}
