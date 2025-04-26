package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/vantu2801se/product-manager-system/static"
	"github.com/vantu2801se/product-manager-system/system"
)

type Handler interface {
	Start() int
	Shutdown(ctx context.Context) error
}

type httpHandler struct {
	sysCtx *system.SystemContext
	router *gin.Engine
}

func NewHttpHandler(sysCtx *system.SystemContext) Handler {
	return &httpHandler{
		sysCtx: sysCtx,
		router: gin.Default(),
	}
}

func (h *httpHandler) Start() int {
	setRoutes(h)

	if err := h.router.Run(h.sysCtx.Config.Port); err != nil {
		// h.sysCtx.Logger.Errorf("failed to start server. err: %s", err.Error())
		return static.ExitError
	}

	return static.ExitOK
}

func (h *httpHandler) Shutdown(ctx context.Context) error {
	return nil
}

func setRoutes(h *httpHandler) {
	h.router.GET(static.PathGetHealcheck, h.healthcheck)
	// h.router.Use(middleware.VerifyPathParamMiddleware(h.sysCtx))
	// h.router.Use(middleware.AuthMiddleware(h.sysCtx))
	h.router.GET(static.PathGetProducts, h.getProducts)
	h.router.GET(static.PathGetCategories, h.getCategories)
	h.router.GET(static.PathGetProductCount, h.getProductCount)
	h.router.POST(static.PathPostProduct, h.postProduct)

}
