package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time-tracker/internal/model"

	"github.com/gin-gonic/gin"
)

// @Summary Получение всех пользователей
// @Tags Users
// @Description Получение массива данных всех пользователей
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Failure 400 {string} string "error"
// @Router /users [get]
func (h *Handler) getUsers(c *gin.Context) {

	filters := make(map[string]string)
	queryValues := c.Request.URL.Query()
	for key, values := range queryValues {
		if key != "limit" && key != "cursor" {

			filters[key] = values[0]
		}
	}

	limitStr := queryValues.Get("limit")
	cursorStr := queryValues.Get("cursor")

	var err error

	limit := 0
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}

	cursor := 0
	if cursorStr != "" {
		cursor, err = strconv.Atoi(cursorStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor parameter"})
			return
		}
	}

	users, err := h.services.Users.GetUsers(filters, limit, cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, users)

}

// @Summary Удаление пользователя
// @Tags Users
// @Description Удаление пользователя по его id
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.UserIDResponse
// @Failure 400 {string} string "error"
// @Failure 404 {string} string "user not found"
// @Router /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userIdStr := c.Param("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.services.Users.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"id": userId})

}

// @Summary Изменение данных пользователя
// @Tags Users
// @Description Изменить данные пользователя по его id (Можно отправлять не все поля, а только требующие замены)
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param data body model.UpdateUserInput true "Update User Data"
// @Success 200 {object} model.User
// @Failure 400 {string} string "error"
// @Failure 404 {string} string "user not found"
// @Router /users/{id} [patch]
func (h *Handler) updateUser(c *gin.Context) {
	userIdStr := c.Param("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req model.UpdateUserInput
	err = json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.services.Users.UpdateUser(userId, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

// @Summary Создание нового пользователя
// @Tags Users
// @Accept  json
// @Produce  json
// @Param data body model.UpdateUserInput true "User Data"
// @Success 200 {object} model.User
// @Failure 400 {string} string "error"
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var user model.User
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "json decoding error"})
		return
	}

	if user.PassportNumber == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "empty passportNumber field"})
		return
	}

	userId, err := h.services.Users.CreateUser(user)

	user.Id = userId

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
