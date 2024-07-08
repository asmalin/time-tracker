package handler

import (
	"time-tracker/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/users", h.getUsers)
	router.POST("/users", h.createUser)
	router.DELETE("/users/:userId", h.deleteUser)
	router.PUT("/users/:userId", h.updateUser)
	router.GET("/users/:userId/tasks", h.getUserTasks)
	router.POST("/users/:userId/task", h.startTask)
	router.POST("/users/:userId/task/:taskId/end", h.endTask)

	return router
}
