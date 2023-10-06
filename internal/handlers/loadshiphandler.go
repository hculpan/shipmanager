package handlers

import (
	"html/template"
	"net/http"

	"github.com/hculpan/shipmanager/internal/db"
	"github.com/hculpan/shipmanager/internal/util"
)

func LoadShipHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./internal/templates/loadship.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		ship, err := db.LoadShip(r.FormValue("shipid"))
		if err != nil {
			msg := "Failed to load ship data: " + err.Error()
			util.LogError(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("./internal/templates/shipregistered.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, ship)
	}
}
