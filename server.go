package main

import (
	"goipmserver/routers"
	//"goipmserver/settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":8000", n)
}
