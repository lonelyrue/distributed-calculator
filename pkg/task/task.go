package task

import (
    "errors"
    "strconv"
    "time"
)

type Task struct {
    ID            string
    Arg1          string 
    Arg2          string 
    Operation     string
    OperationTime int
    Status        string
    Result        float64
}

func (t *Task) Compute() (float64, error) {
    // Преобразуем Arg1 и Arg2 в числа
    arg1, err := strconv.ParseFloat(t.Arg1, 64)
    if err != nil {
        return 0, errors.New("invalid Arg1: not a number")
    }

    arg2, err := strconv.ParseFloat(t.Arg2, 64)
    if err != nil {
        return 0, errors.New("invalid Arg2: not a number")
    }

    // Имитируем задержку выполнения операции
    time.Sleep(time.Duration(t.OperationTime) * time.Millisecond)

    // Выполняем операцию
    switch t.Operation {
    case "+":
        return arg1 + arg2, nil
    case "-":
        return arg1 - arg2, nil
    case "*":
        return arg1 * arg2, nil
    case "/":
        if arg2 == 0 {
            return 0, errors.New("division by zero")
        }
        return arg1 / arg2, nil
    default:
        return 0, errors.New("unsupported operation")
    }
}