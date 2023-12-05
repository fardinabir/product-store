package suppliers

import (
	"github.com/fardinabir/product-store/models"
)

type SupplierStore interface {
	Create(u *models.Supplier) (*models.Supplier, error)
	Update(u *models.Supplier) (*models.Supplier, error)
	Delete(id int) (*models.Supplier, error)
	GetSuppliers(q map[string]interface{}) ([]models.SupplierResp, error)
	GetSupplierById(id int) (*models.Supplier, error)
}
