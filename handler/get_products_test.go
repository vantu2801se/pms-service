package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	rdsMock "github.com/vantu2801se/product-manager-system/client/rds/mock"
	"github.com/vantu2801se/product-manager-system/models"
	"github.com/vantu2801se/product-manager-system/static"
	"github.com/vantu2801se/product-manager-system/system"
	"gotest.tools/assert"
)

func Test_getProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger, err := system.NewLoggerForTest()
	assert.NilError(t, err)
	mockRDS := rdsMock.NewMockClient(ctrl)
	sysCtx := &system.SystemContext{
		Logger: logger,
		RDSCli: mockRDS,
	}
	h := &httpHandler{sysCtx: sysCtx, router: gin.Default()}

	setRoutes(h)

	type respStruct struct {
		Products []models.ProductModel `json:"products,omitempty"`
		Err      string                `json:"error,omitempty"`
	}

	testCases := []struct {
		name          string
		rdsGetProduct func() ([]models.ProductModel, error)
		expStatus     int
		expStructResp respStruct
	}{
		{
			name: "Get product all status success",
			rdsGetProduct: func() ([]models.ProductModel, error) {
				return []models.ProductModel{
					{
						ID:            1,
						Name:          "product 1",
						Description:   "description 1",
						Price:         100,
						StockQuantity: 10,
						Status:        "in stock",
					},
				}, nil
			},
			expStatus:     http.StatusOK,
			expStructResp: respStruct{Products: []models.ProductModel{{ID: 1, Name: "product 1", Description: "description 1", Price: 100, StockQuantity: 10, Status: "in stock"}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			// actual
			if tc.rdsGetProduct != nil {
				mockRDS.EXPECT().GetProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.rdsGetProduct())
			}

			req, _ := http.NewRequest(http.MethodGet, static.PathGetProducts, nil)
			h.router.ServeHTTP(w, req)

			// expected
			resp, _ := json.Marshal(tc.expStructResp)

			assert.Equal(t, tc.expStatus, w.Code)
			assert.Equal(t, string(resp), w.Body.String())
		})
	}
}
