package router

import "github.com/gorilla/mux"

// Generate retorna um novo router comas rotas configuradas
func Generate() *mux.Router {
	return mux.NewRouter()
}

