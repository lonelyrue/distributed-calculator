package orchestrator

import (
    "errors"
    "sync"
)

type Repository struct {
    expressions map[string]Expression
    tasks       map[string]Task
    mu          sync.Mutex
}

func NewRepository() *Repository {
    return &Repository{
        expressions: make(map[string]Expression),
        tasks:       make(map[string]Task),
    }
}

func (r *Repository) AddExpression(id, expr string) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.expressions[id]; exists {
        return errors.New("expression with this ID already exists")
    }

    r.expressions[id] = Expression{ID: id, Expr: expr, Status: "pending"}
    return nil
}

func (r *Repository) AddTask(task Task) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.tasks[task.ID]; exists {
        return errors.New("task with this ID already exists")
    }

    r.tasks[task.ID] = task
    return nil
}

func (r *Repository) GetExpressions() []Expression {
    r.mu.Lock()
    defer r.mu.Unlock()

    var exprs []Expression
    for _, expr := range r.expressions {
        exprs = append(exprs, expr)
    }
    return exprs
}

func (r *Repository) GetExpressionByID(id string) (Expression, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    expr, exists := r.expressions[id]
    if !exists {
        return Expression{}, errors.New("expression not found")
    }
    return expr, nil
}

func (r *Repository) GetTask() (Task, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    for id, task := range r.tasks {
        if task.Status == "pending" {
            task.Status = "in_progress"
            r.tasks[id] = task // Update the task status in the repository
            return task, nil
        }
    }
    return Task{}, errors.New("no tasks available")
}
func (r *Repository) GetTaskResult(taskID string) (float64, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    task, exists := r.tasks[taskID]
    if !exists {
        return 0, errors.New("task not found")
    }

    if task.Status != "completed" {
        return 0, errors.New("task not completed")
    }

    return task.Result, nil
}
// func (r *Repository) UpdateTaskResult(id string, result float64) error {
//     r.mu.Lock()
//     defer r.mu.Unlock()

//     task, exists := r.tasks[id]
//     if !exists {
//         return errors.New("task not found")
//     }

//     task.Result = result
//     task.Status = "completed"
//     r.tasks[id] = task
//     return nil
// }
func (r *Repository) UpdateTaskResult(taskID string, result float64) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    task, exists := r.tasks[taskID]
    if !exists {
        return errors.New("task not found")
    }

    task.Result = result
    task.Status = "completed"
    r.tasks[taskID] = task
    return nil
}

func (r *Repository) UpdateExpressionResult(id string, result float64) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    expr, exists := r.expressions[id]
    if !exists {
        return errors.New("expression not found")
    }

    expr.Result = result
    expr.Status = "completed"
    r.expressions[id] = expr
    return nil
}

func (r *Repository) DeleteExpression(id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, exists := r.expressions[id]; !exists {
        return errors.New("expression not found")
    }

    delete(r.expressions, id)
    return nil
}