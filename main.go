package main

import (
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"net/http"
	"waifu_gcx/models"
	"waifu_gcx/routers"
)

//go:embed assets/*
var directory embed.FS

//go:embed data/*
var data embed.FS

func main() {
	router := httprouter.New()

	waifu, err := models.LoadJson(data)
	if err != nil {
		panic(err)
	}
	dir, err := fs.Sub(directory, "assets")
	if err != nil {
		panic(err)
	}

	routers.AllRoutes(router, waifu, dir)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	color := "\033[1;32m"
	reset := "\033[0m"
	fmt.Println(color + "\n▗▖ ▗▖ ▗▄▖ ▗▄▄▄▖▗▄▄▄▖▗▖ ▗▖     ▗▄▄▖ ▗▄▄▖▗▖  ▗▖\n▐▌ ▐▌▐▌ ▐▌  █  ▐▌   ▐▌ ▐▌    ▐▌   ▐▌    ▝▚▞▘ \n▐▌ ▐▌▐▛▀▜▌  █  ▐▛▀▀▘▐▌ ▐▌    ▐▌▝▜▌▐▌     ▐▌  \n▐▙█▟▌▐▌ ▐▌▗▄█▄▖▐▌   ▝▚▄▞▘    ▝▚▄▞▘▝▚▄▄▖▗▞▘▝▚▖" + reset)
	fmt.Println("running on localhost:8080")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
