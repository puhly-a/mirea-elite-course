package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В процессе."
	case "done":
		return "Завершена."
	default:
		return "Неизвестный статус."
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	var count int = 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	return count
}

// findStatus ищет первый индекс заданного статуса
func findStatus(statuses []string, target string) int {
	for i, status := range statuses {
		if status == target {
			return i
		}
	}
	return -1
}

// completionPercent выводит процент выполнения от списка задач
func completionPercent(statuses []string) float64 {
	already_done := countStatus(statuses, "done")
	f64 := float64(already_done)
	if already_done != 0 {
		return (100 * f64 / float64(len(statuses)))
	} else {
		return 0
	}
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	// TODO: добавьте ещё один статус new через append.
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	// TODO: создайте map[string]int для подсчёта статусов.
	// TODO: выведите new, in_progress, done в фиксированном порядке.
	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}

	fmt.Printf("Колличество разных статусов: new - %d, in_progress - %d, done - %d \n", counts["new"], counts["in_progress"], counts["done"])

	fmt.Println("Index:", findStatus(statuses, "new"))
	fmt.Printf("Процент выполнения: %.2f%%\n", completionPercent(statuses))
}
