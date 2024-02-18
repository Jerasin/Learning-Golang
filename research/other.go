package research

import "fmt"

type Test struct {
	Name string
}

func (s *Test) model(id int) {
	fmt.Println(&s)
}

func model1(id int) {
	fmt.Println(id)
}

func Init() {
	// model1(2)
	test := Test{Name: "gg"}
	test.model(4)

}
