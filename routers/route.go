package routers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"waifu_gcx/controllers"
	"waifu_gcx/models"
)

func AllRoutes(router *httprouter.Router) {
	waifu, err := models.LoadJson("./data/waifu.json")
	if err != nil {
		panic(err)
	}
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	router.GET("/", controllers.Index(waifu))
	router.POST("/", controllers.Index(waifu))

}
