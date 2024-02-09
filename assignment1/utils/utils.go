package utils

import "fmt"

func EvenOrOdd(num int) {
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func OneToFive() {
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}
}
