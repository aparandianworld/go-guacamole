package database

import (
	"context"

	"github.com/aparandianworld/go-guacamole/internal/models"
)

func (c Client) GetAllProducts(ctx context.Context, vendor_id string) ([]models.Product, error) {
	var products []models.Product

	result := c.DB.WithContext(ctx).Where(models.Product{VendorID: vendor_id}).Find(&products)
	return products, result.Error
}
