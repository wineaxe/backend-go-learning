package main

import (
	"fmt"
	"strings"
)

func main() {
	//ЗАДАНИЕ 2.2: Срезы срезов (reslicing)
	fmt.Println("ЗАДАНИЕ 2.2: Срезы срезов (reslicing)")
	fmt.Println(strings.Repeat("-", 50))

	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Исходный срез:", original)

	//Создаём разные срезы
	a := original[2:5] //элементы с индекса 2 до  4 (5 не включается!)
	b := original[:4]  //элементы с начала до индекса 3
	c := original[5:]  //элементы с индекса 5 до конца
	d := original[:]   //всё элементы

	fmt.Println("\na := original[2:5]   ->", a, fmt.Sprintf("len=%d, cap=%d", len(a), cap(a)))
	fmt.Println("b := original[:4]   ->", b, fmt.Sprintf("(len=%d, cap=%d)", len(b), cap(b)))
	fmt.Println("c := original[5:]   ->", c, fmt.Sprintf("(len=%d, cap=%d)", len(c), cap(c)))
	fmt.Println("d := original[:]   ->", d, fmt.Sprintf("(len=%d, cap=%d)", len(d), cap(d)))

	fmt.Println("\nМодифицируем a[] = 999:")
	a[0] = 999
	fmt.Println("a =", a)
	fmt.Println("original =", original, "<- ИЗМЕНИЛСЯ!", fmt.Sprintf("len=%d, cap=%d", len(d), cap(d)))

	fmt.Println("\n📝 Обьяснение:")
	fmt.Println("- Срез 'смотрит' на ту же память, что и исходный")
	fmt.Println("- Изменения в срезе влияет на исходный массив/срез")
	fmt.Println("- Это НЕ копия, а 'вид' на ту же память\n")
}
