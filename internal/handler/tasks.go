package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"time-tracker/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserTasks(c *gin.Context) {

}

func (h *Handler) startTask(c *gin.Context) {

	userIdStr := c.Param("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var task model.Task
	if err := json.NewDecoder(c.Request.Body).Decode(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "json decoding error"})
		return
	}

	user, err := h.services.Users.GetUserById(userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	task.StartTime = time.Now()
	task.UserId = user.Id
	task.User = user

	task, err = h.services.Tasks.StartTask(task)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) endTask(c *gin.Context) {

}
