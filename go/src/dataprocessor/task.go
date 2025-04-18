package main // Important: Belongs to the main package now

import (
    "fmt"
    "math/rand"
    "time"
)

type Task struct {
    ID   string
    Data string
}

func (t *Task) Process() string {
    // Simulate processing time
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    return fmt.Sprintf("Processed: %s (Task ID: %s)", t.Data, t.ID)
}