package suppliers

import (
	"errors"
	"github.com/fardinabir/product-store/database/repos"
	"github.com/go-chi/chi/v5"
)

type SupplierResource struct {
	Suppliers SupplierStore
}

func NewResource() (*SupplierResource, error) {
	supplierStore := repos.NewSupplierStore()
	if supplierStore.DB == nil {
		return nil, errors.New("failed to get Database")
	}
	return &SupplierResource{Suppliers: supplierStore}, nil
}

func (rs *SupplierResource) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", rs.GetSuppliers)
	r.Get("/{id:[0-9]+}", rs.GetSupplierById)
	r.Post("/new", rs.CreateSupplier)
	r.Patch("/{id:[0-9]+}", rs.UpdateSupplier)
	r.Delete("/{id:[0-9]+}", rs.DeleteSupplier)
	return r
}
