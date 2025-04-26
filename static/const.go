package static

const (
	PathGetHealcheck    = "/ping"
	PathGetProducts     = "/v1/products"
	PathPostProduct     = "/v1/products"
	PathGetProductCount = "/v1/products/count"
	// PathPutProduct      = "/v1/products/:product_id"
	// PathDeleteProduct = "/v1/products/:product_id"

	PathGetCategories = "/v1/categories"
)

const (
	ExitOK int = iota
	ExitError
	ExitStartFailed
	ExitPanic = 99
)

const (
	StatusInStock  = "in_stock"
	StatusOutStock = "out_of_stock"
)

var StatusMap = map[string]struct{}{
	StatusInStock:  {},
	StatusOutStock: {},
}

const (
	StatusKey     = "status"
	OffsetKey     = "offset"
	OffsetDefault = "0"
	LimitKey      = "limit"
	LimitDefault  = "10"
)
