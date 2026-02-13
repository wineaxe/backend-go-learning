package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("ЗАДАНИЕ 2.3: Copy слайса")
	fmt.Println(strings.Repeat("-", 50))

	source := []int{1, 2, 3}

	dest := []int{0, 0, 0, 0, 0}

	fmt.Println("До copy:")

	fmt.Println("source =", source)

	fmt.Println("dest =", dest)

	copy(dest, source) // копируем из source в dest

	fmt.Println("\nПосле copy(dest, source):")

	fmt.Println("source =", source)

	fmt.Println("dest =", dest)

	source[0] = 100

	fmt.Println("\nПосле source[0] = 100")

	fmt.Println("source =", source)

	fmt.Println("dest =", dest, "<- НЕ ИЗМЕНИЛСЯ!")

	fmt.Println("\n📝 ОБЬЯСНЕНИЕ:")

	fmt.Println("- copy() создает КОПИЮ элементов")

	fmt.Println("len=%d, cap=%d", fmt.Sprintf("len=%d, cap=%d", len(source), cap(source)))
	fmt.Println("len=%d, cap=%d", fmt.Sprintf("len=%d, cap=%d", len(dest), cap(dest)))

	fmt.Println("- copy() создает КОПИЮ элементов")

}
