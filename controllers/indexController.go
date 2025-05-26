package controllers

import (
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"waifu_gcx/entities"
)

//go:embed views/*.gohtml
var Directory embed.FS

var t = template.Must(template.ParseFS(Directory, "views/*.gohtml"))

func Index(waifu []entities.Waifu) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if r.Method == "POST" {
			picked := waifu[rand.Intn(len(waifu))]
			imageName := strings.ReplaceAll(picked.Name, " ", "_") + ".jpeg"
			imagePath := "/assets/" + imageName
			err := t.ExecuteTemplate(w, "index.gohtml", struct {
				Id    int
				Name  string
				Image string
			}{
				Id:    picked.Id,
				Name:  picked.Name,
				Image: imagePath,
			})
			if err != nil {
				panic(err)
			}
			return
		} else {
			err := t.ExecuteTemplate(w, "index.gohtml", struct {
				Id    int
				Name  string
				Image string
			}{
				Id:    0,
				Name:  "Who??",
				Image: "./assets/Default.jpeg",
			})
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
