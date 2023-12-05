package repos

import (
	"errors"
	"github.com/fardinabir/product-store/database"
	"github.com/fardinabir/product-store/models"
	"gorm.io/gorm"
	"log"
)

type SupplierStore struct {
	DB *gorm.DB
}

func NewSupplierStore() *SupplierStore {
	d := database.GetDbConnection()
	return &SupplierStore{DB: d}
}

func (s *SupplierStore) Create(b *models.Supplier) (*models.Supplier, error) {
	res := s.DB.Create(&b)
	if res.Error != nil {
		log.Println("Error while creating supplier in db", res.Error)
		return nil, res.Error
	}
	return b, nil
}

func (s *SupplierStore) Update(b *models.Supplier) (*models.Supplier, error) {
	res := s.DB.Save(b)
	if res.Error != nil {
		log.Println("Error while updating supplier in db", res.Error)
		return nil, res.Error
	}
	return b, nil
}

func (s *SupplierStore) Delete(id int) (*models.Supplier, error) {
	cat := &models.Supplier{}
	res := s.DB.Delete(&cat, id)
	if res.Error != nil {
		log.Println("Error while deleting supplier in db", id, res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		log.Println("Record with ", id, " not found.")
		return nil, errors.New("not Found")
	}
	return cat, nil
}

func (s *SupplierStore) GetSuppliers(q map[string]interface{}) ([]models.SupplierResp, error) {
	var categories []models.SupplierResp
	res := s.DB.Model(&models.Supplier{}).Where(q).Find(&categories)
	if res.Error != nil {
		log.Println("Fetching suppliers list, ", q, res.Error)
		return nil, res.Error
	}
	return categories, nil
}

func (s *SupplierStore) GetSupplierById(id int) (*models.Supplier, error) {
	cat := &models.Supplier{}
	res := s.DB.Model(&models.Supplier{}).Where("id = ?", id).First(cat)
	if res.Error != nil {
		log.Println("Error while getting supplier in db", id, res.Error)
		return nil, res.Error
	}
	return cat, nil
}
