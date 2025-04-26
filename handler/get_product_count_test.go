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

func Test_getProductCount(t *testing.T) {
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
		ProductCount []models.ProductCount `json:"categories,omitempty"`
		Err          string                `json:"error,omitempty"`
	}

	testCases := []struct {
		name               string
		rdsGetProductCount func() ([]models.ProductCount, error)
		expStatus          int
		expStructResp      respStruct
	}{
		{
			name: "Get product all status success",
			rdsGetProductCount: func() ([]models.ProductCount, error) {
				return []models.ProductCount{{CategoryID: 1, Name: "category 1", Count: 1}}, nil
			},
			expStatus:     http.StatusOK,
			expStructResp: respStruct{ProductCount: []models.ProductCount{{CategoryID: 1, Name: "category 1", Count: 1}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			// actual
			if tc.rdsGetProductCount != nil {
				mockRDS.EXPECT().GetProductCount().Return(tc.rdsGetProductCount())
			}

			req, _ := http.NewRequest(http.MethodGet, static.PathGetProductCount, nil)
			h.router.ServeHTTP(w, req)

			// expected
			resp, _ := json.Marshal(tc.expStructResp)

			assert.Equal(t, tc.expStatus, w.Code)
			assert.Equal(t, string(resp), w.Body.String())
		})
	}
}
