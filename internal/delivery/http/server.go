package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pndwrzk/taskhub-service/internal/common/response"
	"github.com/pndwrzk/taskhub-service/internal/constants"
	"github.com/pndwrzk/taskhub-service/internal/delivery/http/handler"
	"github.com/pndwrzk/taskhub-service/internal/middleware"
)

type Server struct {
	router *gin.Engine
}

func NewServer(
	userHandler *handler.UserHandler,
	taskHandler *handler.TaskHandler,

) *Server {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Response{
			Status:  constants.AppError,
			Message: fmt.Sprintf("Path %s", http.StatusText(http.StatusNotFound)),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Response{
			Status:  constants.AppSuccess,
			Message: "Welcome to taskhub Service",
		})
	})

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	api := r.Group("/api/v1")

	auth := api.Group("auth")
	auth.POST("/register", userHandler.Register)
	auth.POST("/login", userHandler.Login)
	auth.POST("/refresh", userHandler.RefreshToken)

	task := api.Group("tasks")
	task.Use(middleware.JWTAuth())
	{

		task.GET("", taskHandler.GetTasksByUser)
		task.POST("", taskHandler.CreateTask)
		task.PUT("/:id", taskHandler.UpdateTask)
		task.DELETE("/:id", taskHandler.DeleteTask)
	}

	return &Server{router: r}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
