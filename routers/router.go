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
	router = SetPostApiRoutes(router)
	router = SetPutApiRoutes(router)
	router = SetDeleteApiRoutes(router)
	return router
}
