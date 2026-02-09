package main

import (
	"fmt"
)

// printSlice выводит все элементы среза через пробел и переводит строку
func printSlice(x []int) {

	for _, number := range x {

		fmt.Print(number, " ")

	}
	fmt.Println()
}

func main() {

	// Срез (slice): динамический "вид" на массив. Изначально 3 элемента, len=3, cap=3
	aSlice := []int{-1, 0, 4}

	fmt.Printf("aSlice:")

	printSlice(aSlice)

	// cap — вместимость (сколько элементов можно добавить без переаллокации), len — текущая длина
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	// После создания: Cap: 3, Length: 3

	// append добавляет элемент в конец; если места нет (len == cap), выделяется новый массив и cap растёт
	aSlice = append(aSlice, -100)

	fmt.Printf("aSlice:")

	printSlice(aSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	// Обычно cap удваивается: было 3 → стало 6, len = 4

	aSlice = append(aSlice, -2)

	aSlice = append(aSlice, -3)

	aSlice = append(aSlice, -4)
	// Добавили ещё 3 элемента: len = 7. Когда len достиг 6, снова мог вырасти cap (например до 12)

	printSlice(aSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	// Итог: Length = 7, Cap зависит от реализации (часто 8 или 12)
}
