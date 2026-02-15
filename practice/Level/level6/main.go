package main

import (
	"fmt"
	"strings"
)

func main() {
	// ЗАДАНИЕ 2.4: Append и изменение cap
	fmt.Println("ЗАДАНИЕ 2.4: Отслеживание изменения cap")
	fmt.Println(strings.Repeat("─", 50))

	s := make([]int, 0, 2) //len 0, cap 2

	fmt.Printf("Начальное состояние: len=%d, cap=%d\n", len(s), cap(s))

	for i := 1; i <= 5; i++ {

		s = append(s, i)
		fmt.Printf("После append(%d): len=%d, cap=%d, срез=%d\n", i, len(s), cap(s), s)

	}

}
