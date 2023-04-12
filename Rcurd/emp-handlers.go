package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	resp := response{
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	name := emp.Name // assuming the name field is present in the Employee struct
	resp.Message = fmt.Sprintf("record saved successfully for %s", name)

	Database.Create(&emp)

	json.NewEncoder(w).Encode(emp)
	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusOK)
}
