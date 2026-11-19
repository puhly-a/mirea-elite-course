package main

import "fmt"

func main() {
	var t1, t2, t3 float64 = 1, 1, 2
	sum := t1 + t2 + t3
	sr := (t1 + t2 + t3) / 3
	fmt.Println("Сумма:", sum)
	fmt.Printf("Среднее значение: %.2f \n", sr)
}
