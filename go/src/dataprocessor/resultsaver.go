package dataprocessor

import (
	"fmt"
	"os"
)

// ResultSaver struct to hold file writer.
type ResultSaver struct {
	fileWriter *os.File
	filename   string
}

// NewResultSaver creates a new ResultSaver.
func NewResultSaver(filename string) (*ResultSaver, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &ResultSaver{fileWriter: file, filename: filename}, nil
}

// SaveResult saves a result to the file.
func (rs *ResultSaver) SaveResult(result string) error {
	_, err := rs.fileWriter.WriteString(result + "\n")
	if err != nil {
		return err
	}
	return nil
}

// Close closes the file writer.
func (rs *ResultSaver) Close() error {
	if rs.fileWriter != nil {
		return rs.fileWriter.Close()
	}
	return nil
}

//Task struct
type Task struct{
	ID string
	Data string
}

// Process function
func (t *Task) Process() string{
	return fmt.Sprintf("Task %s with data %s", t.ID, t.Data)
}

// SharedQueue struct
type SharedQueue struct {
	queue chan Task
}

// NewSharedQueue creates a new SharedQueue.
func NewSharedQueue() *SharedQueue {
	return &SharedQueue{queue: make(chan Task, 20)} //make the channel buffered
}

// AddTask adds a task to the queue.
func (sq *SharedQueue) AddTask(task Task) {
	sq.queue <- task
}

// GetTask gets a task from the queue.
func (sq *SharedQueue) GetTask() Task {
	return <-sq.queue
}

// IsEmpty checks if the queue is empty
func (sq *SharedQueue) IsEmpty() bool{
	return len(sq.queue) == 0
}
