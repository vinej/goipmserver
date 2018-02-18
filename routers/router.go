package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetGetApiRoutes(router)
	router = SetGetQueryApiRoutes(router)
	router = SetAuthenticationRoutes(router)
	router = SetPostyApiRoutes(router)
	return router
}
