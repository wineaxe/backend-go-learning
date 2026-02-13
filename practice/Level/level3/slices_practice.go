package main

import (
	"fmt"
	"strings"
)

func slices_practice() {
	fmt.Println("=== УРОВЕНЬ 2: СРЕЗЫ (SLICES) ===\n")

	// ЗАДАНИЕ 2.1: Создание и манипуляция срезами
	fmt.Println("ЗАДАНИЕ 2.1: len и cap при append")
	fmt.Println(strings.Repeat("─", 50))

	// 1. Создаём срез
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Исходный срез:", numbers)
	fmt.Printf("len = %d, cap = %d\n", len(numbers), cap(numbers))

	// 2. Добавляем элемент
	numbers = append(numbers, 60)
	fmt.Println("\nПосле append(60):")
	fmt.Println("Срез:", numbers)
	fmt.Printf("len = %d, cap = %d\n", len(numbers), cap(numbers))

	// Объяснение:
	fmt.Println("\n📝 ОБЪЯСНЕНИЕ:")
	fmt.Println("- len увеличился на 1 (добавили один элемент)")
	fmt.Println("- cap увеличился (Go выделил больше памяти)")
	fmt.Println("- Go обычно удваивает cap, когда len == cap\n")

	// ЗАДАНИЕ 2.2: Срезы срезов (reslicing)
	fmt.Println("ЗАДАНИЕ 2.2: Срезы срезов (reslicing)")
	fmt.Println(strings.Repeat("─", 50))

	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Исходный срез:", original)

	// Создаём разные срезы
	a := original[2:5] // элементы с индекса 2 до 4 (5 не включается!)
	b := original[:4]  // элементы с начала до индекса 3
	c := original[5:]  // элементы с индекса 5 до конца
	d := original[:]   // все элементы

	fmt.Println("\na := original[2:5]  →", a, fmt.Sprintf("(len=%d, cap=%d)", len(a), cap(a)))
	fmt.Println("b := original[:4]   →", b, fmt.Sprintf("(len=%d, cap=%d)", len(b), cap(b)))
	fmt.Println("c := original[5:]   →", c, fmt.Sprintf("(len=%d, cap=%d)", len(c), cap(c)))
	fmt.Println("d := original[:]    →", d, fmt.Sprintf("(len=%d, cap=%d)", len(d), cap(d)))

	// Модифицируем элемент в a
	fmt.Println("\nМодифицируем a[0] = 999:")
	a[0] = 999
	fmt.Println("a =", a)
	fmt.Println("original =", original, "← ИЗМЕНИЛСЯ!")

	fmt.Println("\n📝 ОБЪЯСНЕНИЕ:")
	fmt.Println("- Срез 'смотрит' на ту же память, что и исходный")
	fmt.Println("- Изменение в срезе влияет на исходный массив/срез")
	fmt.Println("- Это НЕ копия, а 'вид' на ту же память\n")

	// ЗАДАНИЕ 2.3: Copy слайса
	fmt.Println("ЗАДАНИЕ 2.3: Copy слайса")
	fmt.Println(strings.Repeat("─", 50))

	source := []int{1, 2, 3}
	dest := []int{0, 0, 0, 0, 0}

	fmt.Println("До copy:")
	fmt.Println("source =", source)
	fmt.Println("dest   =", dest)

	copy(dest, source) // копируем из source в dest

	fmt.Println("\nПосле copy(dest, source):")
	fmt.Println("source =", source)
	fmt.Println("dest   =", dest)

	// Изменяем source
	source[0] = 100
	fmt.Println("\nПосле source[0] = 100:")
	fmt.Println("source =", source)
	fmt.Println("dest   =", dest, "← НЕ ИЗМЕНИЛСЯ!")

	fmt.Println("\n📝 ОБЪЯСНЕНИЕ:")
	fmt.Println("- copy() создаёт КОПИЮ элементов")
	fmt.Println("- Изменение source НЕ влияет на dest")
	fmt.Println("- copy() копирует min(len(dest), len(source)) элементов\n")

	// ЗАДАНИЕ 2.4: Append и изменение cap
	fmt.Println("ЗАДАНИЕ 2.4: Отслеживание изменения cap")
	fmt.Println(strings.Repeat("─", 50))

	s := make([]int, 0, 2) // len=0, cap=2
	fmt.Printf("Начальное состояние: len=%d, cap=%d\n", len(s), cap(s))

	for i := 1; i <= 5; i++ {
		s = append(s, i)
		fmt.Printf("После append(%d): len=%d, cap=%d, срез=%v\n", i, len(s), cap(s), s)
	}

	fmt.Println("\n📝 ОБЪЯСНЕНИЕ:")
	fmt.Println("- Когда len == cap, Go выделяет новую память")
	fmt.Println("- Обычно cap удваивается (но не всегда)")
	fmt.Println("- Переаллокация происходит автоматически\n")

	// ДОПОЛНИТЕЛЬНЫЕ ПРИМЕРЫ
	fmt.Println("ДОПОЛНИТЕЛЬНЫЕ ПРИМЕРЫ:")
	fmt.Println(strings.Repeat("─", 50))

	// Пример: Разница между nil и пустым срезом
	fmt.Println("\n1. nil срез vs пустой срез:")
	var nilSlice []int
	emptySlice := []int{}
	makeSlice := make([]int, 0)

	fmt.Printf("nilSlice: %v, len=%d, cap=%d, nil? %v\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptySlice: %v, len=%d, cap=%d, nil? %v\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("makeSlice: %v, len=%d, cap=%d, nil? %v\n", makeSlice, len(makeSlice), cap(makeSlice), makeSlice == nil)

	// Пример: Append к nil срезу
	fmt.Println("\n2. Append к nil срезу работает!")
	nilSlice = append(nilSlice, 1, 2, 3)
	fmt.Println("nilSlice после append:", nilSlice)

	// Пример: Трёхпараметровый срез
	fmt.Println("\n3. Трёхпараметровый срез [low:high:max]:")
	full := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	limited := full[2:5:7] // элементы 2-4, cap ограничен до 7
	fmt.Printf("full = %v\n", full)
	fmt.Printf("limited = full[2:5:7] = %v (len=%d, cap=%d)\n", limited, len(limited), cap(limited))
	fmt.Println("Ограничение cap предотвращает изменение элементов за пределами среза")

	fmt.Println("\n✅ УРОВЕНЬ 2 ЗАВЕРШЁН!")
	fmt.Println("Переходите к Уровню 3: Карты (Maps)")
}
