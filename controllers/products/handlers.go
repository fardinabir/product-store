package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fardinabir/product-store/controllers"
	"github.com/fardinabir/product-store/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// swagger:route GET /products/ Products GetProducts
// For getting the list of all the products, is filterable by fields
// responses:
//
//	500: ErrorResponse
//	200: SuccessRespProducts
func (rs *ProductResource) GetProducts(w http.ResponseWriter, r *http.Request) {
	qry := map[string]string{}

	productName := strings.TrimSpace(r.URL.Query().Get("product_name"))
	if productName != "" {
		qry["name"] = productName
	}
	minPrice := strings.TrimSpace(r.URL.Query().Get("min_price"))
	if minPrice != "" {
		qry["min_price"] = minPrice
	}
	maxPrice := strings.TrimSpace(r.URL.Query().Get("max_price"))
	if maxPrice != "" {
		qry["max_price"] = maxPrice
	}
	brand := strings.TrimSpace(r.URL.Query().Get("brand"))
	if brand != "" {
		qry["brand"] = brand
	}
	category := strings.TrimSpace(r.URL.Query().Get("category"))
	if category != "" {
		qry["category"] = category
	}
	supplier := strings.TrimSpace(r.URL.Query().Get("supplier"))
	if supplier != "" {
		qry["supplier"] = supplier
	}
	verifiedSupplier := strings.TrimSpace(r.URL.Query().Get("is_verified_supplier"))
	if verifiedSupplier != "" {
		qry["is_verified_supplier"] = verifiedSupplier
	}

	page := strings.TrimSpace(r.URL.Query().Get("page"))
	if page != "" {
		qry["page"] = page
	}
	per_page := strings.TrimSpace(r.URL.Query().Get("per_page"))
	if per_page != "" {
		qry["per_page"] = per_page
	}
	res, err := rs.Products.GetProducts(qry)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	controllers.RespondWithJSON(w, http.StatusOK, res)
}

// swagger:route GET /products/{id} Products GetProductById
// Gets the specific product of given product id
// responses:
//
//	500: ErrorResponse
//	200: ProductSuccessResp200
func (rs *ProductResource) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	resp, err := rs.Products.GetProductById(id)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	controllers.RespondWithJSON(w, http.StatusOK, resp.GetProductResp())
}

// swagger:route POST /products/new Products CreateProduct
// For creating new product entry
// responses:
//
//	500: ErrorResponse
//	200: ProductSuccessResp200
func (rs *ProductResource) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var body *models.ProductReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		controllers.ErrInvalidData.ErrorResponse().JSONResponse(w)
		return
	}

	product, err := rs.Products.Create(body.FormatToProduct())
	if err != nil {
		fmt.Println("Can't create the requested : ", err.Error())
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	newResp := product.GetProductResp()
	controllers.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route PATCH /products/{id} Products UpdateProduct
// Updates existing product entry
// responses:
//
//	500: ErrorResponse
//	200: ProductSuccessResp200
func (rs *ProductResource) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var bodyReq models.ProductReq
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	product, err := rs.Products.GetProductById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
			return
		}
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&bodyReq); err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	product.Name = bodyReq.Name
	product.Description = bodyReq.Description
	product.Specifications = bodyReq.Specifications
	product.BrandID = bodyReq.BrandsID
	product.CategoryID = bodyReq.CategoriesID
	product.SupplierID = bodyReq.SuppliersID
	product.UnitPrice = bodyReq.UnitPrice
	product.DiscountPrice = bodyReq.DiscountPrice
	product.Tags = bodyReq.Tags
	product.StatusId = bodyReq.StatusId

	product, err = rs.Products.Update(product)
	if err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("Product Updated : ", product)
	newResp := product.GetProductResp()
	controllers.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route DELETE /products/{id} Products DeleteProduct
// For deleting existing product
// responses:
//
//	500: ErrorResponse
//	200: DeleteResponse
func (rs *ProductResource) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	product, err := rs.Products.Delete(id)
	if err != nil {
		fmt.Println("Delete failed, product not found")
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("Product Deleted : ", product)
	controllers.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Deleted Successfully"})
}
