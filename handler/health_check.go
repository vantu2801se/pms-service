package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
