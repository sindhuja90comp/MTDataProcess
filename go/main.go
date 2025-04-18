package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Task represents a unit of work.
type Task struct {
	ID   int
	Data string
}

// Process simulates some computational work.
func (t *Task) Process() string {
	delay := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(delay)
	return fmt.Sprintf("Processed Task %d: %s", t.ID, t.Data)
}

// SharedQueue is a concurrent-safe queue for tasks.
type SharedQueue struct {
	queue chan Task
	mu    sync.Mutex
}

// NewSharedQueue creates a new shared queue.
func NewSharedQueue(capacity int) *SharedQueue {
	return &SharedQueue{
		queue: make(chan Task, capacity),
	}
}

// AddTask adds a task to the queue.
func (sq *SharedQueue) AddTask(task Task) {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	sq.queue <- task
}

// GetTask retrieves a task from the queue. Returns the task and a boolean indicating if the queue is not empty.
func (sq *SharedQueue) GetTask() (Task, bool) {
	task, ok := <-sq.queue
	return task, ok
}

// Close signals that no more tasks will be added.
func (sq *SharedQueue) Close() {
	close(sq.queue)
}

// ResultSaver is responsible for saving processing results.
type ResultSaver struct {
	file *os.File
	mu   sync.Mutex
}

// NewResultSaver creates a new ResultSaver.
func NewResultSaver(filename string) (*ResultSaver, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create result file: %w", err)
	}
	return &ResultSaver{file: file}, nil
}

// SaveResult saves a processing result to the file.
func (rs *ResultSaver) SaveResult(result string) error {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	_, err := rs.file.WriteString(result + "\n")
	if err != nil {
		return fmt.Errorf("failed to write result: %w", err)
	}
	return nil
}

// Close closes the result file.
func (rs *ResultSaver) Close() error {
	if rs.file != nil {
		return rs.file.Close()
	}
	return nil
}

func worker(id int, queue *SharedQueue, saver *ResultSaver, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Worker %d started", id)
	for {
		task, ok := queue.GetTask()
		if !ok {
			log.Printf("Worker %d finished (no more tasks)", id)
			return
		}
		log.Printf("Worker %d retrieved task %d", id, task.ID)
		result := task.Process()
		err := saver.SaveResult(result)
		if err != nil {
			log.Printf("Worker %d encountered error saving result: %v", id, err)
			break
		}
		log.Printf("Worker %d saved result for task %d", id, task.ID)
	}
}

func main() {
	numWorkers := 5
	numTasks := 20
	queueCapacity := 10

	log.Println("Starting Data Processing System (Go)")

	// Initialize shared queue
	taskQueue := NewSharedQueue(queueCapacity)

	// Initialize result saver
	resultSaver, err := NewResultSaver("go_results.txt")
	if err != nil {
		log.Fatalf("Error creating result saver: %v", err)
	}
	defer resultSaver.Close()

	// Create a WaitGroup to wait for all workers to complete
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, taskQueue, resultSaver, &wg)
	}

	// Add tasks to the queue
	for i := 1; i <= numTasks; i++ {
		task := Task{ID: i, Data: fmt.Sprintf("Data for Task %d", i)}
		taskQueue.AddTask(task)
		log.Printf("Submitted task %d to the queue", i)
		time.Sleep(50 * time.Millisecond) // Simulate task submission rate
	}

	// Close the queue to signal no more tasks
	taskQueue.Close()
	log.Println("Finished submitting tasks")

	// Wait for all workers to finish
	wg.Wait()
	log.Println("All workers finished. Results saved to go_results.txt")
}