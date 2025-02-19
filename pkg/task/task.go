package task

import (
    "time"
)

type Task struct {
    ID             string
    Arg1           float64
    Arg2           float64
    Operation      string
    OperationTime  int
    Status         string
    Result         float64
}

func (t *Task) Compute() float64 {
    time.Sleep(time.Duration(t.OperationTime) * time.Millisecond)

    switch t.Operation {
    case "+":
        return t.Arg1 + t.Arg2
    case "-":
        return t.Arg1 - t.Arg2
    case "*":
        return t.Arg1 * t.Arg2
    case "/":
        return t.Arg1 / t.Arg2
    default:
        return 0
    }
}