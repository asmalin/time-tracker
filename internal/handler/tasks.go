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

	userIdStr := c.Param("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	start, err := time.Parse("2006-01-02", c.Query("start"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time start"})
		return
	}

	end, err := time.Parse("2006-01-02", c.Query("end"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time end"})
		return
	}

	tasks, err := h.services.Tasks.GetTasksForPeriod(userId, start, end)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
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
	userIdStr := c.Param("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	taskIdStr := c.Param("taskId")

	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.services.Tasks.FinishTask(userId, taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
