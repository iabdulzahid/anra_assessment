package handler

import (
	"encoding/json"
	"net/http"

	"github.com/iabdulzahid/anra_assessment/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

type createTaskRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var req createTaskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(req.Title, req.Status)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, task, http.StatusCreated)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {

	tasks := h.service.ListTasks()

	writeJSON(w, tasks, http.StatusOK)
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, message string, status int) {

	resp := errorResponse{
		Error: message,
	}

	writeJSON(w, resp, status)
}
