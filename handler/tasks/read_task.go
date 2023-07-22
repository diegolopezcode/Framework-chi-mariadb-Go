package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/diegolopezcode/api-crud-complete-chi/models/transactions"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user := r.URL.Query().Get("user")
	state := r.URL.Query().Get("state")

	if id == "" && user == "" && state == "" {
		data, err := transactions.GetAllTasks()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("task not found"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
		return
	}

	if id != "" && user == "" && state == "" {
		data, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "id must be a number",
				"code":    "400",
			})
		}

		resp, err := transactions.GetTaskById(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
			return
		}

		if resp == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "task not found",
				"code":    "400",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return

	}

	if id == "" && user != "" && state == "" {
		data, err := strconv.Atoi(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "user must be a number",
				"code":    "400",
			})
		}
		resp, err := transactions.GetTaskByUserId(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
			return
		}

		if resp == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "user not found",
				"code":    "400",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if id == "" && user == "" && state != "" {
		data, err := strconv.ParseBool(state)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "state must be a boolean",
				"code":    "400",
			})
		}
		resp, err := transactions.GetTaskByState(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
				"code":    "400",
			})
			return
		}

		if resp == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "state not found",
				"code":    "400",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

}
