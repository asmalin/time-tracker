package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) deleteUser(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) createUser(c *gin.Context) {

}
