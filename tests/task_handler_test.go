package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/iabdulzahid/anra_assessment/internal/handler"
	"github.com/iabdulzahid/anra_assessment/internal/repository"
	"github.com/iabdulzahid/anra_assessment/internal/service"
)

func setupHandler() *handler.TaskHandler {

	repo := repository.NewTaskRepository()
	service := service.NewTaskService(repo)

	return handler.NewTaskHandler(service)
}

func TestCreateTaskSuccess(t *testing.T) {

	h := setupHandler()

	body := `{"title":"Learn Go"}`

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}
}

func TestCreateTaskValidationError(t *testing.T) {

	h := setupHandler()

	body := `{"title":""}`

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	rec := httptest.NewRecorder()

	h.CreateTask(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestListTasks(t *testing.T) {

	h := setupHandler()

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()

	h.ListTasks(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", rec.Code)
	}
}
