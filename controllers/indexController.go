package controllers

import (
	"embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"waifu_gcx/models"
)

//go:embed views/*.gohtml
var directory embed.FS

var t = template.Must(template.ParseFS(directory, "views/*.gohtml"))

func Index(waifu []models.Waifu) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if r.Method == "POST" {
			picked := waifu[rand.Intn(len(waifu))]
			imageName := strings.ReplaceAll(picked.Name, " ", "_") + ".jpeg"
			imagePath := "/assets/" + imageName
			t.ExecuteTemplate(w, "index.gohtml", struct {
				Id    int
				Name  string
				Image string
			}{
				Id:    picked.Id,
				Name:  picked.Name,
				Image: imagePath,
			})
			return
		} else {
			err := t.ExecuteTemplate(w, "index.gohtml", struct {
				Id    int
				Name  string
				Image string
			}{
				Id:    0,
				Name:  "Default",
				Image: "./assets/Default.jpeg",
			})
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
