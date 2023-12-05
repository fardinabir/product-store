package server

import (
	"github.com/fardinabir/product-store/controllers/brands"
	"github.com/fardinabir/product-store/controllers/categories"
	"github.com/fardinabir/product-store/controllers/suppliers"
	"github.com/fardinabir/product-store/database"
	"github.com/go-chi/chi/v5"
	redocmiddleware "github.com/go-openapi/runtime/middleware"
	"io/ioutil"
	"log"
	"net/http"
)

func New() (*chi.Mux, error) {
	r := chi.NewRouter()

	opts := redocmiddleware.RedocOpts{Path: "docs", SpecURL: "swagger.json"}
	docsHandler := redocmiddleware.Redoc(opts, nil)
	r.Handle("/docs", docsHandler)
	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		spec, err := ioutil.ReadFile("docs/swagger.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(spec)
	})

	_, err := database.InitDb()
	if err != nil {
		log.Println("Postgres connection error", err)
		return nil, err
	}

	brandResource, err := brands.NewResource()
	if err != nil {
		log.Println("Failed to get Brand resource, ", err.Error())
		return nil, err
	}

	catResource, err := categories.NewResource()
	if err != nil {
		log.Println("Failed to get Category resource, ", err.Error())
		return nil, err
	}

	supplierResource, err := suppliers.NewResource()
	if err != nil {
		log.Println("Failed to get Category resource, ", err.Error())
		return nil, err
	}
	r.Route("/api", func(r chi.Router) {
		r.Mount("/brands", brandResource.Router())
		r.Mount("/categories", catResource.Router())
		r.Mount("/suppliers", supplierResource.Router())
	})

	return r, nil
}
