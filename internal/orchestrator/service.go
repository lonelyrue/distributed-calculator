package orchestrator

import (
    "errors"
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

func (s *Service) GetExpressions() []Expression {
    return s.repo.GetExpressions()
}

func (s *Service) GetExpressionByID(id string) (Expression, error) {
    return s.repo.GetExpressionByID(id)
}

func (s *Service) GetTask() (Task, error) {
    return s.repo.GetTask()
}

func (s *Service) UpdateTaskResult(id string, result float64) error {
    return s.repo.UpdateTaskResult(id, result)
}