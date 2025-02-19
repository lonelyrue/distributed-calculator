package agent

import (
    "distributed-calculator/pkg/task"
    "log"
    "time"
)

func Worker(client *Client) {
    for {
        task, err := client.FetchTask()
        if err != nil {
            log.Println("No tasks available, retrying...")
            time.Sleep(1 * time.Second)
            continue
        }

        result := task.Compute()
        if err := client.SendResult(task.ID, result); err != nil {
            log.Println("Failed to send result:", err)
        }
    }
}