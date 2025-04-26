package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type category struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// Get all categories to select when post product
func (h *httpHandler) getCategories(c *gin.Context) {
	categories, err := h.sysCtx.RDSCli.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}

	respBody := []category{}
	for _, c := range categories {
		respBody = append(respBody, category{ID: c.ID, Name: c.Name})
	}

	resp := struct {
		Categories []category `json:"categories"`
	}{
		Categories: respBody,
	}

	c.JSON(http.StatusOK, resp)
}
