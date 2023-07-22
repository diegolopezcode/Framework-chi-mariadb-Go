package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/models/transactions"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	req := new(transactions.TasksUpdate)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Structure json invalid",
			"error":   err.Error(),
		})
		return
	}

	if req.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "id is required",
		})
		return
	}

	data, err := transactions.GetTaskById(int(req.Id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	if data == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "task not found",
		})
		return
	}

	if req.Name == "" {
		req.Name = data.Name
	}

	if req.Description == "" {
		req.Description = data.Description
	}

	if req.Is_complete == "" {
		req.Is_complete = fmt.Sprintf("%v", data.Is_complete)
	}

	resp, err := transactions.UpdateTask(*req)
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
