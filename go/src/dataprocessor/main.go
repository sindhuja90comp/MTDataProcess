package main

import (
    "fmt"
    "os"
    "strconv"
    "sync"
    "time"

    "dataprocessor" // Import the current package using module path
)

// Constants for number of workers and tasks
const numWorkers = 5
const numTasks = 20

func main() {
    fmt.Println("Running main in dataprocessor")
    // Initialize shared queue and result saver
    taskQueue := NewSharedQueue() // Assuming NewSharedQueue is defined in sharedqueue.go
    resultSaver, err := NewResultSaver("results_go.txt") // Assuming NewResultSaver is in resultsaver.go
    if err != nil {
        fmt.Println("Error creating ResultSaver:", err)
        os.Exit(1)
    }
    defer resultSaver.Close()

    // Create a wait group to wait for all workers to finish
    var wg sync.WaitGroup
    wg.Add(numWorkers)

    // Start worker goroutines
    for i := 1; i <= numWorkers; i++ {
        go worker(i, taskQueue, resultSaver, &wg)
    }

    // Add tasks to the queue
    for i := 0; i < numTasks; i++ {
        taskQueue.AddTask(Task{ID: strconv.Itoa(i), Data: fmt.Sprintf("Data-%d", i)}) // Assuming Task is in task.go
    }

    // Close the queue to signal no more tasks
    taskQueue.Close() // Assuming a Close method on SharedQueue

    // Wait for all workers to finish
    wg.Wait()

    fmt.Println("All tasks submitted and workers finished.")
}

// worker function to process tasks from the queue
func worker(workerId int, taskQueue *SharedQueue, resultSaver *ResultSaver, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d started.\n", workerId)

    for task := range taskQueue.Queue { // Assuming Queue is an exported field in SharedQueue
        result := task.Process() // Assuming Process is a method on Task
        err := resultSaver.SaveResult(result)
        if err != nil {
            fmt.Println("Error saving result:", err)
            break
        }
        fmt.Printf("Worker %d processed task %s\n", workerId, task.ID)
        time.Sleep(100 * time.Millisecond)
    }
    fmt.Printf("Worker %d finished.\n", workerId)
}