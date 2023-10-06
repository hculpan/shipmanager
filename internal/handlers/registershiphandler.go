package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/hculpan/shipmanager/internal/data"
	"github.com/hculpan/shipmanager/internal/db"
	"github.com/hculpan/shipmanager/internal/util"
)

func getIntValue(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		result = 0
	}
	return result
}

func getFloatValue(value string) float32 {
	result, err := strconv.ParseFloat(value, 32)
	if err != nil {
		result = 0.0
	}
	return float32(result)
}

func RegisterShipHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./internal/templates/registership.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		tons := getIntValue(r.FormValue("tons"))
		cargo := getIntValue(r.FormValue("cargo"))
		high := getIntValue(r.FormValue("high"))
		middle := getIntValue(r.FormValue("middle"))
		low := getIntValue(r.FormValue("low"))
		basic := getIntValue(r.FormValue("basic"))
		steward := getIntValue(r.FormValue("steward"))

		totalCost := getFloatValue(r.FormValue("cost"))
		remainingCost := getFloatValue(r.FormValue("remaining"))
		monthly := getIntValue(r.FormValue("monthly"))

		ship := data.Ship{
			Id:             data.GenerateRandomId(),
			Name:           r.FormValue("shipname"),
			Tons:           tons,
			CargoCapacity:  cargo,
			HighPassage:    high,
			BasicPassage:   basic,
			LowPassage:     low,
			MiddlePassage:  middle,
			StewardRating:  steward,
			TotalCost:      totalCost,
			RemainingCost:  remainingCost,
			MonthlyPayment: monthly,
		}
		if err := db.SaveShip(&ship); err != nil {
			msg := "Failed to save ship data: " + err.Error()
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
