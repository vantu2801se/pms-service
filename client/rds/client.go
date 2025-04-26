package rds

import (
	"github.com/vantu2801se/product-manager-system/config"
	"github.com/vantu2801se/product-manager-system/models"
	model "github.com/vantu2801se/product-manager-system/models"
	"github.com/vantu2801se/product-manager-system/static"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client interface {
	CreateProduct(p *model.ProductDto) (uint64, error)
	GetProducts(status []string, offset, limit int) ([]models.ProductModel, error)
	GetCategories() ([]*models.Category, error)
	GetProductCount() ([]models.ProductCount, error)
}

type rdsClient struct {
	db *gorm.DB
}

func NewRDSClient(config *config.Config) (Client, error) {
	db, err := gorm.Open(mysql.Open(config.RDS.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(config.RDS.MaxConn)
	sqlDB.SetMaxIdleConns(config.RDS.MaxIdle)
	sqlDB.SetConnMaxLifetime(config.RDS.MaxLifetime)

	return &rdsClient{
		db: db,
	}, nil
}

func (c *rdsClient) CreateProduct(p *model.ProductDto) (uint64, error) {
	productModel := &model.ProductModel{
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
	}

	if p.StockQuantity > 0 {
		productModel.Status = static.StatusInStock
	} else {
		productModel.Status = static.StatusOutStock
	}

	if err := c.db.Create(productModel).Error; err != nil {
		return 0, err
	}

	relatedModel := model.ProductCategoryModel{
		ProductID:  productModel.ID,
		CategoryID: p.CategoryID,
	}
	if err := c.db.Create(relatedModel).Error; err != nil {
		return 0, err
	}

	return p.ID, nil
}

func (c *rdsClient) GetProducts(status []string, offset, limit int) ([]models.ProductModel, error) {
	var products []model.ProductModel

	query := c.db.Where("status IN ?", status).Offset(offset).Limit(limit)

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (c *rdsClient) GetCategories() ([]*models.Category, error) {
	var categories []*model.Category

	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *rdsClient) GetProductCount() ([]models.ProductCount, error) {
	var result []models.ProductCount
	err := c.db.
		Table("pms_categories AS c").
		Select("c.id AS category_id, c.name AS name, COUNT(pc.product_id) AS count").
		Joins("LEFT JOIN pms_products_categories AS pc ON c.id = pc.category_id").
		Group("c.id, c.name").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
