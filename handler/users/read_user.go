package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/diegolopezcode/api-crud-complete-chi/models/transactions"
)

// GetUsers is a function to get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	id_role := r.URL.Query().Get("role")
	if id == "" && id_role == "" {
		data, err := transactions.GetAllUsers()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
		return
	}
	if id_role != "" && id == "" {
		data, err := strconv.Atoi(id_role)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "id_role must be a number",
				"code":    "400",
			})
			return
		}
		resp, err := transactions.GetUsersByRoleId(data)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
			return
		}

		if resp == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("role not found"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

	data, err := strconv.Atoi(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "id must be a number",
			"code":    "400",
		})
		return
	}

	resp, err := transactions.GetUserById(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
			"code":    "400",
		})
		return
	}

	if resp == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("role not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
