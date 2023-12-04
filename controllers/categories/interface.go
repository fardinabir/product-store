package categories

import (
	"github.com/fardinabir/product-store/models"
)

type CategoryStore interface {
	Create(u *models.Category) (*models.Category, error)
	Update(u *models.Category) (*models.Category, error)
	Delete(id int) (*models.Category, error)
	GetCategories(q map[string]interface{}) ([]models.CategoryResp, error)
	GetCategoriesTree(parentId *uint) ([]models.CategoryNode, error)
	GetCategoryById(id int) (*models.Category, error)
	GetMaxSequence(parentId *uint) (uint, error)
}
