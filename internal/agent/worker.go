package orchestrator

import (
    "sync"
    "distributed-calculator/pkg/expression"
)

type Service struct {
    repo *Repository
    mu   sync.Mutex
}

func NewService(repo *Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) AddExpression(expr string) (string, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    id := generateID()
    tasks, err := expression.ParseExpression(expr)
    if err != nil {
        return "", err
    }

    s.repo.AddExpression(id, expr)
    for _, task := range tasks {
        s.repo.AddTask(task)
    }
    return id, nil
}

func (s *Service) GetTask() (task.Task, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    for _, task := range s.repo.tasks {
        if task.Status == "pending" {
            return task, nil
        }
    }
    return task.Task{}, errors.New("no tasks available")
}

func (s *Service) UpdateTaskResult(id string, result float64) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    task, exists := s.repo.tasks[id]
    if !exists {
        return errors.New("task not found")
    }

    task.Result = result
    task.Status = "completed"
    s.repo.tasks[id] = task

    // Обновляем статус выражения
    expr := s.repo.expressions[task.ExpressionID]
    expr.Result = result
    expr.Status = "completed"
    s.repo.expressions[task.ExpressionID] = expr

    return nil
}