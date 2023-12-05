package products

import (
	"github.com/fardinabir/product-store/models"
)

type ProductStore interface {
	Create(u *models.Product) (*models.Product, error)
	Update(u *models.Product) (*models.Product, error)
	Delete(id int) (*models.Product, error)
	GetProducts(q map[string]string) ([]models.Product, error)
	GetProductById(id int) (*models.Product, error)
}
