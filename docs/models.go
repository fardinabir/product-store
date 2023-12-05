package docs

import (
	"github.com/fardinabir/product-store/models"
)

// model for error response
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in:body
	Body struct {
		Msg string `json:"message"`
	} `json:"body"`
}

// model for get brands success response
// swagger:response SuccessRespBrands
type SuccessRespBrands struct {
	// in:body
	data []models.BrandResp
}

// model for create brands success
// swagger:response BrandSuccessResp200
type BrandSuccessResp200 struct {
	// in:body
	data models.BrandResp
}

// model for get categories success response
// swagger:response SuccessRespCategories
type SuccessRespCategories struct {
	// in:body
	data []models.CategoryResp
}

// model for create update categories success
// swagger:response CategorySuccessResp200
type CategorySuccessResp200 struct {
	// in:body
	data models.CategoryResp
}

// category tree response
// swagger:response CategoryTreeResp
type CategoryTreeResp struct {
	// in:body
	data []models.CategoryNode
}

// model for get suppliers success response
// swagger:response SuccessRespSuppliers
type SuccessRespSuppliers struct {
	// in:body
	data []models.SupplierResp
}

// model for create update supplier success
// swagger:response SupplierSuccessResp200
type SupplierSuccessResp200 struct {
	// in:body
	data models.SupplierResp
}

// model for get products success response
// swagger:response SuccessRespProducts
type SuccessRespProducts struct {
	// in:body
	data []models.ProductResp
}

// model for create update products success
// swagger:response ProductSuccessResp200
type ProductSuccessResp200 struct {
	// in:body
	data models.ProductResp
}

// model for delete response
// swagger:response DeleteResponse
type DeleteResponse struct {
	// in:body
	Body struct {
		Msg string `json:"message"`
	} `json:"body"`
}

// ########################### Request Models ############################

// model for adding brands
// swagger:parameters CreateBrand
type AddBrandRequest struct {
	// The request body for adding a brand.
	// in: body
	// required: true
	Body models.BrandReq
}

// model for updating brands
// swagger:parameters UpdateBrand
type UpdateBrandRequest struct {
	// path id for updating brand
	//
	// in:path
	ID int `json:"id"`
	// The request body for adding a brand.
	// in: body
	// required: true
	Body models.BrandReq
}

// model for adding categories
// swagger:parameters CreateCategory
type AddCategoryRequest struct {
	// in: body
	// required: true
	Body models.CategoryReq `json:"body"`
}

// model for updating categories
// swagger:parameters UpdateCategory
type UpdateCategoryRequest struct {
	// path id
	//
	// in:path
	ID int `json:"id"`
	// in: body
	// required: true
	Body models.CategoryReq `json:"body"`
}

// model for adding suppliers
// swagger:parameters CreateSupplier
type AddSupplierRequest struct {
	// in: body
	// required: true
	Body models.SupplierReq `json:"body"`
}

// model for updating suppliers
// swagger:parameters UpdateSupplier
type UpdateSupplierRequest struct {
	// path id
	//
	// in:path
	ID int `json:"id"`
	// in: body
	// required: true
	Body models.SupplierReq `json:"body"`
}

// model for adding products
// swagger:parameters CreateProduct
type AddProductRequest struct {
	// in: body
	// required: true
	Body models.ProductReq `json:"body"`
}

// model for updating products
// swagger:parameters UpdateProdcut
type UpdateProductRequest struct {
	// path id
	//
	// in:path
	ID int `json:"id"`
	// in: body
	// required: true
	Body models.SupplierReq `json:"body"`
}

// model for getting products
// swagger:parameters GetProducts
type GetProductRequest struct {

	// in:query
	Brand string `json:"brand"`
	// in:query
	Category string `json:"category"`
	// in:query
	Supplier string `json:"supplier"`
	// in:query
	ProductName string `json:"product_name"`
}

// model for id path param
// swagger:parameters GetBrandsById DeleteBrand GetCategoryById DeleteCategory GetSupplierById DeleteSupplier GetProductById DeleteProduct
type ReqIDParam struct {
	// The id of the brands, categories, suppliers
	//
	// in:path
	// required:true
	ID int `json:"id"`
}
