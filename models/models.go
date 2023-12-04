package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Brand struct {
	*gorm.Model
	Name     string ` json:"name"`
	StatusId string `json:"status_id"`
}

type BrandReq struct {
	Name     string `json:"name"`
	StatusId string `json:"status_id"`
}

func (b *BrandReq) FormatToBrand() *Brand {
	return &Brand{
		Model:    nil,
		Name:     b.Name,
		StatusId: b.StatusId,
	}
}

func (b *Brand) GetBrandResp() *BrandResp {
	var newResp BrandResp
	bytResp, _ := json.Marshal(b)
	json.Unmarshal(bytResp, &newResp)
	return &newResp
}

type BrandResp struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StatusId string `json:"status_id"`
}

type Category struct {
	*gorm.Model
	Name          string      `json:"name"`
	ParentId      *uint       `json:"parent_id,omitempty"`
	Sequence      uint        `json:"sequence,omitempty"`
	StatusId      string      `json:"status_id"`
	SubCategories []*Category `gorm:"foreignKey:ParentId"`
}

type CategoryNode struct {
	ID       uint           `json:"id"`
	Name     string         `json:"category_name"`
	Children []CategoryNode `json:"children"`
}

type CategoryReq struct {
	Name     string `json:"name"`
	ParentId *uint  `json:"parent_id,omitempty"`
	StatusId string `json:"status_id"`
}

func (b *CategoryReq) FormatToCategory() *Category {
	return &Category{
		Model:         nil,
		Name:          b.Name,
		ParentId:      b.ParentId,
		Sequence:      0,
		StatusId:      b.StatusId,
		SubCategories: nil,
	}
}

type CategoryResp struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ParentId uint   `json:"parent_id"`
	Sequence uint   `json:"sequence"`
	StatusId string `json:"status_id"`
}

func (b *Category) GetCategoryResp() *CategoryResp {
	return &CategoryResp{
		ID:       b.ID,
		Name:     b.Name,
		ParentId: *b.ParentId,
		Sequence: b.Sequence,
		StatusId: b.StatusId,
	}
}

type Supplier struct {
	*gorm.Model
	Name               string
	Email              string
	Phone              string
	StatusId           string
	IsVerifiedSupplier bool
}

type Product struct {
	*gorm.Model
	Name           string
	Description    string
	Specifications string
	BrandsID       uint
	CategoriesID   uint
	SuppliersID    uint
	UnitPrice      int
	DiscountPrice  int
	Tags           string
	StatusId       string
}

type ProductsStock struct {
	*gorm.Model
	ProductsID    uint
	StockQuantity int
}
