package handlers

import (
	"html/template"
	"net/http"

	"github.com/hculpan/shipmanager/internal/data"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Message": data.GetMessage(),
	}

	tmpl, err := template.ParseFiles("./internal/templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
