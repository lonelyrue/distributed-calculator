
package expression

import (
    "errors"
    "distributed-calculator/pkg/task"
    "strconv"
    "strings"
    "unicode"
    "github.com/google/uuid"
    "fmt"
)

// ParseExpression разбирает выражение на задачи
func ParseExpression(expr string) ([]task.Task, error) {
    expr = strings.ReplaceAll(expr, " ", "")
    tokens := tokenize(expr)
    rpn, err := toRPN(tokens)
    if err != nil {
        return nil, err
    }
    return buildTasks(rpn)
}

// tokenize разбивает выражение на токены
func tokenize(expression string) []string {
    var tokens []string
    var current string

    for _, r := range expression {
        if unicode.IsSpace(r) {
            continue
        }
        if isOperator(r) || r == '(' || r == ')' {
            if current != "" {
                tokens = append(tokens, current)
                current = ""
            }
            tokens = append(tokens, string(r))
        } else if unicode.IsDigit(r) || r == '.' {
            current += string(r)
        } else {
            return nil // некорректный символ
        }
    }
    if current != "" {
        tokens = append(tokens, current)
    }
    return tokens
}

// toRPN преобразует токены в обратную польскую запись (RPN)
func toRPN(tokens []string) ([]string, error) {
    var output []string
    var stack []string

    precedence := map[string]int{
        "+": 1,
        "-": 1,
        "*": 2,
        "/": 2,
        "(": 0,
    }

    for _, token := range tokens {
        if isNumber(token) {
            output = append(output, token)
        } else if token == "(" {
            stack = append(stack, token)
        } else if token == ")" {
            for len(stack) > 0 && stack[len(stack)-1] != "(" {
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            if len(stack) == 0 {
                return nil, errors.New("mismatched parentheses")
            }
            stack = stack[:len(stack)-1] // Убираем '('
        } else if isOperator(rune(token[0])) {
            for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, token)
        } else {
            return nil, fmt.Errorf("unexpected token: %s", token)
        }
    }

    for len(stack) > 0 {
        output = append(output, stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }

    return output, nil
}

// buildTasks строит задачи из RPN
func buildTasks(rpn []string) ([]task.Task, error) {
    var stack []string
    var tasks []task.Task

    for _, token := range rpn {
        if isNumber(token) {
            stack = append(stack, token)
        } else if isOperator(rune(token[0])) {
            if len(stack) < 2 {
                return nil, errors.New("invalid expression")
            }

            arg2 := stack[len(stack)-1]
            arg1 := stack[len(stack)-2]
            stack = stack[:len(stack)-2]

            taskID := generateID()
            task := task.Task{
                ID:             taskID,
                Arg1:           arg1,
                Arg2:           arg2,
                Operation:      token,
                OperationTime:  getOperationTime(token),
                Status:         "pending",
            }
            tasks = append(tasks, task)

            stack = append(stack, taskID)
        }
    }

    return tasks, nil
}

// isNumber проверяет, является ли токен числом
func isNumber(token string) bool {
    _, err := strconv.ParseFloat(token, 64)
    return err == nil
}

// isOperator проверяет, является ли символ оператором
func isOperator(r rune) bool {
    return r == '+' || r == '-' || r == '*' || r == '/'
}

// generateID генерирует уникальный идентификатор
func generateID() string {
    return uuid.NewString()
}

// getOperationTime возвращает время выполнения операции
func getOperationTime(op string) int {
    switch op {
    case "+", "-":
        return 1000
    case "*", "/":
        return 2000
    default:
        return 0
    }
}