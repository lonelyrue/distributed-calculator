package orchestrator

import (
    "encoding/json"
    "net/http"
    "strings"
)

type Handler struct {
    service *Service
}

func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Expression string `json:"expression"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
        return
    }

    id, err := h.service.AddExpression(req.Expression)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *Handler) GetExpressions(w http.ResponseWriter, r *http.Request) {
    expressions := h.service.GetExpressions()
    json.NewEncoder(w).Encode(map[string]interface{}{"expressions": expressions})
}

func (h *Handler) GetExpressionByID(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/api/v1/expressions/")
    expr, err := h.service.GetExpressionByID(id)
    if err != nil {
        http.Error(w, "Expression not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{"expression": expr})
}

func (h *Handler) Task(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        task, err := h.service.GetTask()
        if err != nil {
            http.Error(w, "No tasks available", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(map[string]interface{}{"task": task})
    case http.MethodPost:
        var req struct {
            ID     string  `json:"id"`
            Result float64 `json:"result"`
        }
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
            return
        }

        if err := h.service.UpdateTaskResult(req.ID, req.Result); err != nil {
            http.Error(w, "Task not found", http.StatusNotFound)
            return
        }

        w.WriteHeader(http.StatusOK)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}