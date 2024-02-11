package tutorial

import "fmt"

type Product struct {
	name     string
	price    float64
	category string
	discount float64
}

func Struct() {
	product1 := Product{name: "car", price: 99.99, category: "item", discount: 10}
	product2 := Product{name: "pen", price: 50, category: "item", discount: 20}

	fmt.Println("product1", product1)
	fmt.Println("product2", product2)
}
