package feat

const RouterFileTemplate = `package %[1]s

import "github.com/go-chi/chi"

func Router() *chi.Mux {
	var router = chi.NewRouter()

	router.Get("/", get%[2]sHandler)

	return router
}
`
