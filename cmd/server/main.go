package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iabdulzahid/anra_assessment/internal/handler"
	"github.com/iabdulzahid/anra_assessment/internal/repository"
	"github.com/iabdulzahid/anra_assessment/internal/service"
)

func main() {

	// Read port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	addr := ":" + port

	// Initialize dependencies
	repo := repository.NewTaskRepository()
	taskService := service.NewTaskService(repo)
	taskHandler := handler.NewTaskHandler(taskService)

	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			taskHandler.CreateTask(w, r)

		case http.MethodGet:
			taskHandler.ListTasks(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// HTTP Server configuration
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in goroutine
	go func() {

		log.Printf("Server started on %s\n", addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}

	}()

	// Listen for OS signals
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutdown signal received")

	// Graceful shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
