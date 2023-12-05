package models

import (
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
	return &BrandResp{
		ID:        b.ID,
		Name:      b.Name,
		StatusId:  b.StatusId,
		CreatedAt: b.CreatedAt.String(),
	}
}

type BrandResp struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	StatusId  string `json:"status_id"`
	CreatedAt string `json:"created_at"`
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
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ParentId  *uint  `json:"parent_id"`
	Sequence  uint   `json:"sequence"`
	StatusId  string `json:"status_id"`
	CreatedAt string `json:"created_at"`
}

func (b *Category) GetCategoryResp() *CategoryResp {
	return &CategoryResp{
		ID:        b.ID,
		Name:      b.Name,
		ParentId:  b.ParentId,
		Sequence:  b.Sequence,
		StatusId:  b.StatusId,
		CreatedAt: b.CreatedAt.String(),
	}
}

type Supplier struct {
	*gorm.Model
	SupplierReq
}

type SupplierResp struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	StatusId           string `json:"status_id"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier"`
	CreatedAt          string `json:"created_at"`
}

type SupplierReq struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	StatusId           string `json:"status_id"`
	IsVerifiedSupplier bool   `json:"is_verified_supplier"`
}

func (b *SupplierReq) FormatToSupplier() *Supplier {
	return &Supplier{nil, *b}
}

func (s *Supplier) GetSupplierResp() *SupplierResp {
	return &SupplierResp{
		ID:                 s.ID,
		Name:               s.Name,
		Email:              s.Email,
		Phone:              s.Phone,
		StatusId:           s.StatusId,
		IsVerifiedSupplier: s.IsVerifiedSupplier,
		CreatedAt:          s.CreatedAt.String(),
	}
}

type Product struct {
	*gorm.Model
	Name           string `json:"name"`
	Description    string `json:"description" gorm:"column:description;size:100"`
	Specifications string `json:"specifications"`

	UnitPrice     int    `json:"unit_price"`
	DiscountPrice int    `json:"discount_price"`
	Tags          string `json:"tags"`
	StatusId      string `json:"status_id"`
	BrandID       uint   `json:"brands_id"`
	Brand         Brand
	CategoryID    uint `json:"categories_id"`
	Category      Category
	SupplierID    uint `json:"suppliers_id"`
	Supplier      Supplier
}

type ProductResp struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description" gorm:"column:description;size:100"`
	Specifications string `json:"specifications"`
	BrandID        uint   `json:"brand_id"`
	CategoryID     uint   `json:"category_id"`
	SupplierID     uint   `json:"supplier_id"`
	UnitPrice      int    `json:"unit_price"`
	DiscountPrice  int    `json:"discount_price"`
	Tags           string `json:"tags"`
	StatusId       string `json:"status_id"`
}

func (p *Product) GetProductResp() *ProductResp {
	return &ProductResp{
		ID:             p.ID,
		Name:           p.Name,
		Description:    p.Description,
		Specifications: p.Specifications,
		BrandID:        p.BrandID,
		CategoryID:     p.CategoryID,
		SupplierID:     p.SupplierID,
		UnitPrice:      p.UnitPrice,
		DiscountPrice:  p.DiscountPrice,
		Tags:           p.Tags,
		StatusId:       p.StatusId,
	}
}

type ProductReq struct {
	Name           string `json:"name"`
	Description    string `json:"description" gorm:"column:description;size:100"`
	Specifications string `json:"specifications"`
	BrandsID       uint   `json:"brands_id"`
	CategoriesID   uint   `json:"categories_id"`
	SuppliersID    uint   `json:"suppliers_id"`
	UnitPrice      int    `json:"unit_price"`
	DiscountPrice  int    `json:"discount_price"`
	Tags           string `json:"tags"`
	StatusId       string `json:"status_id"`
}

func (p *ProductReq) FormatToProduct() *Product {
	return &Product{
		Model:          nil,
		Name:           p.Name,
		Description:    p.Description,
		Specifications: p.Specifications,
		UnitPrice:      p.UnitPrice,
		DiscountPrice:  p.DiscountPrice,
		Tags:           p.Tags,
		StatusId:       p.StatusId,
		BrandID:        p.BrandsID,
		CategoryID:     p.CategoriesID,
		SupplierID:     p.SuppliersID,
	}
}

type ProductsStock struct {
	*gorm.Model
	ProductsID    uint
	StockQuantity int
}
