package handler

import (
	"encoding/json"
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/models/database/transactions"
)

type RolePermission struct {
	Id_role       int
	Id_permission int
}

func CreateRolePermission(w http.ResponseWriter, r *http.Request) {
	req := new(RolePermission)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	if req.Id_role == 0 || req.Id_permission == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "id_role and id_permission are required",
		})
		return
	}

	rep, err := transactions.CreateRolePermission(req.Id_role, req.Id_permission)
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
