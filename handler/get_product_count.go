package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vantu2801se/product-manager-system/models"
)

func (h *httpHandler) getProductCount(c *gin.Context) {
	productCount, err := h.sysCtx.RDSCli.GetProductCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	resp := struct {
		ProductCount []models.ProductCount `json:"categories"`
	}{
		ProductCount: productCount,
	}

	c.JSON(http.StatusOK, resp)
}
