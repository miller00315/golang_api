package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate retorna um novo router comas rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurate(r)
}
