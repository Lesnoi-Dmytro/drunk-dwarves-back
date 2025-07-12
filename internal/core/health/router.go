package health

import "github.com/go-chi/chi"

func Router() *chi.Mux {
	var router = chi.NewRouter()

	router.Get("/", getHealthHandler)

	return router
}
