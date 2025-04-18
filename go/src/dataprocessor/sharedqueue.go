package dataprocessor

type SharedQueue struct {
	queue chan Task
}

func NewSharedQueue() *SharedQueue {
	return &SharedQueue{
		queue: make(chan Task),
	}
}

func (sq *SharedQueue) AddTask(task Task) {
	sq.queue <- task
}

func (sq *SharedQueue) GetTask() (Task, bool) {
	task, ok := <-sq.queue
	return task, ok
}