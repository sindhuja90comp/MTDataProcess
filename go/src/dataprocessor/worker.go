package dataprocessor

import (
	"fmt"
	"sync"
)

func worker(id int, queue *SharedQueue, saver *ResultSaver, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started.\n", id)
	for {
		task, ok := queue.GetTask()
		if !ok {
			fmt.Printf("Worker %d finished (no more tasks).\n", id)
			return
		}
		result := task.Process()
		saver.SaveResult(result)
		fmt.Printf("Worker %d processed task %s\n", id, task.ID)
	}
}