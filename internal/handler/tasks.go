package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"time-tracker/internal/model"

	"github.com/gin-gonic/gin"
)

// @Summary Получение задач пользователя за период
// @Description Получение списка задач пользователя за указанный период с указанием количества времени, затраченное на их выполенение. Список отсортирован по убыванию трудозатрат.
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Param start query string true "Period start time" example("2024-07-08")
// @Param end query string true "Period end time example("2024-07-22")"
// @Success 200 {array} model.TaskSummary
// @Failure 400 {string} string "error"
// @Failure 404 {string} string "user not found"
// @Router /users/{userID}/tasks [get]
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

// @Summary Начать отсчет времени задачи для пользователя
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Param data body model.TaskDataToCreate true "Task data to create"
// @Success 200 {object} model.Task
// @Failure 400 {string} string "error"
// @Failure 404 {string} string "user not found"
// @Router /users/{userID}/task [post]
func (h *Handler) startTask(c *gin.Context) {

	userIdStr := c.Param("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var taskData model.TaskDataToCreate
	if err := json.NewDecoder(c.Request.Body).Decode(&taskData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "json decoding error"})
		return
	}

	user, err := h.services.Users.GetUserById(userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.services.Tasks.StartTask(user.Id, taskData)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

// @Summary Закночить отсчет времени задачи для пользователя
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Param taskID path int true "Task ID"
// @Success 200 {object} model.Task
// @Failure 400 {string} string "error"
// @Failure 404 {string} string "user not found"
// @Router /users/{userID}/task/{taskID}/end [post]
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
