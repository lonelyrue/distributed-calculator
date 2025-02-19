package expression

import (
    "distributed-calculator/pkg/task"
    "errors"
    "strconv"
    "strings"
)

// ParseExpression разбирает выражение на задачи
func ParseExpression(expr string) ([]task.Task, error) {
    // Удаляем пробелы
    expr = strings.ReplaceAll(expr, " ", "")

    // Разбираем выражение на токены
    tokens := tokenize(expr)

    // Преобразуем токены в обратную польскую запись (RPN)
    rpn, err := shuntingYard(tokens)
    if err != nil {
        return nil, err
    }

    // Строим задачи из RPN
    tasks, err := buildTasks(rpn)
    if err != nil {
        return nil, err
    }

    return tasks, nil
}

// tokenize разбивает выражение на токены
func tokenize(expr string) []string {
    var tokens []string
    var buffer strings.Builder

    for _, char := range expr {
        if isOperator(string(char)) || char == '(' || char == ')' {
            if buffer.Len() > 0 {
                tokens = append(tokens, buffer.String())
                buffer.Reset()
            }
            tokens = append(tokens, string(char))
        } else {
            buffer.WriteRune(char)
		}
    }

    if buffer.Len() > 0 {
        tokens = append(tokens, buffer.String())
    }

    return tokens
}

// shuntingYard преобразует токены в обратную польскую запись
func shuntingYard(tokens []string) ([]string, error) {
    var output []string
    var stack []string

    for _, token := range tokens {
        if isNumber(token) {
            output = append(output, token)
        } else if isOperator(token) {
            for len(stack) > 0 && isOperator(stack[len(stack)-1]) && precedence(stack[len(stack)-1]) >= precedence(token) {
                output = append(output, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, token)
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
            stack = stack[:len(stack)-1]
        }
    }

    for len(stack) > 0 {
        if stack[len(stack)-1] == "(" || stack[len(stack)-1] == ")" {
            return nil, errors.New("mismatched parentheses")
        }
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
        } else if isOperator(token) {
            if len(stack) < 2 {
                return nil, errors.New("invalid expression")
            }

            arg2 := stack[len(stack)-1]
            arg1 := stack[len(stack)-2]
            stack = stack[:len(stack)-2]

            taskID := generateID()
            task := task.Task{
                ID:             taskID,
                Arg1:           parseNumber(arg1),
                Arg2:           parseNumber(arg2),
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

// Вспомогательные функции
func isNumber(token string) bool {
    _, err := strconv.ParseFloat(token, 64)
    return err == nil
}

func isOperator(token string) bool {
    return token == "+" || token == "-" || token == "*" || token == "/"
}

func precedence(op string) int {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    default:
        return 0
    }
}

func parseNumber(token string) float64 {
    num, _ := strconv.ParseFloat(token, 64)
    return num
}

func getOperationTime(op string) int {
    switch op {
    case "+":
        return 1000
    case "-":
        return 1000
    case "*":
        return 2000
    case "/":
        return 3000
    default:
        return 0
    }
}

func generateID() string {
    return strconv.Itoa(int(time.Now().UnixNano()))
}