package suppliers

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
)

// swagger:route GET /suppliers/ Suppliers GetSuppliers
// For getting the list of all the suppliers, shows both active and inactive suppliers
// responses:
//
//	500: ErrorResponse
//	200: SuccessRespSuppliers
func (rs *SupplierResource) GetSuppliers(w http.ResponseWriter, r *http.Request) {
	qry := map[string]interface{}{}
	res, err := rs.Suppliers.GetSuppliers(qry)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	controllers.RespondWithJSON(w, http.StatusOK, res)
}

// swagger:route GET /suppliers/{id} Suppliers GetSupplierById
// Gets the specific supplier of given supplier id
// responses:
//
//	500: ErrorResponse
//	200: SupplierSuccessResp200
func (rs *SupplierResource) GetSupplierById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	resp, err := rs.Suppliers.GetSupplierById(id)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	controllers.RespondWithJSON(w, http.StatusOK, resp.GetSupplierResp())
}

// swagger:route POST /suppliers/new Suppliers CreateSupplier
// For creating new supplier entry
// responses:
//
//	500: ErrorResponse
//	200: SupplierSuccessResp200
func (rs *SupplierResource) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	var body models.SupplierReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		controllers.ErrInvalidData.ErrorResponse().JSONResponse(w)
		return
	}
	supplier, err := rs.Suppliers.Create(body.FormatToSupplier())
	if err != nil {
		fmt.Println("Can't create the requested : ", err.Error())
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	newResp := supplier.GetSupplierResp()
	controllers.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route PATCH /suppliers/{id} Suppliers UpdateSupplier
// Updates existing supplier entry
// responses:
//
//	500: ErrorResponse
//	200: SupplierSuccessResp200
func (rs *SupplierResource) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	var bodyReq models.SupplierReq
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	supplier, err := rs.Suppliers.GetSupplierById(id)
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
	supplier.Name = bodyReq.Name
	supplier.Email = bodyReq.Email
	supplier.Phone = bodyReq.Phone
	supplier.StatusId = bodyReq.StatusId
	supplier.IsVerifiedSupplier = bodyReq.IsVerifiedSupplier

	supplier, err = rs.Suppliers.Update(supplier)
	if err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("Supplier Updated : ", supplier)
	newResp := supplier.GetSupplierResp()
	controllers.RespondWithJSON(w, http.StatusOK, newResp)
}

// swagger:route DELETE /suppliers/{id} Suppliers DeleteSupplier
// For deleting existing supplier
// responses:
//
//	500: ErrorResponse
//	200: DeleteResponse
func (rs *SupplierResource) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	supplier, err := rs.Suppliers.Delete(id)
	if err != nil {
		fmt.Println("Delete failed, supplier not found")
		controllers.ErrNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("Supplier Deleted : ", supplier)
	controllers.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Deleted Successfully"})
}
