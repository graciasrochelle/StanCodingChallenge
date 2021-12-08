package handler

import (
	"net/http"
)

func GetServiceStatus(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "Service up and running")
}
