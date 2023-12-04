package brands

import (
	"errors"
	"github.com/fardinabir/product-store/database/repos"
	"github.com/go-chi/chi/v5"
)

type BrandResource struct {
	Brands BrandStore
}

func NewResource() (*BrandResource, error) {
	brandStore := repos.NewBrandStore()
	if brandStore.DB == nil {
		return nil, errors.New("failed to get Database")
	}
	return &BrandResource{Brands: brandStore}, nil
}

func (rs *BrandResource) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", rs.GetBrands)
	r.Get("/{id:[0-9]+}", rs.GetBrandsById)
	r.Post("/new", rs.CreateBrand)
	r.Patch("/{id:[0-9]+}", rs.UpdateBrand)
	r.Delete("/{id:[0-9]+}", rs.DeleteBrand)
	return r
}
