package tutorial

import "fmt"

func loop() {
	for i := 0; i < 10; i++ {

		if i == 5 {
			fmt.Println("continue", i)
			continue
		} else if i == 9 {
			fmt.Println("break", i)
			break
		}

		fmt.Println("index", i)
	}
}

func loopRange() {
	scores := [...]string{"jame", "jin", "mo", "jak"}

	for i, v := range scores {
		fmt.Println("v in scores", i, v)
	}
}

func Loop() {
	// loop()
	loopRange()
}
