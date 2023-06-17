package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/diegolopezcode/api-crud-complete-chi/models/database/transactions"
)

func GetRolePermissionById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("Id_role")
	idPermission := r.URL.Query().Get("Id_permission")

	fmt.Println("id", id)
	fmt.Println("idPermission", idPermission)

	if id == "" && idPermission == "" {
		data, err := transactions.GetRolePermissions()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
		return
	}

	if id == "" && idPermission != "" {
		data, err := strconv.Atoi(idPermission)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id must be a number"))
			return
		}
		resp, err := transactions.GetRolePermissionByPermissionId(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if id != "" && idPermission == "" {
		data, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id must be a number"))
			return
		}
		resp, err := transactions.GetRolePermissionByRoleId(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if id != "" && idPermission != "" {
		data, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id must be a number"))
			return
		}

		data2, err := strconv.Atoi(idPermission)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id must be a number"))
			return
		}

		resp, err := transactions.GetRolePermissionByRoleIdAndPermissionId(data, data2)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

}
