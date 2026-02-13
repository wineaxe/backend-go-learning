package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Repeat("-", 50))

	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("Исходный срез:", numbers)
	fmt.Printf("len = %d, cap = %d\n", len(numbers), cap(numbers))

	numbers = append(numbers, 60)

	fmt.Println("\nПосле append(60):")

	fmt.Println("Срез:", numbers)

	fmt.Printf("len = %d, cap = %d\n", len(numbers), cap(numbers))

	fmt.Println("\n📝 Обьяснение:")

	fmt.Println("- len увеличился на 1 (добавили один элемент)")

	fmt.Println("- cap увеличился (Go выделил больше памяти)")

	fmt.Println("- Go обычно удваивает cap, когда len == cap\n")
}
