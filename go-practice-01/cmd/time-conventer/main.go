package main

import "fmt"

func main() {
	var time int
	const limit = 60
	fmt.Println("Введите кол-во минут:")
	_, err := fmt.Scan(&time)
	if err != nil {
		fmt.Println("Ошибка! Введите целое число.")
		return
	}
	if time < 0 {
		fmt.Println("Ошибка! Кол-во. минут не может быть < 0.")
		return
	}
	//fmt.Println(time)
	fmt.Println("Полных часов:", time/60)
	fmt.Println("Оставшихся минут:", time%60)
}
