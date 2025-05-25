package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"waifu_gcx/routers"
)

func main() {
	router := httprouter.New()
	routers.AllRoutes(router)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
