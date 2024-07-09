package handler

import (
	"os"
	"time"
	"time-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(log *logrus.Logger) *gin.Engine {

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()

	router.Use(Logger(log), gin.Recovery())

	router.GET("/users", h.getUsers)
	router.POST("/users", h.createUser)
	router.DELETE("/users/:userId", h.deleteUser)
	router.PUT("/users/:userId", h.updateUser)
	router.GET("/users/:userId/tasks", h.getUserTasks)
	router.POST("/users/:userId/task", h.startTask)
	router.POST("/users/:userId/task/:taskId/end", h.endTask)

	return router
}

func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		log.WithFields(logrus.Fields{
			"status_code":  c.Writer.Status(),
			"latency_time": latencyTime,
			"client_ip":    c.ClientIP(),
			"method":       c.Request.Method,
			"path":         c.Request.RequestURI,
		}).Info("Request completed")
	}
}
