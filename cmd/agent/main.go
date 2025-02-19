package main

import (
    "distributed-calculator/internal/agent"
    "log"
    "os"
    "strconv"
)

func main() {
    // Получение значения COMPUTING_POWER из переменных окружения
    computingPowerStr := os.Getenv("COMPUTING_POWER")
    computingPower, err := strconv.Atoi(computingPowerStr)
    if err != nil {
        computingPower = 5 // Значение по умолчанию
    }

    // Запуск агента
    log.Println("Агент запущен с количеством горутин:", computingPower)
    agent.Start(computingPower)
}