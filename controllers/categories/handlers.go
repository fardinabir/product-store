package categories

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

// swagger:route GET /categories/ Categories GetCategories
// For getting the list of all the categories, shows both active and inactive categories
// responses:
//
//	500: ErrorResponse
//	200: SuccessRespCategories
func (rs *CategoryResource) GetCategories(w http.ResponseWriter, r *http.Request) {
	qry := map[string]interface{}{}
	res, err := rs.Categories.GetCategories(qry)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	controllers.RespondWithJSON(w, http.StatusOK, res)
}

// swagger:route GET /categories/{id} Categories GetCategoryById
// Gets the specific category of given category id
// responses:
//
//	500: ErrorResponse
//	200: CategorySuccessResp200
func (rs *CategoryResource) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	resp, err := rs.Categories.GetCategoryById(id)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	newResp := resp.GetCategoryResp()
	controllers.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route POST /categories/new Categories CreateCategory
// For creating new category entry
// responses:
//
//	500: ErrorResponse
//	200: CategorySuccessResp200
func (rs *CategoryResource) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var body models.CategoryReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		controllers.ErrInvalidData.ErrorResponse().JSONResponse(w)
		return
	}
	maxSeqId, err := rs.Categories.GetMaxSequence(body.ParentId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Failed to get categories : ", err.Error())
			controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
			return
		}
	}
	categoryBody := body.FormatToCategory()
	categoryBody.Sequence = maxSeqId + 1
	category, err := rs.Categories.Create(categoryBody)
	if err != nil {
		fmt.Println("Can't create the requested : ", err.Error())
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	newResp := category.GetCategoryResp()
	controllers.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route PATCH /categories/{id} Categories UpdateCategory
// Updates existing category entry
// responses:
//
//	500: ErrorResponse
//	200: CategorySuccessResp200
func (rs *CategoryResource) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var bodyReq models.CategoryReq
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	category, err := rs.Categories.GetCategoryById(id)
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
	category.Name = bodyReq.Name
	category.StatusId = bodyReq.StatusId
	if *category.ParentId != *bodyReq.ParentId {
		maxSeqId, err := rs.Categories.GetMaxSequence(bodyReq.ParentId)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println("Failed to get categories : ", err.Error())
				controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
				return
			}
		}
		category.ParentId = bodyReq.ParentId
		category.Sequence = maxSeqId + 1
	}
	category, err = rs.Categories.Update(category)
	if err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("Category Updated : ", category)
	newResp := category.GetCategoryResp()
	service.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route DELETE /categories/{id} Categories DeleteCategory
// For deleting existing category
// responses:
//
//	500: ErrorResponse
//	200: DeleteResponse
func (rs *CategoryResource) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	category, err := rs.Categories.Delete(id)
	if err != nil {
		fmt.Println("Delete failed, category not found")
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("Category Deleted : ", category)
	service.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Deleted Successfully"})
}
