package repos

import (
	"errors"
	"github.com/fardinabir/product-store/database"
	"github.com/fardinabir/product-store/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"strconv"
)

type ProductStore struct {
	DB *gorm.DB
}

func NewProductStore() *ProductStore {
	d := database.GetDbConnection()
	return &ProductStore{DB: d}
}

func (s *ProductStore) Create(b *models.Product) (*models.Product, error) {
	brnd := &models.Brand{}
	res := s.DB.Model(&models.Brand{}).Where("id = ?", b.BrandID).First(brnd)
	if res.Error != nil {
		log.Println("Error while getting brand in db", res.Error)
		return nil, res.Error
	}

	cat := &models.Category{}
	res = s.DB.Model(&models.Category{}).Where("id = ?", b.CategoryID).First(cat)
	if res.Error != nil {
		log.Println("Error while getting category in db", res.Error)
		return nil, res.Error
	}

	sup := &models.Supplier{}
	res = s.DB.Model(&models.Supplier{}).Where("id = ?", b.SupplierID).First(sup)
	if res.Error != nil {
		log.Println("Error while getting supplier in db", res.Error)
		return nil, res.Error
	}

	b.Brand = *brnd
	b.Category = *cat
	b.Supplier = *sup
	res = s.DB.Create(&b)
	if res.Error != nil {
		log.Println("Error while creating product in db", res.Error)
		return nil, res.Error
	}
	return b, nil
}

func (s *ProductStore) Update(b *models.Product) (*models.Product, error) {
	res := s.DB.Save(b)
	if res.Error != nil {
		log.Println("Error while updating product in db", res.Error)
		return nil, res.Error
	}
	return b, nil
}

func (s *ProductStore) Delete(id int) (*models.Product, error) {
	cat := &models.Product{}
	res := s.DB.Delete(&cat, id)
	if res.Error != nil {
		log.Println("Error while deleting product in db", id, res.Error)
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		log.Println("Record with ", id, " not found.")
		return nil, errors.New("not Found")
	}
	return cat, nil
}

func (s *ProductStore) GetProducts(qry map[string]string) ([]models.Product, error) {
	var products []models.Product
	tx := s.DB.Model(&models.Product{})

	if val, ok := qry["max_price"]; ok {
		tx = tx.Where("discount_price <= ? ", val)
		delete(qry, "max_price")
	}
	if val, ok := qry["min_price"]; ok {
		tx = tx.Where("discount_price >= ? ", val)
		delete(qry, "min_price")
	}
	var perPage, page int
	if val, ok := qry["per_page"]; ok {
		perPage, _ = strconv.Atoi(val)
		delete(qry, "per_page")
	}
	if val, ok := qry["page"]; ok {
		page, _ = strconv.Atoi(val)
		delete(qry, "page")
	}
	if perPage > 0 && page > 0 {
		offset := (page - 1) * perPage
		tx = tx.Limit(perPage).Offset(offset)
	}

	//res := s.DB.Model(&models.Brand{}).Where(q).Find(&brands)
	res := tx.Preload(clause.Associations).Where(qry).Find(&products)
	if res.Error != nil {
		log.Println("Fetching products list, ", qry, res.Error)
		return nil, res.Error
	}
	return products, nil
}

func (s *ProductStore) GetProductById(id int) (*models.Product, error) {
	cat := &models.Product{}
	res := s.DB.Model(&models.Product{}).Where("id = ?", id).First(cat)
	if res.Error != nil {
		log.Println("Error while getting product in db", id, res.Error)
		return nil, res.Error
	}
	return cat, nil
}
