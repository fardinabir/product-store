package brands

import (
	"github.com/fardinabir/product-store/models"
)

type BrandStore interface {
	Create(u *models.Brand) (*models.Brand, error)
	Update(u *models.Brand) (*models.Brand, error)
	Delete(id int) (*models.Brand, error)
	GetBrands(q map[string]interface{}) ([]models.BrandResp, error)
	GetBrandById(id int) (*models.Brand, error)
}
