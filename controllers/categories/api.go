package categories

import (
	"errors"
	"github.com/fardinabir/product-store/database/repos"
	"github.com/go-chi/chi/v5"
)

type CategoryResource struct {
	Categories CategoryStore
}

func NewResource() (*CategoryResource, error) {
	categoryStore := repos.NewCategoryStore()
	if categoryStore.DB == nil {
		return nil, errors.New("failed to get Database")
	}
	return &CategoryResource{Categories: categoryStore}, nil
}

func (rs *CategoryResource) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", rs.GetCategories)
	r.Get("/{id:[0-9]+}", rs.GetCategoryById)
	r.Post("/new", rs.CreateCategory)
	r.Patch("/{id:[0-9]+}", rs.UpdateCategory)
	r.Delete("/{id:[0-9]+}", rs.DeleteCategory)
	r.Get("/get-tree", rs.GetCategoriesTree)
	return r
}
