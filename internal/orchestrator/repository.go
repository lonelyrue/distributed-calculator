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

func (r *Repository) AddExpression(id, expr string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.expressions[id] = Expression{ID: id, Expr: expr, Status: "pending"}
}

func (r *Repository) AddTask(task Task) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.tasks[task.ID] = task
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

    for _, task := range r.tasks {
        if task.Status == "pending" {
            return task, nil
        }
    }
    return Task{}, errors.New("no tasks available")
}

func (r *Repository) UpdateTaskResult(id string, result float64) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    task, exists := r.tasks[id]
    if !exists {
        return errors.New("task not found")
    }

    task.Result = result
    task.Status = "completed"
    r.tasks[id] = task
    return nil
}