package repos

import (
	"errors"
	"github.com/fardinabir/product-store/database"
	"github.com/fardinabir/product-store/models"
	"gorm.io/gorm"
	"log"
)

type BrandStore struct {
	DB *gorm.DB
}

func NewBrandStore() *BrandStore {
	d := database.GetDbConnection()
	return &BrandStore{DB: d}
}

func (s *BrandStore) Create(b *models.Brand) (*models.Brand, error) {
	res := s.DB.Create(&b)
	if res.Error != nil {
		log.Println("Error while creating brand in db", res.Error)
		return nil, res.Error
	}
	return b, nil
}

func (s *BrandStore) Update(b *models.Brand) (*models.Brand, error) {
	res := s.DB.Updates(b)
	if res.Error != nil {
		log.Println("Error while updating brand in db", res.Error)
		return nil, res.Error
	}
	return b, nil
}

func (s *BrandStore) Delete(id int) (*models.Brand, error) {
	brnd := &models.Brand{}
	res := s.DB.Delete(&brnd, id)
	if res.Error != nil {
		log.Println("Error while deleting brand in db", id, res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		log.Println("Record with ", id, " not found.")
		return nil, errors.New("not Found")
	}
	return brnd, nil
}

func (s *BrandStore) GetBrands(q map[string]interface{}) ([]models.BrandResp, error) {
	var brands []models.BrandResp
	res := s.DB.Model(&models.Brand{}).Where(q).Find(&brands)
	if res.Error != nil {
		log.Println("Fetching Brands list, ", q, res.Error)
		return nil, res.Error
	}
	return brands, nil
}

func (s *BrandStore) GetBrandById(id int) (*models.Brand, error) {
	brnd := &models.Brand{}
	res := s.DB.Model(&models.Brand{}).Where("id = ?", id).First(brnd)
	if res.Error != nil {
		log.Println("Error while getting brand in db", id, res.Error)
		return nil, res.Error
	}
	return brnd, nil
}
