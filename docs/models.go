package docs

import "github.com/fardinabir/product-store/models"

// model for PetID param
// swagger:parameters GetPet DeletePet
type PetIDParam struct {
	// The id of the pet
	//
	// in:path
	// required:true
	ID int `json:"id"`
}

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

// model for create categories success
// swagger:response CategorySuccessResp200
type CategorySuccessResp200 struct {
	// in:body
	data models.CategoryResp
}

// model for delete response
// swagger:response DeleteResponse
type DeleteResponse struct {
	// in:body
	Body struct {
		Msg string `json:"message"`
	} `json:"body"`
}
