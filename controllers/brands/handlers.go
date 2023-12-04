package brands

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fardinabir/auth-guard/service"
	"github.com/fardinabir/product-store/controllers"
	"github.com/fardinabir/product-store/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// swagger:route GET /brands/ Brands GetBrands
// For getting the list of all the brands, shows both active and inactive brands
// responses:
//
//	500: ErrorResponse
//	200: SuccessRespBrands
func (rs *BrandResource) GetBrands(w http.ResponseWriter, r *http.Request) {
	qry := map[string]interface{}{}
	res, err := rs.Brands.GetBrands(qry)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	controllers.RespondWithJSON(w, http.StatusOK, res)
}

// swagger:route GET /brands/{id} Brands GetBrandsById
// Get's the specific brand of given brand id
// responses:
//
//	500: ErrorResponse
//	200: BrandSuccessResp200
func (rs *BrandResource) GetBrandsById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	resp, err := rs.Brands.GetBrandById(id)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	newResp := resp.GetBrandResp()
	controllers.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route POST /brands/new Brands CreateBrand
// For creating new brand entry
// responses:
//
//	500: ErrorResponse
//	200: BrandSuccessResp200
func (rs *BrandResource) CreateBrand(w http.ResponseWriter, r *http.Request) {
	var body models.BrandReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		controllers.ErrInvalidData.ErrorResponse().JSONResponse(w)
		return
	}
	brand, err := rs.Brands.Create(body.FormatToBrand())
	if err != nil {
		fmt.Println("Can't create the requested : ", err.Error())
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	newResp := brand.GetBrandResp()
	controllers.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route PATCH /brands/{id} Brands UpdateBrand
// Updates existing brand entry
// responses:
//
//	500: ErrorResponse
//	200: BrandSuccessResp200
func (rs *BrandResource) UpdateBrand(w http.ResponseWriter, r *http.Request) {
	var body models.BrandReq
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	brand, err := rs.Brands.GetBrandById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
			return
		}
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	brand.StatusId = body.StatusId
	brand.Name = body.Name
	brand, err = rs.Brands.Update(brand)
	if err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("Brand Updated : ", brand)
	newResp := brand.GetBrandResp()
	service.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route DELETE /brands/{id} Brands DeleteBrand
// For deleting existing brand
// responses:
//
//	500: ErrorResponse
//	200: DeleteResponse
func (rs *BrandResource) DeleteBrand(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	brand, err := rs.Brands.Delete(id)
	if err != nil {
		fmt.Println("Delete failed, brand not found")
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("Brand Deleted : ", brand)
	service.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Deleted Successfully"})
}
