package main

import (
	"github.com/pndwrzk/taskhub-service/config"
	"github.com/pndwrzk/taskhub-service/internal/delivery/http"
	"github.com/pndwrzk/taskhub-service/internal/delivery/http/handler"
	"github.com/pndwrzk/taskhub-service/internal/repository"
	"github.com/pndwrzk/taskhub-service/internal/usecase"
	"github.com/pndwrzk/taskhub-service/pkg/db"
)

func main() {
	config.Load()
	db.Connect()

	userRepo := repository.NewUserRepository(db.DB)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	taskRepo := repository.NewTaskRepository(db.DB)
	taskUC := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUC)

	server := http.NewServer(userHandler, taskHandler)

	server.Start(":" + config.App.AppPort)
}
