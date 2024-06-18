package main

import (
	"fmt"

	"sync"
)

const tasksCount int = 15
const workersCount int = 5

type Task struct {
	taskName string
	IsDone   bool
}

func main() {
	tasks := generateTasks(tasksCount)

	wg := sync.WaitGroup{}

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go doTask(tasks, &wg)
	}

	wg.Wait()
}

func doTask(taskQueue <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskQueue {
		task.IsDone = true
		fmt.Printf("Task %s is done\n", task.taskName)
	}
}

func generateTasks(n int) <-chan Task {
	tasks := make(chan Task, n)

	for i := 0; i < n; i++ {
		task := Task{fmt.Sprintf("task%d", i), false}
		tasks <- task
	}

	close(tasks)
	return tasks
}
