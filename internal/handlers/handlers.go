package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(jet.NewOSFileSystemLoader("./html"), jet.InDevelopmentMode())

func Home(rw http.ResponseWriter, r *http.Request) {
	err := renderPage(rw, "home.jet", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
