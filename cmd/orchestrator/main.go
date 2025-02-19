package main

import (
    "log"
    "net/http"
    "distributed-calculator/internal/orchestrator"
)

func main() {
    // Инициализация репозитория и сервиса
    repo := orchestrator.NewRepository()
    service := orchestrator.NewService(repo)
    handler := orchestrator.NewHandler(service)

    // Настройка маршрутов
    http.HandleFunc("/api/v1/calculate", handler.Calculate)
    http.HandleFunc("/api/v1/expressions", handler.GetExpressions)
    http.HandleFunc("/api/v1/expressions/", handler.GetExpressionByID)
    http.HandleFunc("/internal/task", handler.Task)

    // Запуск сервера
    log.Println("Оркестратор запущен на :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
