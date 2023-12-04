package repos

import (
	"errors"
	"github.com/fardinabir/product-store/database"
	"github.com/fardinabir/product-store/models"
	"gorm.io/gorm"
	"log"
)

type CategoryStore struct {
	DB *gorm.DB
}

func NewCategoryStore() *CategoryStore {
	d := database.GetDbConnection()
	return &CategoryStore{DB: d}
}

func (s *CategoryStore) Create(b *models.Category) (*models.Category, error) {
	res := s.DB.Create(&b)
	if res.Error != nil {
		log.Println("Error while creating category in db", res.Error)
		return nil, res.Error
	}
	return b, nil
}

func (s *CategoryStore) Update(b *models.Category) (*models.Category, error) {
	res := s.DB.Updates(b)
	if res.Error != nil {
		log.Println("Error while updating category in db", res.Error)
		return nil, res.Error
	}
	return b, nil
}

func (s *CategoryStore) Delete(id int) (*models.Category, error) {
	cat := &models.Category{}
	res := s.DB.Delete(&cat, id)
	if res.Error != nil {
		log.Println("Error while deleting category in db", id, res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		log.Println("Record with ", id, " not found.")
		return nil, errors.New("not Found")
	}
	return cat, nil
}

func (s *CategoryStore) GetCategories(q map[string]interface{}) ([]models.CategoryResp, error) {
	var categories []models.CategoryResp
	res := s.DB.Model(&models.Category{}).Where(q).Find(&categories)
	if res.Error != nil {
		log.Println("Fetching categories list, ", q, res.Error)
		return nil, res.Error
	}
	return categories, nil
}

func (s *CategoryStore) GetMaxSequence(parentId *uint) (uint, error) {
	var category models.CategoryResp
	query := s.DB.Model(&models.Category{})
	if parentId == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", parentId)
	}
	res := query.Order("sequence desc").First(&category)
	if res.Error != nil {
		log.Println("Fetching max sequence, ", res.Error)
		return 0, res.Error
	}
	return category.Sequence, nil
}

func (s *CategoryStore) GetCategoryById(id int) (*models.Category, error) {
	cat := &models.Category{}
	res := s.DB.Model(&models.Category{}).Where("id = ?", id).First(cat)
	if res.Error != nil {
		log.Println("Error while getting category in db", id, res.Error)
		return nil, res.Error
	}
	return cat, nil
}

func (s *CategoryStore) GetCategoriesTree(parentId *uint) ([]models.CategoryNode, error) {
	var categories []*models.Category
	var err error
	query := s.DB.Model(&models.Category{})
	if parentId == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", parentId)
	}
	res := query.Order("sequence").Find(&categories)
	if res.Error != nil {
		log.Println("Error while getting category tree in db", res.Error)
		return nil, res.Error
	}

	categoryNodes := make([]models.CategoryNode, len(categories))
	for i := range categories {
		categoryNodes[i].ID = categories[i].ID
		categoryNodes[i].Name = categories[i].Name
		categoryNodes[i].Children, err = s.GetCategoriesTree(&categories[i].ID)
		if err != nil {
			log.Println("Error while getting category tree in db", res.Error)
			return nil, res.Error
		}
	}
	return categoryNodes, nil
}
