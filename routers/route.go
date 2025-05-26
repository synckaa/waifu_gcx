package routers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/fs"
	"net/http"
	"waifu_gcx/controllers"
	"waifu_gcx/entities"
)

func AllRoutes(router *httprouter.Router, waifu []entities.Waifu, dir fs.FS) {
	router.ServeFiles("/assets/*filepath", http.FS(dir))
	router.GET("/", controllers.Index(waifu))
	router.POST("/", controllers.Index(waifu))
	router.GET("/about", controllers.About())
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		_, err2 := fmt.Fprint(w, err)
		if err2 != nil {
			return
		}
	}
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	})

}
