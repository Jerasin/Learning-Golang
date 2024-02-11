package tutorial

import "fmt"

func dataTypeConst() {
	const name = "Jame"
	fmt.Println("Data Type Const name", name)
}

func stringType() {
	// fmt.Println("Data Type String")
	var name1 string
	fmt.Println("Data Type String name1", name1)
	name1 = "gg"
	name2 := "jin"
	fmt.Println("Data Type String name1", name1)
	fmt.Println("Data Type String name2", name2)
}

func intType() {
	// fmt.Println("Data Type Int")
	var int1 int
	fmt.Println("Data Type Int int1", int1)
	int1 = 4
	int2 := 10
	fmt.Println("Data Type Int int1", int1)
	fmt.Println("Data Type Int int2", int2)

	score3, score4, score5 := 2, 3, 4
	fmt.Println("Data Type Int score3", score3)
	fmt.Println("Data Type Int score4", score4)
	fmt.Println("Data Type Int score5", score5)
}

func floatType() {
	var score1 float64
	fmt.Println("Data Type Float score1", score1)
	score1 = 100.30
	score2 := 99.99
	fmt.Println("Data Type Float score1", score1)
	fmt.Println("Data Type Float score2", score2)
}

func boolType() {
	var isPass1 bool
	isPass2 := true
	isPass1 = false
	fmt.Println("Data Type Bool isPass1", isPass1)
	fmt.Println("Data Type Bool isPass2", isPass2)
}

func arrayType() {
	var numbers1 [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(numbers1)

	numbers2 := [4]int{5, 6, 7, 8}
	fmt.Println(numbers2)

	names1 := [2]string{"5", "6"}
	fmt.Println(names1)
	fmt.Printf("%T\n", names1)
	y := len(names1)
	fmt.Println(y)

	var numbers3 [3]int
	fmt.Println(numbers3)
	numbers3[0] = 4
	fmt.Println(numbers3)

	var numbers4 = [...]int{1, 3, 4, 5}
	fmt.Println(numbers4)
	numbers4[3] = 66
	fmt.Println(numbers4)
}

func slicesType() {
	names := []string{"jame", "jin"}
	fmt.Println(names)
	names = append(names, "fff")
	fmt.Println(names)
	fmt.Println("----------------------------------------------------------")
	fmt.Println(names[0])
	fmt.Println(names[:])  // เอาทุก index
	fmt.Println(names[1:]) // index ที่ 1 ถึงสุดท้าย
	fmt.Println(names[:2]) // index ที่ 0 ถึง 1
}

func mapType() {
	country := map[string]string{}
	country["TH"] = "Thailand"
	country["EN"] = "England"
	fmt.Println(country)

	fmt.Println("----------------------------------------------------------")
	coin := map[string]int{"TH": 87}
	fmt.Println(coin)

	fmt.Println("----------------------------------------------------------")
	value1, check1 := country["TH"]
	fmt.Println("value", value1)
	fmt.Println("check", check1)

	fmt.Println("----------------------------------------------------------")
	value, check := coin["GG"]
	fmt.Println("value", value)
	fmt.Println("check", check)
}

func DataType() {
	// stringType()
	// intType()
	// floatType()
	// boolType()
	// dataTypeConst()
	// arrayType()
	// slicesType()
	mapType()
}
