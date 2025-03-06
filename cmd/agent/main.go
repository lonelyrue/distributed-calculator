package main

import (
    "distributed-calculator/internal/agent"
    "log"
    "os"
    "strconv"
)

func main() {
    // Получаем значение COMPUTING_POWER из переменных окружения
    computingPowerStr := os.Getenv("COMPUTING_POWER")
    computingPower, err := strconv.Atoi(computingPowerStr)
    if err != nil {
        computingPower = 5 // Значение по умолчанию
    }

    // Запускаем агента
    log.Println("Агент запущен с количеством горутин:", computingPower)
    for i := 0; i < computingPower; i++ {
        go agent.Worker(agent.NewClient("http://localhost:8080"))
    }

    // Бесконечный цикл, чтобы агент продолжал работать
    select {}
}