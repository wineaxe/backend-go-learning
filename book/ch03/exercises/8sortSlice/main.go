package main

import (
	"fmt"
	"sort" // пакет для сортировки
)

// aStructure — структура (тип данных с несколькими полями)
type aStructure struct {
	person string // имя человека

	height int // рост в см

	weight int // вес в кг
}

func main() {

	// Создаём пустой срез структур: make([]aStructure, 0) — длина 0, но не nil
	mySlice := make([]aStructure, 0)

	// Добавляем элементы в срез через append
	// Каждый элемент — структура aStructure с тремя полями
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})

	mySlice = append(mySlice, aStructure{"Bill", 134, 45})

	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})

	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})

	mySlice = append(mySlice, aStructure{"Athina", 134, 40})

	fmt.Println("0:", mySlice) // вывод до сортировки (в порядке добавления)

	// sort.Slice сортирует срез по заданному критерию
	// Второй параметр — функция сравнения: func(i, j int) bool
	// Функция должна вернуть true, если элемент i должен идти ПЕРЕД элементом j
	sort.Slice(mySlice, func(i, j int) bool {

		// Сортировка по возрастанию роста: если рост i меньше роста j, i идёт первым
		return mySlice[i].height < mySlice[j].height

	})

	fmt.Println("<:", mySlice) // вывод после сортировки по возрастанию роста
	// Результат: от меньшего роста к большему
	// Bill (134), Athina (134), Epifanios (144), Marietta (155), Mihalis (180)

	// Сортировка по убыванию роста
	sort.Slice(mySlice, func(i, j int) bool {

		// Если рост i больше роста j, i идёт первым → сортировка по убыванию
		return mySlice[i].height > mySlice[j].height

	})

	fmt.Println(">:", mySlice) // вывод после сортировки по убыванию роста
	// Результат: от большего роста к меньшему
	// Mihalis (180), Marietta (155), Epifanios (144), Bill (134), Athina (134)
}
