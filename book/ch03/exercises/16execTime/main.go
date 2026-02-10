package main

import (
	"fmt"
	"time"
)

func main() {

	// ПРИМЕР 1: Замер времени выполнения time.Sleep(1 секунда)
	// time.Now() возвращает текущий момент времени (тип time.Time)
	start := time.Now()

	// Приостанавливаем выполнение на 1 секунду
	time.Sleep(time.Second)

	// time.Since(start) возвращает длительность, прошедшую с момента start (тип time.Duration)
	duration := time.Since(start)

	fmt.Println("It took time.Sleep(1)", duration, "to finish.")
	// Вывод: примерно "It took time.Sleep(1) 1.000123456s to finish." (чуть больше 1 сек из-за накладных расходов)

	// ПРИМЕР 2: Замер времени выполнения time.Sleep(2 секунды)
	start = time.Now()

	// 2 * time.Second — пауза на 2 секунды
	time.Sleep(2 * time.Second)

	duration = time.Since(start)

	fmt.Println("It took time.Sleep(2)", duration, "to finish.")
	// Вывод: примерно "It took time.Sleep(2) 2.000xxx s to finish."

	// ПРИМЕР 3: Замер времени пустого цикла (200 миллионов итераций)
	start = time.Now()

	for i := 0; i < 200000000; i++ {

		_ = i // пустая итерация (компилятор может её оптимизировать)
	}

	duration = time.Since(start)

	fmt.Println("It took the for loop", duration, "to finish.")
	// Вывод зависит от процессора, например: "It took the for loop 45.234ms to finish."

	// ПРИМЕР 4: Замер времени цикла с реальной работой (суммирование)
	sum := 0

	start = time.Now()

	for i := 0; i < 200000000; i++ {

		sum += i // на каждой итерации добавляем i к sum
	}

	duration = time.Since(start)

	fmt.Println("It took the for loop", duration, "to finish.")
	// Этот цикл обычно выполняется дольше, чем пустой, так как есть обращение к памяти (sum)
}
