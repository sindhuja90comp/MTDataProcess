package main // Important: Belongs to the main package now

// SharedQueue struct
type SharedQueue struct {
    Queue chan Task // Exported field
}

// NewSharedQueue creates a new SharedQueue.
func NewSharedQueue() *SharedQueue {
    return &SharedQueue{Queue: make(chan Task, 20)} // Make the channel buffered and export Queue
}

// AddTask adds a task to the queue.
func (sq *SharedQueue) AddTask(task Task) {
    sq.Queue <- task
}

// GetTask gets a task from the queue.
func (sq *SharedQueue) GetTask() (Task, bool) {
    task, ok := <-sq.Queue
    return task, ok
}

// Close closes the queue channel
func (sq *SharedQueue) Close() {
    close(sq.Queue)
}