package main

import (
	"fmt"
	"sync"
	"time"
	"strconv"
	"os"
	"MTDataProcess/go/dataprocessor" // CRITICAL:  Import with your module path
)

// Constants for number of workers and tasks
const numWorkers = 5
const numTasks = 20

func main() {
	// Initialize shared queue and result saver
	taskQueue := dataprocessor.NewSharedQueue()
	resultSaver, err := dataprocessor.NewResultSaver("results_go.txt") // Corrected filename
	if err != nil {
		fmt.Println("Error creating ResultSaver:", err)
		os.Exit(1) // Use os.Exit for consistent error handling
	}
	defer resultSaver.Close() // Ensure resultSaver.Close() is called

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, taskQueue, resultSaver, &wg)
	}

	// Add tasks to the queue
	for i := 0; i < numTasks; i++ {
		taskQueue.AddTask(dataprocessor.Task{ID: strconv.Itoa(i), Data: fmt.Sprintf("Data-%d", i)})
	}

	//close the channel, no more tasks are coming
	close(taskQueue.queue)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All tasks submitted and workers finished.")
}

// worker function to process tasks from the queue
func worker(workerId int, taskQueue *dataprocessor.SharedQueue, resultSaver *dataprocessor.ResultSaver, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group counter when the worker finishes
	fmt.Printf("Worker %d started.\n", workerId)

	for task := range taskQueue.queue { //changed from for true to for range
		//removed the check for empty queue, the loop will exit automatically
		// when the channel is closed
		result := task.Process()
		err := resultSaver.SaveResult(result)
		if err != nil {
			fmt.Println("Error saving result:", err)
			// Consider what to do here:  Retry?  Abort?
			break // IMPORTANT: Exit the worker loop on error
		}
		fmt.Printf("Worker %d processed task %s\n", workerId, task.ID)
		time.Sleep(100 * time.Millisecond) // Simulate processing time
	}
	fmt.Printf("Worker %d finished.\n", workerId)
}

