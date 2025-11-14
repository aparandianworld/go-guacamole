package database

import (
	"context"
	"errors"

	"github.com/aparandianworld/go-guacamole/internal/dberrors"
	"github.com/aparandianworld/go-guacamole/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllProducts(ctx context.Context, vendor_id string) ([]models.Product, error) {
	var products []models.Product

	result := c.DB.WithContext(ctx).Where(models.Product{VendorID: vendor_id}).Find(&products)
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()

	result := c.DB.WithContext(ctx).Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return product, nil
}
