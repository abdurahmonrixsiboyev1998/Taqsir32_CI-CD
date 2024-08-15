package handlers

import (
	"90HM/internal/models"
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	res := models.Response{
		Status:  "OK",
		Message: "Health check passed",
	}
	jsonResponse(w, res)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	res := models.Response{
		Status:  "OK",
		Message: "Root handler",
	}
	jsonResponse(w, res)
}

func jsonResponse(w http.ResponseWriter, res models.Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

