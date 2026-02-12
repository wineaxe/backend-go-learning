package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 3

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Целочисленное деление (результат int)
	intDiv := a / b
	fmt.Println("Целочисленное деление:", intDiv) // 3

	// Деление с плавающей точкой (результат float64)
	floatDiv := float64(a) / float64(b)
	fmt.Println("Деление с float64:", floatDiv) // 3.333...

	// Объяснение разницы:
	// - int / int = int (дробная часть отбрасывается)
	// - float64 / float64 = float64 (сохраняет дробную часть)
}
