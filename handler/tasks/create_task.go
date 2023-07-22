package handler

import (
	"encoding/json"
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/models/transactions"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	req := new(transactions.Tasks)
	req.Is_complete = false
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Structure json invalid",
		})

		return
	}

	if req.Name == "" || req.Description == "" || req.Id_user == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "name, description and id_user are required",
		})
		return
	}

	_, err := transactions.GetUserById(int(req.Id_user))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user not found",
		})
		return
	}

	resp, err := transactions.CreateTask(*req)
	if err != nil {
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
