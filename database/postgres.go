package database

import (
	"fmt"
	"github.com/fardinabir/product-store/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host           = "localhost"
	user           = "admin"
	password       = "admin"
	dbName         = "productsdb"
	port           = 5432
	runAutoMigrate = false
)

var dbInstance *gorm.DB

func GetDbConnection() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}
	var err error
	dbInstance, err = InitDb()
	if err != nil {
		log.Println("Failed to init DB")
		return nil
	}
	return dbInstance
}

func InitDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect Database")
		return nil, err
	}
	db.AutoMigrate(&models.Brand{}, &models.Category{}, &models.Supplier{}, &models.Product{}, &models.ProductsStock{})
	return db, nil
}
