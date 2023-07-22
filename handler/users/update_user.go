package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/models/transactions"
)

// UpdateUser is a function to update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := new(transactions.UsersUpdate)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	if req.Id == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": errors.New("name is required").Error(),
		})
		return
	}

	data, err := transactions.GetUserById(int(req.Id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user not found"))
		return
	}

	if data == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user not found"))
		return
	}

	if req.Email == "" {
		req.Email = data.Email
	}

	if req.Name == "" {
		req.Name = data.Name
	}

	if req.Password == "" {
		req.Password = data.Password
	}

	if req.Id_role == 0 {
		req.Id_role = data.Id_role
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

	resp, err := transactions.UpdateUser(*req)
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
	json.NewEncoder(w).Encode(resp)

}
