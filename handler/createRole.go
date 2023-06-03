package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/models/database/transactions"
)

type Role struct {
	Name string `json:"name"`
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	req := new(Role)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	if req.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": errors.New("name is required").Error(),
		})
		return
	}

	rep, err := transactions.CreateRole(req.Name)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rep)

}
