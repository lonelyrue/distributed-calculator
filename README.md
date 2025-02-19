# distributed-calculator
# Distributed Arithmetic Expression Calculator (короче, калькулятор арифмитических выражений)

## Описание проекта:

Это распределенный калькулятор арифметических выражений. Он состоит из двух компонентов:
1. **Оркестратор** — он принимает выражения, разбивает их на задачи и управляет их выполнением
2. **Агент** — выполняет задачи и возвращает результаты

## Как всё это работает:)

1. Пользователь отправляет выражение на оркестратор
2. Оркестратор разбивает выражение на задачи и добавляет их в очередь
3. Агенты запрашивают задачи, выполняют их и возвращают результаты
4. Оркестратор обновляет статус выражения и результат
   
# Запуск проекта

### Сборка и запуск

1. Необходимо установить Docker и Docker Compose.
2. Теперь клонируйте репозиторий:
  
   git clone https://github.com/yourusername/distributed-calculator.git
   cd distributed-calculator
   
3. И запустите проект с помощью докера:
  
   docker-compose up --build
## Примеры запросов

### Добавление выражения
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
### Получение списка выражений
curl --location 'localhost:8080/api/v1/expressions'
### Получение выражения по ID
curl --location 'localhost:8080/api/v1/expressions/1'
## Тестирование

Для тестирования используйте Postman коллекцию: [tests/postman_collection.json](tests/postman_collection.json).
---

### 4. **Примеры для проверки работы программы**

#### Примеры curl-запросов

1. **Добавление выражения**:
   ```bash
   curl --location 'localhost:8080/api/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data '{
     "expression": "2+2*2"
   }'
   
2. Получение списка выражений:
  
   curl --location 'localhost:8080/api/v1/expressions'
   
3. Получение выражения по ID:
  
   curl --location 'localhost:8080/api/v1/expressions/1'
   




## схема работы

Пользователь -> Оркестратор -> Агенты -> Оркестратор -> Пользователь

Структура проекта:
distributed-calculator/
├── cmd/
│   ├── orchestrator/
│   │   └── main.go
│   ├── agent/
│   │   └── main.go
├── internal/
│   ├── orchestrator/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── models.go
│   ├── agent/
│   │   ├── worker.go
│   │   └── client.go
├── pkg/
│   ├── task/
│   │   └── task.go
│   ├── expression/
│   │   └── expression.go
├── docker-compose.yml
├── Makefile
├── README.md
├── go.mod
├── go.sum
└── tests/
    ├── postman_collection.json
    └── curl_examples.txt

Описание структуры:

1. `cmd/`:
   - orchestrator/main.go — вход для оркестратора
   - agent/main.go — вход для агента

2. `internal/`:
   - orchestrator/ — логика оркестратора:
     - handler.go — HTTP-обработчики (endpoint'ы)
     - service.go — разбор выражений, управление задачами
     - repository.go — хранение данных (выражения, задачи)
     - models.go — структуры данных (выражения, задачи)
   - agent/ — логика агента:
     - worker.go — горутины для выполнения задач
     - client.go — HTTP-клиент для взаимодействия с оркестратором

3. `pkg/`:
   - task/ — общие структуры и функции для задач.
   - expression/ — общие структуры и функции для выражений

4. `docker-compose.yml`:
   - Конфигурация для запуска оркестратора и агентов в Docker.

5. `Makefile`:
   - Упрощает запуск и сборку проекта

6. `README.md`:
   - Документация проекта

7. `tests/`:
   - postman_collection.json — коллекция Postman для тестирования API.
   - curl_examples.txt — примеры curl-запросов


