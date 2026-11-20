package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t *Task) Summary() string {
	// TODO: верните "#ID Название [статус], приоритет N".
	return fmt.Sprintf("ID: %d, Title: [%s], Status: %s, Priority: %d", t.ID, t.Title, t.Status, t.Priority)
}
func (t *Task) Complete() {
	// TODO: присвойте статус done исходной задаче.
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	// TODO: верните &tasks[i] при совпадении, иначе nil.
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

// возвращает задачи с нужным статусом не редактируя исходный срез
func tasksByStatus(tasks []Task, status string) []Task {
	new_res := []Task{}
	for i := range tasks {
		if tasks[i].Status == status {
			new_res = append(new_res, tasks[i])
		}
	}
	if len(new_res) > 0 {
		return new_res
	} else {
		return nil
	}
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	// TODO: найдите ID=2, завершите и установите приоритет 3.
	// TODO: посчитайте завершённые задачи.

	// Output before: ID: 2, Title: [Освоить функции], Status: in_progress, Priority: 2
	target := findTaskByID(tasks, 2)
	if target != nil {
		target.Complete()
		target.SetPriority(3)
	}
	// Output after: ID: 2, Title: [Освоить функции], Status: done, Priority: 3

	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	done_k := 0
	for i := range tasks {
		if tasks[i].Status == "done" {
			done_k++
		}
	}
	fmt.Printf("Выполненных задач: %d \n", done_k)

	found := findTaskByID(tasks, 99)
	if found == nil {
		fmt.Println("Задача не найдена")
	}
	// Output: Задача не найдена

	by_status := tasksByStatus(tasks, "done")
	if by_status == nil {
		fmt.Println("Нету задачи с таким статусом")
	} else {
		for i := range by_status {
			fmt.Println(by_status[i].Summary())
		}
	}
	// Output: ID: 1, Title: [Установить Go], Status: done, Priority: 1
	//		   ID: 2, Title: [Освоить функции], Status: done, Priority: 3
}
