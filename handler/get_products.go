package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vantu2801se/product-manager-system/static"
)

type ProductResponse struct {
	Name         string  `json:"production_name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Quantity     uint64  `json:"quantity"`
	CategoryName uint64  `json:"category_id"`
	Status       string  `json:"status"`
}

func (h *httpHandler) getProducts(c *gin.Context) {
	status := c.QueryArray(static.StatusKey)
	if len(status) == 0 {
		status = []string{static.StatusInStock, static.StatusOutStock}
	}

	for _, s := range status {
		_, ok := static.StatusMap[s]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
			return
		}
	}

	offset, _ := strconv.Atoi(c.DefaultQuery(static.OffsetKey, static.OffsetDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery(static.LimitKey, static.LimitDefault))

	products, err := h.sysCtx.RDSCli.GetProducts(status, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}
