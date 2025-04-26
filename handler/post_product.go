package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/vantu2801se/product-manager-system/models"
)

type ProductRequest struct {
	Name        string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    uint64  `json:"quantity"`
	CategoryID  uint64  `json:"category_id"`
}

func (h *httpHandler) postProduct(c *gin.Context) {
	var product ProductRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if _, err := h.sysCtx.RDSCli.CreateProduct(&model.ProductDto{
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		StockQuantity: product.Quantity,
		CategoryID:    product.CategoryID,
	}); err != nil {
		h.sysCtx.Logger.Errorf("failed to create product. err: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
