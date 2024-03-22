package routes 

// import packages
import (
	"net/http"

	"github.com/cletushunsu/chi_sample/Handler"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

)


// go-chi router function  
func GoChiRouter() http.Handler {

	// define chi router and configure middleware 
	apiRouter := chi.NewRouter()

	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.AllowContentType("application/json"))
	apiRouter.Use(middleware.CleanPath)
	apiRouter.Use(middleware.AllowContentEncoding("deflate", "gzip"))

	// set routes
	apiRouter.Get("/", handler.GetAllItems)
	apiRouter.Get("/{id}", handler.GetItem)
	apiRouter.Post("/", handler.CreateItem)
	apiRouter.Put("/{id}", handler.UpdateItem)
	apiRouter.Patch("/{id}", handler.UpdateItem)
	apiRouter.Delete("/{id}", handler.DeleteItem)

	return apiRouter
}