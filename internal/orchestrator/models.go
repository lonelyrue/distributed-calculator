package orchestrator

import "distributed-calculator/pkg/task"

// Alias the Task type from the task package
type Task = task.Task

type Expression struct {
    ID     string
    Expr   string
    Status string
    Result float64 
}

