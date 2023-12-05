package products

import (
	"errors"
	"github.com/fardinabir/product-store/database/repos"
	"github.com/go-chi/chi/v5"
)

type ProductResource struct {
	Products ProductStore
}

func NewResource() (*ProductResource, error) {
	productStore := repos.NewProductStore()
	if productStore.DB == nil {
		return nil, errors.New("failed to get Database")
	}
	return &ProductResource{Products: productStore}, nil
}

func (rs *ProductResource) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", rs.GetProducts)
	r.Get("/{id:[0-9]+}", rs.GetProductById)
	r.Post("/new", rs.CreateProduct)
	r.Patch("/{id:[0-9]+}", rs.UpdateProduct)
	r.Delete("/{id:[0-9]+}", rs.DeleteProduct)
	return r
}
