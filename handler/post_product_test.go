package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	rdsMock "github.com/vantu2801se/product-manager-system/client/rds/mock"
	"github.com/vantu2801se/product-manager-system/static"
	"github.com/vantu2801se/product-manager-system/system"
	"gotest.tools/assert"
)

func Test_postProduct(t *testing.T) {
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
		Message string `json:"message,omitempty"`
		Err     string `json:"error,omitempty"`
	}

	testCases := []struct {
		name           string
		requestBody    []byte
		rdsPostProduct func() (uint64, error)
		expStatus      int
		expStructResp  respStruct
	}{
		{
			name: "Post product success",
			requestBody: []byte(`{
				"name": "test",
				"description": "test",
				"price": 1.0,
				"quantity": 1,
				"category_id": 1
			}`),
			rdsPostProduct: func() (uint64, error) {
				return 1, nil
			},
			expStatus:     http.StatusOK,
			expStructResp: respStruct{Message: "success"},
		},
		{
			name: "Post product failed with invalid request",
			requestBody: []byte(`{
				"name": "test",
				"description": "test",
				"price": 1.0,
				"quantity": 1
				"category_id": 1
			}`),
			rdsPostProduct: nil,
			expStatus:      http.StatusBadRequest,
			expStructResp:  respStruct{Err: "invalid request body"},
		},
		{
			name: "Post product failed with invalid request",
			requestBody: []byte(`{
				"name": "test",
				"description": "test",
				"price": 1.0,
				"quantity": 1,
				"category_id": 1
			}`),
			rdsPostProduct: func() (uint64, error) {
				return 0, errors.New("failed to create product")
			},
			expStatus:     http.StatusInternalServerError,
			expStructResp: respStruct{Err: "internal server error"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			// actual
			if tc.rdsPostProduct != nil {
				mockRDS.EXPECT().CreateProduct(gomock.Any()).Return(tc.rdsPostProduct())
			}

			req, _ := http.NewRequest(http.MethodPost, static.PathPostProduct, bytes.NewBuffer(tc.requestBody))
			h.router.ServeHTTP(w, req)

			// expected
			resp, _ := json.Marshal(tc.expStructResp)

			assert.Equal(t, tc.expStatus, w.Code)
			assert.Equal(t, string(resp), w.Body.String())
		})
	}
}
