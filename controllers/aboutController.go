package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func About() func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := t.ExecuteTemplate(w, "about.gohtml", nil)
		if err != nil {
			panic(err)
		}
	}
}
