package agent

import (
    "bytes"
    "encoding/json"
    "distributed-calculator/pkg/task"
    "net/http"
)

type Client struct {
    BaseURL string
}

func NewClient(baseURL string) *Client {
    return &Client{BaseURL: baseURL}
}

func (c *Client) FetchTask() (*task.Task, error) {
    resp, err := http.Get(c.BaseURL + "/internal/task")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, err
    }

    var result struct {
        Task task.Task `json:"task"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return &result.Task, nil
}

func (c *Client) SendResult(taskID string, result float64) error {
    reqBody, err := json.Marshal(map[string]interface{}{
        "id":     taskID,
        "result": result,
    })
    if err != nil {
        return err
    }

    resp, err := http.Post(c.BaseURL+"/internal/task", "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return err
    }

    return nil
}