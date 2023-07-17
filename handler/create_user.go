package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/models/database/transactions"
)

// CreateUser is a function to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	req := new(transactions.Users)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	if req.Name == "" || req.Password == "" || req.Email == "" || req.Id_role == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": errors.New("name is required").Error(),
		})
		return
	}

	_, errr := transactions.GetRoleById(int(req.Id_role))
	if errr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": errors.New("role not found").Error(),
		})
		return
	}

	rep, err := transactions.CreateUser(*req)
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
