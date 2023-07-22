package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/diegolopezcode/api-crud-complete-chi/models/transactions"
)

func GetRoleById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		data, err := transactions.GetRoles()
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

	resp, err := transactions.GetRoleById(data)
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

func GetRoles(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "id is required",
			"code":    "400",
		})
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

	resp, err := transactions.GetRoleById(data)
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
