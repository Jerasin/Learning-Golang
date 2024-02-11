package tutorial

import "fmt"

func Input() {
	var name string
	var score int
	fmt.Print("ป้อนชื่อนักเรียน = ")
	fmt.Scanf("%s", &name)
	fmt.Println("สวัสดี = ", name)

	fmt.Println("ป้อนคะแนน = ")
	fmt.Scanf("%d", &score)
	fmt.Println("ได้คะแนน = ", score)
}
