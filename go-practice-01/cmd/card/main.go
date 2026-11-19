package main

import "fmt"

func main() {
	var name string = "Виталий"
	var group string = "ИКБО-01"
	count := 1
	const course = "Бэкэнд разработка на Golang"
	var status bool = true

	fmt.Printf("Участник - %s \n", name)
	fmt.Printf("Группа - %s\n", group)
	fmt.Println("Курс -", course)
	fmt.Printf("Колличество работ - %d \n", count)
	fmt.Printf("Интерес - %t \n", status)

}
