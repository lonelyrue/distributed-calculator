package orchestrator

import (
    "sync"
    "distributed-calculator/pkg/expression"
    "distributed-calculator/pkg/task"
    "github.com/google/uuid"
    "errors"
    "strconv"
    "log"
)

type Service struct {
    repo *Repository
    mu   sync.Mutex
}

func NewService(repo *Repository) *Service {
    return &Service{repo: repo}
}

// generateID генерирует уникальный идентификатор.
func generateID() string {
    return uuid.NewString()
}

func (s *Service) AddExpression(expr string) (string, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    id := generateID()

    // Парсинг выражения для получения задач
    tasks, err := expression.ParseExpression(expr)
    if err != nil {
        return "", err
    }

    // Добавление выражения в репозиторий
    if err := s.repo.AddExpression(id, expr); err != nil {
        return "", err
    }

    // Добавление задач в репозиторий
    for _, t := range tasks {
        if err := s.repo.AddTask(t); err != nil {
            // Откат: удаление выражения, если добавление задач не удалось
            _ = s.repo.DeleteExpression(id)
            return "", err
        }
    }

    // Запуск вычисления задач
    go s.processTasks(id)

    return id, nil
}

func (s *Service) processTasks(expressionID string) {
    for {
        task, err := s.repo.GetTask()
        if err != nil {
            break // Нет доступных задач
        }

        // Вычисление результата задачи
        result, err := evaluateTask(task, s.repo)
        if err != nil {
            log.Printf("Error evaluating task %s: %v", task.ID, err)
            continue
        }

        // Обновление результата задачи
        if err := s.repo.UpdateTaskResult(task.ID, result); err != nil {
            log.Printf("Error updating task result %s: %v", task.ID, err)
            continue
        }

        // Обновление результата выражения
        if err := s.repo.UpdateExpressionResult(expressionID, result); err != nil {
            log.Printf("Error updating expression result %s: %v", expressionID, err)
            continue
        }
    }
}

// evaluateTask вычисляет результат задачи
func evaluateTask(task task.Task, repo *Repository) (float64, error) {
    var arg1, arg2 float64
    var err error

    // Вычисляем Arg1
    if isNumber(task.Arg1) {
        arg1, err = strconv.ParseFloat(task.Arg1, 64)
        if err != nil {
            return 0, err
        }
    } else {
        arg1, err = repo.GetTaskResult(task.Arg1)
        if err != nil {
            return 0, err
        }
    }

    // Вычисляем Arg2
    if isNumber(task.Arg2) {
        arg2, err = strconv.ParseFloat(task.Arg2, 64)
        if err != nil {
            return 0, err
        }
    } else {
        arg2, err = repo.GetTaskResult(task.Arg2)
        if err != nil {
            return 0, err
        }
    }

    // Выполняем операцию
    switch task.Operation {
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

// isNumber проверяет, является ли строка числом
func isNumber(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}
// GetExpressions возвращает все выражения из репозитория.
func (s *Service) GetExpressions() []Expression {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.repo.GetExpressions()
}

// GetExpressionByID возвращает выражение по его идентификатору.
func (s *Service) GetExpressionByID(id string) (Expression, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.repo.GetExpressionByID(id)
}

// GetTask получает следующую задачу из репозитория.
func (s *Service) GetTask() (task.Task, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.repo.GetTask()
}

// UpdateTaskResult обновляет результат выполнения задачи по ее идентификатору.
func (s *Service) UpdateTaskResult(id string, result float64) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.repo.UpdateTaskResult(id, result)
}