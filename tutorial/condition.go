package tutorial

import "fmt"

func ifElse() {
	score := 10

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 50 && score < 90 {
		fmt.Println("B")
	} else {
		fmt.Println("F")
	}

	if score%2 == 0 {
		fmt.Println("เลขคู่")
	} else {
		fmt.Println("เลขคี่")
	}
}

func switchCase() {
	name := "gg"
	switch name {
	case "gg":
		fmt.Println("gg")
	case "aa":
		fmt.Println("aa")
	}

	score := 90
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 50:
		fmt.Println("B")
	default:
		fmt.Println("F")
	}
}

func Condition() {
	ifElse()
	switchCase()
}
