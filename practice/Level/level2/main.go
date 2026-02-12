package main

import (
	"fmt"
)

func main() {
	fmt.Println("Чётное число от 1 до 20:")

	for i := 1; i <= 20; i++ {

		if i%2 == 0 {

			fmt.Print(i, "-чётное")
		}
	}

	fmt.Println()
	fmt.Println("\nОстаток при делении на 5:")

	for i := 1; i <= 20; i++ {

		remainder := i % 3

		fmt.Printf("%d %% 5 = %d\n", i, remainder)
	}
}
