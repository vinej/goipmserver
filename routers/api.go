package routers

import (
	"goipmserver/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"goipmserver/core/authentication"
)

func SetGetApiRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/{collection}",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.GetHandler),
		)).Methods("GET")

	return router
}

func SetGetQueryApiRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/{collection}/{squery}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.GetQueryHandler),
		)).Methods("GET")

	return router
}

func SetPostApiRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/{collection}",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.PostHandler),
		)).Methods("POST")

	return router
}

func SetPutApiRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/{collection}",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.PutHandler),
		)).Methods("PUT")

	return router
}

func SetDeleteApiRoutes(router *mux.Router) *mux.Router {
	router.Handle("/api/{collection}",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.DeleteHandler),
		)).Methods("DELETE")

	return router
}
