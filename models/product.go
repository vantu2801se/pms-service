package models

type ProductDto struct {
	ID            uint64
	Name          string
	Description   string
	Price         float64
	StockQuantity uint64
	CategoryID    uint64
	CategoryName  string
}

type ProductModel struct {
	ID            uint64  `json:"id"`
	Name          string  `json:"product_name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity uint64  `json:"quantity"`
	Status        string  `json:"status"`
}

func (p *ProductModel) TableName() string {
	return "pms_products"
}

type ProductCategoryModel struct {
	ProductID  uint64
	CategoryID uint64
}

func (pc *ProductCategoryModel) TableName() string {
	return "pms_products_categories"
}

type ProductCount struct {
	CategoryID uint64 `json:"category_id"`
	Name       string `json:"name"`
	Count      int64  `json:"count"`
}
